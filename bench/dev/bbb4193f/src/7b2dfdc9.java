                  props_,
                  ((bitField0_ & 0x00000004) == 0x00000004),
                  getParentForChildren(),
                  isClean());
          props_ = null;
        }
        return propsBuilder_;
      }

      // @@protoc_insertion_point(builder_scope:com.alibaba.otter.canal.protocol.TransactionEnd)
    }

    static {
      defaultInstance = new TransactionEnd(true);
      defaultInstance.initFields();
    }

    // @@protoc_insertion_point(class_scope:com.alibaba.otter.canal.protocol.TransactionEnd)
  }

  public interface PairOrBuilder extends
      // @@protoc_insertion_point(interface_extends:com.alibaba.otter.canal.protocol.Pair)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>optional string key = 1;</code>
     */
    boolean hasKey();
    /**
     * <code>optional string key = 1;</code>
     */
    String getKey();
    /**
     * <code>optional string key = 1;</code>
     */
    com.google.protobuf.ByteString
        getKeyBytes();

    /**
     * <code>optional string value = 2;</code>
     */
    boolean hasValue();
    /**
     * <code>optional string value = 2;</code>
     */
    String getValue();
    /**
     * <code>optional string value = 2;</code>
     */
    com.google.protobuf.ByteString
        getValueBytes();
  }
  /**
   * Protobuf type {@code com.alibaba.otter.canal.protocol.Pair}
   *
   * <pre>
   **预留扩展*
   * </pre>
   */
        internalGetFieldAccessorTable() {
      return CanalEntry.internal_static_com_alibaba_otter_canal_protocol_Pair_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              Pair.class, Builder.class);
    }

    public static com.google.protobuf.Parser<Pair> PARSER =
        new com.google.protobuf.AbstractParser<Pair>() {
      public Pair parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Pair(input, extensionRegistry);
      }
    };

    @Override
    public com.google.protobuf.Parser<Pair> getParserForType() {
      return PARSER;
    }

    private int bitField0_;
    public static final int KEY_FIELD_NUMBER = 1;
    private Object key_;
    /**
     * <code>optional string key = 1;</code>
     */
    public boolean hasKey() {
      return ((bitField0_ & 0x00000001) == 0x00000001);
    }
    /**
     * <code>optional string key = 1;</code>
     */
    public String getKey() {
      Object ref = key_;
      if (ref instanceof String) {
        return (String) ref;
      } else {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        String s = bs.toStringUtf8();
        if (bs.isValidUtf8()) {
          key_ = s;
        }
        return s;
      }
    }
    /**
     * <code>optional string key = 1;</code>
     */
    public com.google.protobuf.ByteString
        getKeyBytes() {
      Object ref = key_;
      if (ref instanceof String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (String) ref);
        key_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int VALUE_FIELD_NUMBER = 2;
    private Object value_;
    /**
     * <code>optional string value = 2;</code>
     */
    public boolean hasValue() {
      return ((bitField0_ & 0x00000002) == 0x00000002);
    }
    /**
     * <code>optional string value = 2;</code>
      public final boolean isInitialized() {
        return true;
      }

      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        Pair parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (Pair) e.getUnfinishedMessage();
          throw e;
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int bitField0_;

      private Object key_ = "";
      /**
       * <code>optional string key = 1;</code>
       */
      public boolean hasKey() {
        return ((bitField0_ & 0x00000001) == 0x00000001);
      }
      /**
       * <code>optional string key = 1;</code>
       */
      public String getKey() {
        Object ref = key_;
        if (!(ref instanceof String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          String s = bs.toStringUtf8();
          if (bs.isValidUtf8()) {
            key_ = s;
          }
          return s;
        } else {
          return (String) ref;
        }
      }
      /**
       * <code>optional string key = 1;</code>
       */
      public com.google.protobuf.ByteString
          getKeyBytes() {
        Object ref = key_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b =
              com.google.protobuf.ByteString.copyFromUtf8(
                  (String) ref);
          key_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>optional string key = 1;</code>
       */
      public Builder setKey(
          String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
        key_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>optional string key = 1;</code>
       */
      public Builder clearKey() {
        bitField0_ = (bitField0_ & ~0x00000001);
        key_ = getDefaultInstance().getKey();
        onChanged();
        return this;
      }
      /**
       * <code>optional string key = 1;</code>
       */
      public Builder setKeyBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  bitField0_ |= 0x00000001;
        key_ = value;
        onChanged();
        return this;
      }

      private Object value_ = "";
      /**
       * <code>optional string value = 2;</code>
       */
      public boolean hasValue() {
        return ((bitField0_ & 0x00000002) == 0x00000002);
      }
      /**
       * <code>optional string value = 2;</code>
       */
      public String getValue() {
        Object ref = value_;
        if (!(ref instanceof String)) {