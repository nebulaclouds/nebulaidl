// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/array_job.proto

package nebulaidl.plugins;

public final class ArrayJobOuterClass {
  private ArrayJobOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface ArrayJobOrBuilder extends
      // @@protoc_insertion_point(interface_extends:nebulaidl.plugins.ArrayJob)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <pre>
     * Defines the maximum number of instances to bring up concurrently at any given point. Note that this is an
     * optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
     * running instances might be more. This has to be a positive number if assigned. Default value is size.
     * </pre>
     *
     * <code>int64 parallelism = 1;</code>
     */
    long getParallelism();

    /**
     * <pre>
     * Defines the number of instances to launch at most. This number should match the size of the input if the job
     * requires processing of all input data. This has to be a positive number.
     * In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
     * </pre>
     *
     * <code>int64 size = 2;</code>
     */
    long getSize();

    /**
     * <pre>
     * An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
     * the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
     * assigned. Default value is size (if specified).
     * </pre>
     *
     * <code>int64 min_successes = 3;</code>
     */
    long getMinSuccesses();

    /**
     * <pre>
     * If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
     * job can be marked successful.
     * </pre>
     *
     * <code>float min_success_ratio = 4;</code>
     */
    float getMinSuccessRatio();

    public nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.SuccessCriteriaCase getSuccessCriteriaCase();
  }
  /**
   * <pre>
   * Describes a job that can process independent pieces of data concurrently. Multiple copies of the runnable component
   * will be executed concurrently.
   * </pre>
   *
   * Protobuf type {@code nebulaidl.plugins.ArrayJob}
   */
  public  static final class ArrayJob extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:nebulaidl.plugins.ArrayJob)
      ArrayJobOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use ArrayJob.newBuilder() to construct.
    private ArrayJob(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private ArrayJob() {
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private ArrayJob(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      int mutable_bitField0_ = 0;
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 8: {

              parallelism_ = input.readInt64();
              break;
            }
            case 16: {

              size_ = input.readInt64();
              break;
            }
            case 24: {
              successCriteriaCase_ = 3;
              successCriteria_ = input.readInt64();
              break;
            }
            case 37: {
              successCriteriaCase_ = 4;
              successCriteria_ = input.readFloat();
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return nebulaidl.plugins.ArrayJobOuterClass.internal_static_nebulaidl_plugins_ArrayJob_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return nebulaidl.plugins.ArrayJobOuterClass.internal_static_nebulaidl_plugins_ArrayJob_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.class, nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.Builder.class);
    }

    private int successCriteriaCase_ = 0;
    private java.lang.Object successCriteria_;
    public enum SuccessCriteriaCase
        implements com.google.protobuf.Internal.EnumLite {
      MIN_SUCCESSES(3),
      MIN_SUCCESS_RATIO(4),
      SUCCESSCRITERIA_NOT_SET(0);
      private final int value;
      private SuccessCriteriaCase(int value) {
        this.value = value;
      }
      /**
       * @deprecated Use {@link #forNumber(int)} instead.
       */
      @java.lang.Deprecated
      public static SuccessCriteriaCase valueOf(int value) {
        return forNumber(value);
      }

      public static SuccessCriteriaCase forNumber(int value) {
        switch (value) {
          case 3: return MIN_SUCCESSES;
          case 4: return MIN_SUCCESS_RATIO;
          case 0: return SUCCESSCRITERIA_NOT_SET;
          default: return null;
        }
      }
      public int getNumber() {
        return this.value;
      }
    };

    public SuccessCriteriaCase
    getSuccessCriteriaCase() {
      return SuccessCriteriaCase.forNumber(
          successCriteriaCase_);
    }

    public static final int PARALLELISM_FIELD_NUMBER = 1;
    private long parallelism_;
    /**
     * <pre>
     * Defines the maximum number of instances to bring up concurrently at any given point. Note that this is an
     * optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
     * running instances might be more. This has to be a positive number if assigned. Default value is size.
     * </pre>
     *
     * <code>int64 parallelism = 1;</code>
     */
    public long getParallelism() {
      return parallelism_;
    }

    public static final int SIZE_FIELD_NUMBER = 2;
    private long size_;
    /**
     * <pre>
     * Defines the number of instances to launch at most. This number should match the size of the input if the job
     * requires processing of all input data. This has to be a positive number.
     * In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
     * </pre>
     *
     * <code>int64 size = 2;</code>
     */
    public long getSize() {
      return size_;
    }

    public static final int MIN_SUCCESSES_FIELD_NUMBER = 3;
    /**
     * <pre>
     * An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
     * the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
     * assigned. Default value is size (if specified).
     * </pre>
     *
     * <code>int64 min_successes = 3;</code>
     */
    public long getMinSuccesses() {
      if (successCriteriaCase_ == 3) {
        return (java.lang.Long) successCriteria_;
      }
      return 0L;
    }

    public static final int MIN_SUCCESS_RATIO_FIELD_NUMBER = 4;
    /**
     * <pre>
     * If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
     * job can be marked successful.
     * </pre>
     *
     * <code>float min_success_ratio = 4;</code>
     */
    public float getMinSuccessRatio() {
      if (successCriteriaCase_ == 4) {
        return (java.lang.Float) successCriteria_;
      }
      return 0F;
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (parallelism_ != 0L) {
        output.writeInt64(1, parallelism_);
      }
      if (size_ != 0L) {
        output.writeInt64(2, size_);
      }
      if (successCriteriaCase_ == 3) {
        output.writeInt64(
            3, (long)((java.lang.Long) successCriteria_));
      }
      if (successCriteriaCase_ == 4) {
        output.writeFloat(
            4, (float)((java.lang.Float) successCriteria_));
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (parallelism_ != 0L) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt64Size(1, parallelism_);
      }
      if (size_ != 0L) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt64Size(2, size_);
      }
      if (successCriteriaCase_ == 3) {
        size += com.google.protobuf.CodedOutputStream
          .computeInt64Size(
              3, (long)((java.lang.Long) successCriteria_));
      }
      if (successCriteriaCase_ == 4) {
        size += com.google.protobuf.CodedOutputStream
          .computeFloatSize(
              4, (float)((java.lang.Float) successCriteria_));
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof nebulaidl.plugins.ArrayJobOuterClass.ArrayJob)) {
        return super.equals(obj);
      }
      nebulaidl.plugins.ArrayJobOuterClass.ArrayJob other = (nebulaidl.plugins.ArrayJobOuterClass.ArrayJob) obj;

      if (getParallelism()
          != other.getParallelism()) return false;
      if (getSize()
          != other.getSize()) return false;
      if (!getSuccessCriteriaCase().equals(other.getSuccessCriteriaCase())) return false;
      switch (successCriteriaCase_) {
        case 3:
          if (getMinSuccesses()
              != other.getMinSuccesses()) return false;
          break;
        case 4:
          if (java.lang.Float.floatToIntBits(getMinSuccessRatio())
              != java.lang.Float.floatToIntBits(
                  other.getMinSuccessRatio())) return false;
          break;
        case 0:
        default:
      }
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + PARALLELISM_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
          getParallelism());
      hash = (37 * hash) + SIZE_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
          getSize());
      switch (successCriteriaCase_) {
        case 3:
          hash = (37 * hash) + MIN_SUCCESSES_FIELD_NUMBER;
          hash = (53 * hash) + com.google.protobuf.Internal.hashLong(
              getMinSuccesses());
          break;
        case 4:
          hash = (37 * hash) + MIN_SUCCESS_RATIO_FIELD_NUMBER;
          hash = (53 * hash) + java.lang.Float.floatToIntBits(
              getMinSuccessRatio());
          break;
        case 0:
        default:
      }
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(nebulaidl.plugins.ArrayJobOuterClass.ArrayJob prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * <pre>
     * Describes a job that can process independent pieces of data concurrently. Multiple copies of the runnable component
     * will be executed concurrently.
     * </pre>
     *
     * Protobuf type {@code nebulaidl.plugins.ArrayJob}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:nebulaidl.plugins.ArrayJob)
        nebulaidl.plugins.ArrayJobOuterClass.ArrayJobOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return nebulaidl.plugins.ArrayJobOuterClass.internal_static_nebulaidl_plugins_ArrayJob_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return nebulaidl.plugins.ArrayJobOuterClass.internal_static_nebulaidl_plugins_ArrayJob_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.class, nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.Builder.class);
      }

      // Construct using nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        parallelism_ = 0L;

        size_ = 0L;

        successCriteriaCase_ = 0;
        successCriteria_ = null;
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return nebulaidl.plugins.ArrayJobOuterClass.internal_static_nebulaidl_plugins_ArrayJob_descriptor;
      }

      @java.lang.Override
      public nebulaidl.plugins.ArrayJobOuterClass.ArrayJob getDefaultInstanceForType() {
        return nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.getDefaultInstance();
      }

      @java.lang.Override
      public nebulaidl.plugins.ArrayJobOuterClass.ArrayJob build() {
        nebulaidl.plugins.ArrayJobOuterClass.ArrayJob result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public nebulaidl.plugins.ArrayJobOuterClass.ArrayJob buildPartial() {
        nebulaidl.plugins.ArrayJobOuterClass.ArrayJob result = new nebulaidl.plugins.ArrayJobOuterClass.ArrayJob(this);
        result.parallelism_ = parallelism_;
        result.size_ = size_;
        if (successCriteriaCase_ == 3) {
          result.successCriteria_ = successCriteria_;
        }
        if (successCriteriaCase_ == 4) {
          result.successCriteria_ = successCriteria_;
        }
        result.successCriteriaCase_ = successCriteriaCase_;
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof nebulaidl.plugins.ArrayJobOuterClass.ArrayJob) {
          return mergeFrom((nebulaidl.plugins.ArrayJobOuterClass.ArrayJob)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(nebulaidl.plugins.ArrayJobOuterClass.ArrayJob other) {
        if (other == nebulaidl.plugins.ArrayJobOuterClass.ArrayJob.getDefaultInstance()) return this;
        if (other.getParallelism() != 0L) {
          setParallelism(other.getParallelism());
        }
        if (other.getSize() != 0L) {
          setSize(other.getSize());
        }
        switch (other.getSuccessCriteriaCase()) {
          case MIN_SUCCESSES: {
            setMinSuccesses(other.getMinSuccesses());
            break;
          }
          case MIN_SUCCESS_RATIO: {
            setMinSuccessRatio(other.getMinSuccessRatio());
            break;
          }
          case SUCCESSCRITERIA_NOT_SET: {
            break;
          }
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        nebulaidl.plugins.ArrayJobOuterClass.ArrayJob parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (nebulaidl.plugins.ArrayJobOuterClass.ArrayJob) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }
      private int successCriteriaCase_ = 0;
      private java.lang.Object successCriteria_;
      public SuccessCriteriaCase
          getSuccessCriteriaCase() {
        return SuccessCriteriaCase.forNumber(
            successCriteriaCase_);
      }

      public Builder clearSuccessCriteria() {
        successCriteriaCase_ = 0;
        successCriteria_ = null;
        onChanged();
        return this;
      }


      private long parallelism_ ;
      /**
       * <pre>
       * Defines the maximum number of instances to bring up concurrently at any given point. Note that this is an
       * optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
       * running instances might be more. This has to be a positive number if assigned. Default value is size.
       * </pre>
       *
       * <code>int64 parallelism = 1;</code>
       */
      public long getParallelism() {
        return parallelism_;
      }
      /**
       * <pre>
       * Defines the maximum number of instances to bring up concurrently at any given point. Note that this is an
       * optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
       * running instances might be more. This has to be a positive number if assigned. Default value is size.
       * </pre>
       *
       * <code>int64 parallelism = 1;</code>
       */
      public Builder setParallelism(long value) {
        
        parallelism_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * Defines the maximum number of instances to bring up concurrently at any given point. Note that this is an
       * optimistic restriction and that, due to network partitioning or other failures, the actual number of currently
       * running instances might be more. This has to be a positive number if assigned. Default value is size.
       * </pre>
       *
       * <code>int64 parallelism = 1;</code>
       */
      public Builder clearParallelism() {
        
        parallelism_ = 0L;
        onChanged();
        return this;
      }

      private long size_ ;
      /**
       * <pre>
       * Defines the number of instances to launch at most. This number should match the size of the input if the job
       * requires processing of all input data. This has to be a positive number.
       * In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
       * </pre>
       *
       * <code>int64 size = 2;</code>
       */
      public long getSize() {
        return size_;
      }
      /**
       * <pre>
       * Defines the number of instances to launch at most. This number should match the size of the input if the job
       * requires processing of all input data. This has to be a positive number.
       * In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
       * </pre>
       *
       * <code>int64 size = 2;</code>
       */
      public Builder setSize(long value) {
        
        size_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * Defines the number of instances to launch at most. This number should match the size of the input if the job
       * requires processing of all input data. This has to be a positive number.
       * In the case this is not defined, the back-end will determine the size at run-time by reading the inputs.
       * </pre>
       *
       * <code>int64 size = 2;</code>
       */
      public Builder clearSize() {
        
        size_ = 0L;
        onChanged();
        return this;
      }

      /**
       * <pre>
       * An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
       * the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
       * assigned. Default value is size (if specified).
       * </pre>
       *
       * <code>int64 min_successes = 3;</code>
       */
      public long getMinSuccesses() {
        if (successCriteriaCase_ == 3) {
          return (java.lang.Long) successCriteria_;
        }
        return 0L;
      }
      /**
       * <pre>
       * An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
       * the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
       * assigned. Default value is size (if specified).
       * </pre>
       *
       * <code>int64 min_successes = 3;</code>
       */
      public Builder setMinSuccesses(long value) {
        successCriteriaCase_ = 3;
        successCriteria_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * An absolute number of the minimum number of successful completions of subtasks. As soon as this criteria is met,
       * the array job will be marked as successful and outputs will be computed. This has to be a non-negative number if
       * assigned. Default value is size (if specified).
       * </pre>
       *
       * <code>int64 min_successes = 3;</code>
       */
      public Builder clearMinSuccesses() {
        if (successCriteriaCase_ == 3) {
          successCriteriaCase_ = 0;
          successCriteria_ = null;
          onChanged();
        }
        return this;
      }

      /**
       * <pre>
       * If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
       * job can be marked successful.
       * </pre>
       *
       * <code>float min_success_ratio = 4;</code>
       */
      public float getMinSuccessRatio() {
        if (successCriteriaCase_ == 4) {
          return (java.lang.Float) successCriteria_;
        }
        return 0F;
      }
      /**
       * <pre>
       * If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
       * job can be marked successful.
       * </pre>
       *
       * <code>float min_success_ratio = 4;</code>
       */
      public Builder setMinSuccessRatio(float value) {
        successCriteriaCase_ = 4;
        successCriteria_ = value;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * If the array job size is not known beforehand, the min_success_ratio can instead be used to determine when an array
       * job can be marked successful.
       * </pre>
       *
       * <code>float min_success_ratio = 4;</code>
       */
      public Builder clearMinSuccessRatio() {
        if (successCriteriaCase_ == 4) {
          successCriteriaCase_ = 0;
          successCriteria_ = null;
          onChanged();
        }
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:nebulaidl.plugins.ArrayJob)
    }

    // @@protoc_insertion_point(class_scope:nebulaidl.plugins.ArrayJob)
    private static final nebulaidl.plugins.ArrayJobOuterClass.ArrayJob DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new nebulaidl.plugins.ArrayJobOuterClass.ArrayJob();
    }

    public static nebulaidl.plugins.ArrayJobOuterClass.ArrayJob getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<ArrayJob>
        PARSER = new com.google.protobuf.AbstractParser<ArrayJob>() {
      @java.lang.Override
      public ArrayJob parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new ArrayJob(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<ArrayJob> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<ArrayJob> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public nebulaidl.plugins.ArrayJobOuterClass.ArrayJob getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_nebulaidl_plugins_ArrayJob_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_nebulaidl_plugins_ArrayJob_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n!nebulaidl/plugins/array_job.proto\022\021neb" +
      "ulaidl.plugins\"w\n\010ArrayJob\022\023\n\013parallelis" +
      "m\030\001 \001(\003\022\014\n\004size\030\002 \001(\003\022\027\n\rmin_successes\030\003" +
      " \001(\003H\000\022\033\n\021min_success_ratio\030\004 \001(\002H\000B\022\n\020s" +
      "uccess_criteriaB?Z=github.com/nebulaclou" +
      "ds/nebulaidl/gen/pb-go/nebulaidl/plugins" +
      "b\006proto3"
    };
    com.google.protobuf.Descriptors.FileDescriptor.InternalDescriptorAssigner assigner =
        new com.google.protobuf.Descriptors.FileDescriptor.    InternalDescriptorAssigner() {
          public com.google.protobuf.ExtensionRegistry assignDescriptors(
              com.google.protobuf.Descriptors.FileDescriptor root) {
            descriptor = root;
            return null;
          }
        };
    com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        }, assigner);
    internal_static_nebulaidl_plugins_ArrayJob_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_nebulaidl_plugins_ArrayJob_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_nebulaidl_plugins_ArrayJob_descriptor,
        new java.lang.String[] { "Parallelism", "Size", "MinSuccesses", "MinSuccessRatio", "SuccessCriteria", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
