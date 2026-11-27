      ANeuralNetworksDevice* device = nullptr;
      const char* buffer = nullptr;
      RETURN_TFLITE_ERROR_IF_NN_ERROR(
          context, nnapi->ANeuralNetworks_getDevice(i, &device),
          "Getting list of available devices", nnapi_errno);
      RETURN_TFLITE_ERROR_IF_NN_ERROR(
          context, nnapi->ANeuralNetworksDevice_getName(device, &buffer),
          "Getting list of available devices", nnapi_errno);
      if (nnapi_cpu != buffer) {
        result->push_back(device);
      }
    }
  }

  return kTfLiteOk;
}

}  // namespace

namespace delegate {
namespace nnapi {

#ifdef TFLITE_NNAPI_ALLOW_MMAP_SHARING
NNMemory::NNMemory(const NnApi* nnapi, const char* name, size_t size) {
  if (name && size > 0) {
    nnapi_ = nnapi;
    byte_size_ = size;
#ifdef __ANDROID__
    fd_ = nnapi_->ASharedMemory_create(name, size);
#else
    // For non-Android platforms ASharedMemory_create needs unique name to
    // create a shared memory object (see nnapi_implementation.cc).
    char shm_name_buffer[L_tmpnam];
    if (tmpnam(shm_name_buffer) == nullptr) {
      shm_name_buffer[0] = '\0';
    }
    // tmpnam will produce a string containing with slashes, but shm_open
    // won't like that.
    shm_region_name_ = std::string(name) + std::string(shm_name_buffer);
    std::replace(shm_region_name_.begin(), shm_region_name_.end(), '/', '-');
    fd_ = nnapi_->ASharedMemory_create(shm_region_name_.c_str(), size);
#endif

    data_ptr_ = reinterpret_cast<uint8_t*>(
        mmap(nullptr, size, PROT_READ | PROT_WRITE, MAP_SHARED, fd_, 0));
    nnapi_->ANeuralNetworksMemory_createFromFd(size, PROT_READ | PROT_WRITE,
                                               fd_, 0, &nn_memory_handle_);
  }
}
#else
NNMemory::NNMemory(const NnApi* /*nnapi*/, const char* /*name*/,
        nn_type, type, dim_array, tensor_value, quant_params, tensor_index);
    TfLiteIntArrayFree(dim_array);
    return result;
  }

  void ClearInputOuputLists() {
    augmented_inputs_.clear();
    augmented_outputs_.clear();
  }

 private:
  // Returns a TF Lite type which has the same memory representation as a
  // provided NN API type.
  TfLiteStatus GetEquivalentToANNType(TfLiteContext* context, int nn_type,
                                      TfLiteType* type) {
    switch (nn_type) {
      case ANEURALNETWORKS_INT32:
        *type = kTfLiteInt32;
        return kTfLiteOk;
      case ANEURALNETWORKS_FLOAT32:
        *type = kTfLiteFloat32;
        return kTfLiteOk;
      default:
        context->ReportError(context,
                             "NN API Delegate: Can't get an equivalent TF Lite "
                             "type for provided NN API type: %d.\n",
                             nn_type);
        return kTfLiteError;
    }
  }

  template <typename T>
  TfLiteStatus AddScalarOperand(T value, int32_t nn_type) {
    ANeuralNetworksOperandType operand_type{.type = nn_type};
    RETURN_TFLITE_ERROR_IF_NN_ERROR(
        context_,
        nnapi_->ANeuralNetworksModel_addOperand(nn_model_, &operand_type),
        "adding operand", nnapi_errno_);
    const int ann_index = operand_mapping_->add_new_non_tensor_operand();
    RETURN_TFLITE_ERROR_IF_NN_ERROR(
        context_,
        nnapi_->ANeuralNetworksModel_setOperandValue(nn_model_, ann_index,
                                                     &value, sizeof(T)),
        "setting new operand value", nnapi_errno_);
    augmented_inputs_.push_back(ann_index);
    return kTfLiteOk;
  }

  template <typename T>
  TfLiteStatus AddVectorOperand(const T* values, uint32_t num_values,
                                int32_t nn_type, float scale,