    Value result = rewriter.create<SelectOp>(loc, op.getType(), compare,
                                             on_value, off_value);

    rewriter.replaceOp(op, {result});

    return success();
  }
};

// Converts InfeedDequeueTuple to XLA HLO create_token, infeed and
// get_tuple_element ops.
//
// All HLO infeed ops expect a HLO token type operand and produce a tuple
// containing a token. This HLO token type is used to order multiple infeed
// operations within a computation. The token type can come from other
// infeed/outfeed/send/recv ops or can be generated using create_token op with
// no operands. Here we emit a create_token op to generate the token type
// operand of infeed.
//
// For example the following IR:
// %0:2 = "tf.InfeedDequeueTuple"() : () -> (tensor<3xi32>, tensor<4xf32>)
//
// would be lowered to
//
// %token = "mhlo.create_token"() : () -> !mhlo.token
// %data_and_token = "mhlo.infeed"(%token) {infeed_config = ""} :
//      (!mhlo.token) -> tuple<tuple<tensor<3xi32>, tensor<4xf32>>,
//      !mhlo.token>
// %data = "mhlo.get_tuple_element"(%data_and_token) {index = 0}
// %0#0 = "mhlo.get_tuple_element"(%data) {index = 0}
// %0#1 = "mhlo.get_tuple_element"(%data) {index = 1}
//
class ConvertInfeedDequeueTupleOp
    : public OpRewritePattern<TF::InfeedDequeueTupleOp> {
 public:
  using OpRewritePattern::OpRewritePattern;

  FailureOr<std::vector<int64_t>> GetTPUInfeedLayoutFromAPI(
      RankedTensorType t) const {
    // Call the TPU API to determine the right infeed layout. Note that
    // this can fail if we're not running on a TPU-enabled node.
    // TODO(kramm): Move this into a separate pass. See b/181724526
    xla::Shape old_shape = xla::TypeToShape(t);
    XLA_Shape old_shape_c = {};
    XLA_Shape new_shape_c = {};
    TfTpu_ExecutorApiFn *executor = tensorflow::tpu::ExecutorApiFn();
    if (!tensorflow::tpu::IsInitialized(executor)) {
      return failure();
    }
    ApiConverter::ToC(old_shape, &old_shape_c);
    executor->TpuTransferManager_GetInfeedLayoutFn(&old_shape_c, &new_shape_c);
    xla::Shape new_shape = ApiConverter::FromC(&new_shape_c);
    ApiConverter::Free(&old_shape_c);
    ApiConverter::Free(&new_shape_c);

    xla::Layout layout = new_shape.layout();
    auto minor_to_major = layout.minor_to_major();
    return std::vector<int64_t>(minor_to_major.begin(), minor_to_major.end());
  }

  FailureOr<Attribute> GetLayout(const Type &type,
                                 PatternRewriter &rewriter) const {
    auto i64_type = rewriter.getIntegerType(64);
    if (type.isa<TupleType>()) {
      TupleType tuple_type = type.dyn_cast<TupleType>();
      std::vector<mlir::Attribute> v;
      for (const mlir::Type &t : tuple_type.getTypes()) {
        auto layout = GetLayout(t, rewriter);
        if (failed(layout)) return failure();
        v.push_back(layout.getValue());
      }
      auto tuple_element = rewriter.create<GetTupleElementOp>(
          op.getLoc(), idx_and_type.value(), data_tuple,
          rewriter.getI32IntegerAttr(idx_and_type.index()));
      results.push_back(tuple_element);
    }
    rewriter.replaceOp(op, ValueRange(results));
    return success();
  }
};

// Converts tf.OutfeedEnqueueTuple to XLA HLO tuple, create_token and outfeed
// ops.
//
// XLA HLO outfeed op expects a token, which we generate by emitting an
// create_token op.
//
// For example the following IR:
// "tf.OutfeedEnqueueTuple"(%val_1, %val_2) : (tensor<3xi32>, tensor<4xf32>) ->
//      ()
//
// would be lowered to
//
// %tuple = "mhlo.tuple"(%val_1, %val_2) : (tensor<3xi32>, tensor<4xf32>) ->
//      tuple<tensor<3xi32>, tensor<4xf32>>
// %token = "mhlo.create_token"() : () -> !mhlo.token
// %outfeed_token = "mhlo.outfeed"(%tuple, %token) {outfeed_config = ""} :
//      (tuple<tensor<3xi32>, tensor<4xf32>>, !mhlo.token) -> !mhlo.token
//
class ConvertOutfeedEnqueueTupleOp
    : public OpRewritePattern<TF::OutfeedEnqueueTupleOp> {
 public:
  using OpRewritePattern::OpRewritePattern;

  LogicalResult matchAndRewrite(TF::OutfeedEnqueueTupleOp op,
                                PatternRewriter &rewriter) const override {
    auto token_type = mhlo::TokenType::get(rewriter.getContext());
    auto tuple = rewriter.create<TupleOp>(op.getLoc(), op.inputs());
    auto token = rewriter.create<CreateTokenOp>(op.getLoc(), token_type);
    rewriter.create<OutfeedOp>(op.getLoc(), token_type, tuple, token,
                               /*outfeed_config=*/rewriter.getStringAttr(""));
    rewriter.eraseOp(op);
    return success();
  }
};

// Converts tf.TopKV2 to XLA HLO iota, sort, and slice ops when k is a constant.
//
// tf.TopKV2 sorts along last dimension of the input tensor and then returns
// the top K components' values and indices. This is translated into a few
// ops in XLA HLO: first generating an integer sequence for the indices,
// then sort both the original input tensor and the indices togheter, and