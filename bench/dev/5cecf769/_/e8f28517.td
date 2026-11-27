}

def TF_CollectiveReduceV2Op : TF_Op<"CollectiveReduceV2", []> {
  let summary = [{
Mutually reduces multiple tensors of identical type and shape.
  }];

  let arguments = (ins
    TF_FpOrI32OrI64Tensor:$input,
    TF_Int32Tensor:$group_size,
    TF_Int32Tensor:$group_key,
    TF_Int32Tensor:$instance_key,
    Variadic<TF_ResourceTensor>:$ordering_token,

    TF_AnyStrAttrOf<["Min", "Max", "Mul", "Add"]>:$merge_op,
    TF_AnyStrAttrOf<["Id", "Div"]>:$final_op,
    DefaultValuedAttr<StrAttr, "auto">:$communication_hint,
    DefaultValuedAttr<F32Attr, "0.0f">:$timeout_seconds,
    DefaultValuedAttr<I64Attr, "-1">:$max_subdivs_per_device
  );

  let results = (outs
    TF_FpOrI32OrI64Tensor:$data
  );

  TF_DerivedOperandSizeAttr Nordering_token = TF_DerivedOperandSizeAttr<4>;
  TF_DerivedOperandTypeAttr T = TF_DerivedOperandTypeAttr<0>;
}

def TF_ComplexOp : TF_Op<"Complex", [NoSideEffect, ResultsBroadcastableShape]> {
  let summary = "Converts two real numbers to a complex number.";

  let description = [{
Given a tensor `real` representing the real part of a complex number, and a
tensor `imag` representing the imaginary part of a complex number, this
operation returns complex numbers elementwise of the form \\(a + bj\\), where
*a* represents the `real` part and *b* represents the `imag` part.

The input tensors `real` and `imag` must have the same shape.

For example:

```
# tensor 'real' is [2.25, 3.25]
# tensor `imag` is [4.75, 5.75]
tf.complex(real, imag) ==> [[2.25 + 4.75j], [3.25 + 5.75j]]
```
  }];

  let arguments = (ins
    TF_F32OrF64Tensor:$real,
    Arg<TF_StrTensor, [{Filename of a vocabulary text file.}]>:$filename,

    Confined<I64Attr, [IntMinValue<-2>]>:$key_index,
    Confined<I64Attr, [IntMinValue<-2>]>:$value_index,
    Confined<DefaultValuedAttr<I64Attr, "-1">, [IntMinValue<-1>]>:$vocab_size,
    DefaultValuedAttr<StrAttr, "\t">:$delimiter,
    DefaultValuedAttr<I64Attr, "0">:$offset
  );

  let results = (outs);
}

def TF_InitializeTableV2Op : TF_Op<"InitializeTableV2", []> {
  let summary = [{
Table initializer that takes two tensors for keys and values respectively.
  }];

  let arguments = (ins
    Arg<TF_ResourceTensor, [{Handle to a table which will be initialized.}], [TF_LookupTableWrite]>:$table_handle,
    Arg<TF_Tensor, [{Keys of type Tkey.}]>:$keys,
    Arg<TF_Tensor, [{Values of type Tval.}]>:$values
  );

  let results = (outs);

  TF_DerivedOperandTypeAttr Tkey = TF_DerivedOperandTypeAttr<1>;
  TF_DerivedOperandTypeAttr Tval = TF_DerivedOperandTypeAttr<2>;
}

