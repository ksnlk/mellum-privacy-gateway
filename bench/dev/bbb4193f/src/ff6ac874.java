    public com.alibaba.otter.canal.protocol.AdminPacket.Handshake getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  public interface ClientAuthOrBuilder extends
      // @@protoc_insertion_point(interface_extends:com.alibaba.otter.canal.protocol.ClientAuth)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string username = 1;</code>
     */
    java.lang.String getUsername();
    /**
     * <code>string username = 1;</code>
     */
    com.google.protobuf.ByteString
        getUsernameBytes();

    /**
     * <pre>
     * hashed password with seeds from Handshake message
     * </pre>
     *
     * <code>bytes password = 2;</code>
     */
    com.google.protobuf.ByteString getPassword();

    /**
     * <pre>
     * in seconds
     * </pre>
     *
     * <code>int32 net_read_timeout = 3;</code>
     */
    int getNetReadTimeout();

    /**
     * <pre>
     * in seconds
     * </pre>
     *
     * <code>int32 net_write_timeout = 4;</code>
     */
    int getNetWriteTimeout();

    public com.alibaba.otter.canal.protocol.AdminPacket.ClientAuth.NetReadTimeoutPresentCase getNetReadTimeoutPresentCase();

    public com.alibaba.otter.canal.protocol.AdminPacket.ClientAuth.NetWriteTimeoutPresentCase getNetWriteTimeoutPresentCase();
  }
    }
    /**
     * <code>string username = 1;</code>
     */
    public com.google.protobuf.ByteString
        getUsernameBytes() {
      java.lang.Object ref = username_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        username_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int PASSWORD_FIELD_NUMBER = 2;
    private com.google.protobuf.ByteString password_;
    /**
     * <pre>
     * hashed password with seeds from Handshake message
     * </pre>
     *
     * <code>bytes password = 2;</code>
     */
    public com.google.protobuf.ByteString getPassword() {
      return password_;
    }

    public static final int NET_READ_TIMEOUT_FIELD_NUMBER = 3;
    /**
     * <pre>
     * in seconds
     * </pre>
     *
     * <code>int32 net_read_timeout = 3;</code>
     */
    public int getNetReadTimeout() {
      if (netReadTimeoutPresentCase_ == 3) {
        return (java.lang.Integer) netReadTimeoutPresent_;
      }
      return 0;
    }

    public static final int NET_WRITE_TIMEOUT_FIELD_NUMBER = 4;
    /**
     * <pre>
     * in seconds
     * </pre>
        username_ = getDefaultInstance().getUsername();
        onChanged();
        return this;
      }
      /**
       * <code>string username = 1;</code>
       */
      public Builder setUsernameBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        username_ = value;
        onChanged();
        return this;
      }

      private com.google.protobuf.ByteString password_ = com.google.protobuf.ByteString.EMPTY;
      /**
       * <pre>
       * hashed password with seeds from Handshake message
       * </pre>
       *
       * <code>bytes password = 2;</code>
       */
      public com.google.protobuf.ByteString getPassword() {
        return password_;
      }
      /**
       * <pre>
       * hashed password with seeds from Handshake message
       * </pre>
       *
       * <code>bytes password = 2;</code>
       */
      public Builder setPassword(com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        password_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * hashed password with seeds from Handshake message
       * </pre>
       *
       * <code>bytes password = 2;</code>
       */
      public Builder clearPassword() {
        
        password_ = getDefaultInstance().getPassword();
        onChanged();
        return this;
      }

      /**
       * <pre>
       * in seconds
       * </pre>
       *
       * <code>int32 net_read_timeout = 3;</code>
       */
      public int getNetReadTimeout() {
        if (netReadTimeoutPresentCase_ == 3) {
          return (java.lang.Integer) netReadTimeoutPresent_;
        }
        return 0;
      }
      /**
       * <pre>
       * in seconds
       * </pre>