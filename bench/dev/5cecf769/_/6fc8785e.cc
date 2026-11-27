    return Unimplemented(
        "ReducePrecision cannot accept non-floating-point type %s.",
        PrimitiveType_Name(src_ty));
  }

  // Integer and float types for casting and constant generation.
  llvm::Type* float_type = x->getType();
  int64 nbits = float_type->getPrimitiveSizeInBits();
  llvm::IntegerType* int_type = b->getIntNTy(nbits);

  // SignificandWidth includes the implicit extra bit.
  int src_mantissa_bits = primitive_util::SignificandWidth(src_ty) - 1;
  int src_exponent_bits = nbits - 1 - src_mantissa_bits;

  // Cast the input value to an integer for bitwise manipulation.
  llvm::Value* x_as_int = b->CreateBitCast(x, int_type);

  if (dest_mantissa_bits < src_mantissa_bits) {
    // Last remaining mantissa bit.
    APInt last_mantissa_bit_mask(nbits, 1);
    last_mantissa_bit_mask <<= src_mantissa_bits - dest_mantissa_bits;

    // Compute rounding bias for round-to-nearest with ties to even.  This is
    // equal to a base value of 0111... plus one bit if the last remaining
    // mantissa bit is 1.
    APInt base_rounding_bias = last_mantissa_bit_mask.lshr(1) - 1;
    llvm::Value* x_last_mantissa_bit = b->CreateLShr(
        b->CreateAnd(x_as_int,
                     llvm::ConstantInt::get(int_type, last_mantissa_bit_mask)),
        (src_mantissa_bits - dest_mantissa_bits));
    llvm::Value* x_rounding_bias =
        b->CreateAdd(x_last_mantissa_bit,
                     llvm::ConstantInt::get(int_type, base_rounding_bias));

    // Add rounding bias, and mask out truncated bits.  Note that the case
    // where adding the rounding bias overflows into the exponent bits is
    // correct; the non-masked mantissa bits will all be zero, and the
    // exponent will be incremented by one.
    APInt truncation_mask = ~(last_mantissa_bit_mask - 1);
    x_as_int = b->CreateAdd(x_as_int, x_rounding_bias);
    x_as_int = b->CreateAnd(x_as_int,
                            llvm::ConstantInt::get(int_type, truncation_mask));
  }

  if (dest_exponent_bits < src_exponent_bits) {
    APInt sign_bit_mask(nbits, 1);
    sign_bit_mask <<= nbits - 1;

    APInt exp_bits_mask(nbits, 1);
    exp_bits_mask = ((exp_bits_mask << src_exponent_bits) - 1)
                    << src_mantissa_bits;

    // An exponent of 2^(n-1)-1 -- that is, 0111... with the zero in the most-
    // significant bit -- is equal to 1.0f for all exponent sizes.  Adding
    // 2^(n-1)-1 to this gives us the highest non-infinite exponent for a bit-
    // size of n, and subtracting 2^(n-1)-1 from this gives us the lowest'
    // exponent (corresponding to 0.0f).
    //
    // Thus, the f32 exponent corresponding to the highest non-infinite
    // exponent for a bit size of n is (2^7-1) + 2^(n-1)-1, and the f32
    // exponent corresponding to the lowest exponent for a bit size of n is
    // (2^7-1) - 2^(n-1)-1.
    //
    // Note that we have already checked that exponents_bits >= 1.
    APInt exponent_bias(nbits, 1);
    exponent_bias = (exponent_bias << (src_exponent_bits - 1)) - 1;

    APInt reduced_exponent_bias(nbits, 1);
    reduced_exponent_bias =
        (reduced_exponent_bias << (dest_exponent_bits - 1)) - 1;

    APInt reduced_max_exponent = exponent_bias + reduced_exponent_bias;
    APInt reduced_min_exponent = exponent_bias - reduced_exponent_bias;

    // Do we overflow or underflow?
    llvm::Value* x_exponent =
        b->CreateAnd(x_as_int, llvm::ConstantInt::get(int_type, exp_bits_mask));
    llvm::Value* x_overflows = b->CreateICmpUGT(
        x_exponent, llvm::ConstantInt::get(
                        int_type, reduced_max_exponent << src_mantissa_bits));
    llvm::Value* x_underflows = b->CreateICmpULE(
        x_exponent, llvm::ConstantInt::get(
                        int_type, reduced_min_exponent << src_mantissa_bits));

    // Compute appropriately-signed values of zero and infinity.
    llvm::Value* x_signed_zero =
        b->CreateAnd(x_as_int, llvm::ConstantInt::get(int_type, sign_bit_mask));
    llvm::Value* x_signed_inf = b->CreateOr(
        x_signed_zero, llvm::ConstantInt::get(int_type, exp_bits_mask));

    // Force to zero or infinity if overflow or underflow.  (Note that this
    // truncates all denormal values to zero, rather than rounding them.)
    x_as_int = b->CreateSelect(x_overflows, x_signed_inf, x_as_int);
    x_as_int = b->CreateSelect(x_underflows, x_signed_zero, x_as_int);
  }

  // Cast the result back to a floating-point type.
  llvm::Value* result = b->CreateBitCast(x_as_int, float_type);