def TF_InplaceAddOp : TF_Op<"InplaceAdd", [NoSideEffect, TF_AllTypesMatch<["x", "y"]>]> {
  let summary = "Adds v into specified rows of x.";

  let description = [{
Computes y = x; y[i, :] += v; return y.
  }];

  let arguments = (ins
    Arg<TF_Tensor, [{A `Tensor` of type T.}]>:$x,
    Arg<TF_Int32Tensor, [{A vector. Indices into the left-most dimension of `x`.}]>:$i,
    Arg<TF_Tensor, [{A `Tensor` of type T. Same dimension sizes as x except the first dimension, which must be the same as i's size.}]>:$v
  );

  let results = (outs
    Res<TF_Tensor, [{A `Tensor` of type T. An alias of `x`. The content of `y` is undefined if there are duplicates in `i`.}]>:$y
  );

  TF_DerivedOperandTypeAttr T = TF_DerivedOperandTypeAttr<0>;
}

def TF_InplaceUpdateOp : TF_Op<"InplaceUpdate", [NoSideEffect]> {
  let summary = "Updates specified rows 'i' with values 'v'.";
  );

  let results = (outs
    Res<TF_Tensor, [{The folded tensor.}]>:$output
  );

  TF_DerivedOperandTypeAttr T = TF_DerivedOperandTypeAttr<0>;
  TF_DerivedOperandTypeAttr Tpaddings = TF_DerivedOperandTypeAttr<1>;
}

def TF_MlirLocalVarOp : TF_Op<"MlirLocalVarOp", []> {
  let summary = "Creates a handle to an in-scope variable.";

  let description = [{
Used by internal passes for temporary representation of local state, which will
be eventually removed.
  }];

  let arguments = (ins);

  let results = (outs
    Res<TF_ResourceTensor, "", [TF_VariableAlloc]>:$resource
  );
}

def TF_MlirPassthroughOp : TF_Op<"MlirPassthroughOp", [NoSideEffect]> {
  let summary = [{
Wraps an arbitrary MLIR computation expressed as a module with a main() function.
  }];

  let description = [{
This operation does not have an associated kernel and is not intended to be
executed in a regular TensorFlow session. Instead it is intended to be used for
testing or for special case where a user intends to pass custom MLIR computation
through a TensorFlow graph with the intent of having custom tooling processing
it downstream (when targeting a different environment, like TensorFlow lite for
example).
The MLIR module is expected to have a main() function that will be used as an
entry point. The inputs to the operations will be passed as argument to the
main() function and the returned values of the main function mapped to the
outputs.
Example usage:

```
import tensorflow as tf
from tensorflow.compiler.mlir.tensorflow.gen_mlir_passthrough_op import mlir_passthrough_op

mlir_module = '''python
func @main(%arg0 : tensor<10xf32>, %arg1 : tensor<10xf32>) -> tensor<10x10xf32> {
   %add = "magic.op"(%arg0, %arg1) : (tensor<10xf32>, tensor<10xf32>) -> tensor<10x10xf32>
   return %ret : tensor<10x10xf32>
collisions.

This op creates a mutable hash table, specifying the type of its keys and
values. Each value must be a scalar. Data can be inserted into the table using
the insert operations. It does not support the initialization operation.
  }];

  let arguments = (ins
    Arg<TF_Tensor, [{The key used to represent empty key buckets internally. Must not
be used in insert or lookup operations.}]>:$empty_key,
    TF_Tensor:$deleted_key,

    DefaultValuedAttr<StrAttr, "">:$container,
    DefaultValuedAttr<StrAttr, "">:$shared_name,
    DefaultValuedAttr<BoolAttr, "false">:$use_node_name_sharing,
    TypeAttr:$value_dtype,
    DefaultValuedAttr<TF_ShapeAttr, "llvm::ArrayRef<int64_t>({})">:$value_shape,
    DefaultValuedAttr<I64Attr, "131072">:$initial_num_buckets,
    DefaultValuedAttr<F32Attr, "0.8f">:$max_load_factor
  );

  let results = (outs
    Res<TF_ResourceTensor, [{Handle to a table.}], [TF_LookupTableAlloc]>:$table_handle
  );

  TF_DerivedOperandTypeAttr key_dtype = TF_DerivedOperandTypeAttr<0>;
}

def TF_MutableHashTableOfTensorsV2Op : TF_Op<"MutableHashTableOfTensorsV2", []> {
  let summary = "Creates an empty hash table.";

  let description = [{
This op creates a mutable hash table, specifying the type of its keys and
values. Each value must be a vector. Data can be inserted into the table using
the insert operations. It does not support the initialization operation.
  }];

  let arguments = (ins
    DefaultValuedAttr<StrAttr, "">:$container,
    DefaultValuedAttr<StrAttr, "">:$shared_name,
    DefaultValuedAttr<BoolAttr, "false">:$use_node_name_sharing,
    TypeAttr:$key_dtype,
    TypeAttr:$value_dtype,
    DefaultValuedAttr<TF_ShapeAttr, "llvm::ArrayRef<int64_t>({})">:$value_shape
  );

  let results = (outs
    Res<TF_ResourceTensor, [{Handle to a table.}], [TF_LookupTableAlloc]>:$table_handle
  );
}

  );

  let results = (outs
    Res<TensorOf<[TF_Float16, TF_Float32, TF_Float64]>, [{Random values with specified shape.}]>:$output
  );

  TF_DerivedOperandTypeAttr T = TF_DerivedOperandTypeAttr<0>;
  TF_DerivedOperandTypeAttr Tseed = TF_DerivedOperandTypeAttr<1>;
  TF_DerivedOperandTypeAttr dtype = TF_DerivedOperandTypeAttr<2>;
}

def TF_StatelessRandomGetAlgOp : TF_Op<"StatelessRandomGetAlg", []> {
  let summary = "Picks the best counter-based RNG algorithm based on device.";

  let description = [{
This op picks the best counter-based RNG algorithm based on device.
  }];

  let arguments = (ins);

  let results = (outs
    Res<TF_Int32Tensor, [{The RNG algorithm (shape int32[]).}]>:$alg
  );
}

def TF_StatelessRandomGetKeyCounterOp : TF_Op<"StatelessRandomGetKeyCounter", []> {
  let summary = [{
Scrambles seed into key and counter, using the best algorithm based on device.
  }];

  let description = [{
This op scrambles a shape-[2] seed into a key and a counter, both needed by counter-based RNG algorithms. The scrambing uses the best algorithm based on device. The scrambling is opaque but approximately satisfies the property that different seed results in different key/counter pair (which will in turn result in different random numbers).
  }];

  let arguments = (ins
    Arg<TF_I32OrI64Tensor, [{2 seeds (shape [2]).}]>:$seed
  );

  let results = (outs
    Res<TF_Uint64Tensor, [{Key for the counter-based RNG algorithm (shape uint64[1]).}]>:$key,
    Res<TF_Uint64Tensor, [{Counter for the counter-based RNG algorithm. Since counter size is algorithm-dependent, this output will be right-padded with zeros to reach shape uint64[2] (the current maximal counter size among algorithms).}]>:$counter
  );

  TF_DerivedOperandTypeAttr Tseed = TF_DerivedOperandTypeAttr<0>;
}

def TF_StatelessRandomGetKeyCounterAlgOp : TF_Op<"StatelessRandomGetKeyCounterAlg", [NoSideEffect, TF_NoConstantFold]> {
  let summary = [{
Picks the best algorithm based on device, and scrambles seed into key and counter.
  }];

  let description = [{
This op picks the best counter-based RNG algorithm based on device, and scrambles a shape-[2] seed into a key and a counter, both needed by the counter-based algorithm. The scrambling is opaque but approximately satisfies the property that different seed results in different key/counter pair (which will in turn result in different random numbers).
  }];

  let arguments = (ins
    Arg<TF_I32OrI64Tensor, [{2 seeds (shape [2]).}]>:$seed
  );

  let results = (outs
    Res<TF_Uint64Tensor, [{Key for the counter-based RNG algorithm (shape uint64[1]).}]>:$key,
    Res<TF_Uint64Tensor, [{Counter for the counter-based RNG algorithm. Since counter size is algorithm-dependent, this output will be right-padded with zeros to reach shape uint64[2] (the current maximal counter size among algorithms).}]>:$counter,
    Res<TF_Int32Tensor, [{The RNG algorithm (shape int32[]).}]>:$alg
  );

  TF_DerivedOperandTypeAttr Tseed = TF_DerivedOperandTypeAttr<0>;
}

def TF_StatelessRandomNormalOp : TF_Op<"StatelessRandomNormal", [NoSideEffect, TF_NoConstantFold]> {
  let summary = [{
Outputs deterministic pseudorandom values from a normal distribution.
  }];

def TF_XlaGatherOp : TF_Op<"XlaGather", [NoSideEffect]> {
  let summary = "Wraps the XLA Gather operator documented at";

  let description = [{
https://www.tensorflow.org/xla/operation_semantics#gather
  }];

  let arguments = (ins
    Arg<TensorOf<[TF_Bfloat16, TF_Bool, TF_Complex128, TF_Complex64, TF_Float16, TF_Float32, TF_Float64, TF_Int16, TF_Int32, TF_Int64, TF_Int8, TF_Qint32, TF_Qint8, TF_Quint8, TF_Uint16, TF_Uint32, TF_Uint64, TF_Uint8]>, [{The array we're gathering from.}]>:$operand,
    Arg<TF_I32OrI64Tensor, [{Array containing the starting indices of the slices we gather.}]>:$start_indices,
    Arg<TF_I32OrI64Tensor, [{slice_sizes[i] is the bounds for the slice on dimension i.}]>:$slice_sizes,

    StrAttr:$dimension_numbers,
    BoolAttr:$indices_are_sorted
  );

  let results = (outs
    TensorOf<[TF_Bfloat16, TF_Bool, TF_Complex128, TF_Complex64, TF_Float16, TF_Float32, TF_Float64, TF_Int16, TF_Int32, TF_Int64, TF_Int8, TF_Qint32, TF_Qint8, TF_Quint8, TF_Uint16, TF_Uint32, TF_Uint64, TF_Uint8]>:$output
  );

  TF_DerivedOperandTypeAttr T = TF_DerivedOperandTypeAttr<0>;
  TF_DerivedOperandTypeAttr Tindices = TF_DerivedOperandTypeAttr<1>;
}

def TF_XlaKeyValueSortOp : TF_Op<"XlaKeyValueSort", [NoSideEffect]> {
  let summary = "Wraps the XLA Sort operator, documented at";

  let description = [{
https://www.tensorflow.org/performance/xla/operation_semantics#sort
.

Sorts a tensor. Currently only sorts in ascending order are supported.
  }];

  let arguments = (ins
    Arg<TF_IntOrFpTensor, [{A `Tensor` of type K.}]>:$keys,
    Arg<TF_Tensor, [{A `Tensor` of type V.}]>:$values
  );

  let results = (outs
    Res<TF_IntOrFpTensor, [{A `Tensor` of type K.}]>:$sorted_keys,
    Res<TF_Tensor, [{A `Tensor` of type V.}]>:$sorted_values
  );

  TF_DerivedOperandTypeAttr K = TF_DerivedOperandTypeAttr<0>;
  TF_DerivedOperandTypeAttr V = TF_DerivedOperandTypeAttr<1>;
}

def TF_XlaPadOp : TF_Op<"XlaPad", [NoSideEffect]> {
  let summary = "Wraps the XLA Pad operator, documented at";
'dynamic_shapes' contains dynamic shapes of arguments whose shapes were not
known statically at TPUReplication rewrite time.
'metadata' is a serialized TPUCompileMetadataProto describing the shapes and
types of the inputs to the computation, as well as a mapping onto the TPU pod
topology.
'program' output is a string key that is passed to the TPUExecute op and used to
look up the program in the compilation cache.
  }];

  let arguments = (ins
    Variadic<TF_Int64Tensor>:$dynamic_shapes,

    DefaultValuedAttr<StrAttr, "">:$mlir_module,
    StrAttr:$metadata
  );

  let results = (outs
    TF_StrTensor:$compilation_status,
    Variadic<TF_StrTensor>:$program
  );

  TF_DerivedOperandSizeAttr NumDynamicShapes = TF_DerivedOperandSizeAttr<0>;
  TF_DerivedResultSizeAttr num_computations = TF_DerivedResultSizeAttr<1>;
}

def TF__TPUCompileMlirPlaceholderProgramKeyOp : TF_Op<"_TPUCompileMlirPlaceholderProgramKey", []> {
  let summary = [{
Placeholder program key (compilation cache key) of a _TPUCompileMlir `program`.
  }];

  let description = [{
This op can be used when certain rewrite passes materialize ops that require a
program key but the _TPUCompileMlir op has not been added yet. Subsequent
rewrite passes must replace this op with a _TPUCompileMlir op `program` output.
  }];

  let arguments = (ins);

  let results = (outs
    TF_StrTensor:$program
  );
}

def TF__UnaryOpsCompositionOp : TF_Op<"_UnaryOpsComposition", [NoSideEffect, TF_SameOperandsAndResultTypeResolveRef]> {
  let summary = [{
*NOTE*: Do not invoke this operator directly in Python. Graph rewrite pass is
  }];

  let description = [{
expected to create these operators.
  }];