// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/waitable.proto

package nebulaidl.plugins;

public final class WaitableOuterClass {
  private WaitableOuterClass() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface WaitableOrBuilder extends
      // @@protoc_insertion_point(interface_extends:nebulaidl.plugins.Waitable)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    boolean hasWfExecId();
    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier getWfExecId();
    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder getWfExecIdOrBuilder();

    /**
     * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
     */
    int getPhaseValue();
    /**
     * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
     */
    nebulaidl.core.Execution.WorkflowExecution.Phase getPhase();

    /**
     * <code>string workflow_id = 3;</code>
     */
    java.lang.String getWorkflowId();
    /**
     * <code>string workflow_id = 3;</code>
     */
    com.google.protobuf.ByteString
        getWorkflowIdBytes();
  }
  /**
   * <pre>
   * Represents an Execution that was launched and could be waited on.
   * </pre>
   *
   * Protobuf type {@code nebulaidl.plugins.Waitable}
   */
  public  static final class Waitable extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:nebulaidl.plugins.Waitable)
      WaitableOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use Waitable.newBuilder() to construct.
    private Waitable(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private Waitable() {
      phase_ = 0;
      workflowId_ = "";
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private Waitable(
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
            case 10: {
              nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder subBuilder = null;
              if (wfExecId_ != null) {
                subBuilder = wfExecId_.toBuilder();
              }
              wfExecId_ = input.readMessage(nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.parser(), extensionRegistry);
              if (subBuilder != null) {
                subBuilder.mergeFrom(wfExecId_);
                wfExecId_ = subBuilder.buildPartial();
              }

              break;
            }
            case 16: {
              int rawValue = input.readEnum();

              phase_ = rawValue;
              break;
            }
            case 26: {
              java.lang.String s = input.readStringRequireUtf8();

              workflowId_ = s;
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
      return nebulaidl.plugins.WaitableOuterClass.internal_static_nebulaidl_plugins_Waitable_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return nebulaidl.plugins.WaitableOuterClass.internal_static_nebulaidl_plugins_Waitable_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              nebulaidl.plugins.WaitableOuterClass.Waitable.class, nebulaidl.plugins.WaitableOuterClass.Waitable.Builder.class);
    }

    public static final int WF_EXEC_ID_FIELD_NUMBER = 1;
    private nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier wfExecId_;
    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    public boolean hasWfExecId() {
      return wfExecId_ != null;
    }
    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    public nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier getWfExecId() {
      return wfExecId_ == null ? nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.getDefaultInstance() : wfExecId_;
    }
    /**
     * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
     */
    public nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder getWfExecIdOrBuilder() {
      return getWfExecId();
    }

    public static final int PHASE_FIELD_NUMBER = 2;
    private int phase_;
    /**
     * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
     */
    public int getPhaseValue() {
      return phase_;
    }
    /**
     * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
     */
    public nebulaidl.core.Execution.WorkflowExecution.Phase getPhase() {
      @SuppressWarnings("deprecation")
      nebulaidl.core.Execution.WorkflowExecution.Phase result = nebulaidl.core.Execution.WorkflowExecution.Phase.valueOf(phase_);
      return result == null ? nebulaidl.core.Execution.WorkflowExecution.Phase.UNRECOGNIZED : result;
    }

    public static final int WORKFLOW_ID_FIELD_NUMBER = 3;
    private volatile java.lang.Object workflowId_;
    /**
     * <code>string workflow_id = 3;</code>
     */
    public java.lang.String getWorkflowId() {
      java.lang.Object ref = workflowId_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        workflowId_ = s;
        return s;
      }
    }
    /**
     * <code>string workflow_id = 3;</code>
     */
    public com.google.protobuf.ByteString
        getWorkflowIdBytes() {
      java.lang.Object ref = workflowId_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        workflowId_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
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
      if (wfExecId_ != null) {
        output.writeMessage(1, getWfExecId());
      }
      if (phase_ != nebulaidl.core.Execution.WorkflowExecution.Phase.UNDEFINED.getNumber()) {
        output.writeEnum(2, phase_);
      }
      if (!getWorkflowIdBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 3, workflowId_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (wfExecId_ != null) {
        size += com.google.protobuf.CodedOutputStream
          .computeMessageSize(1, getWfExecId());
      }
      if (phase_ != nebulaidl.core.Execution.WorkflowExecution.Phase.UNDEFINED.getNumber()) {
        size += com.google.protobuf.CodedOutputStream
          .computeEnumSize(2, phase_);
      }
      if (!getWorkflowIdBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(3, workflowId_);
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
      if (!(obj instanceof nebulaidl.plugins.WaitableOuterClass.Waitable)) {
        return super.equals(obj);
      }
      nebulaidl.plugins.WaitableOuterClass.Waitable other = (nebulaidl.plugins.WaitableOuterClass.Waitable) obj;

      if (hasWfExecId() != other.hasWfExecId()) return false;
      if (hasWfExecId()) {
        if (!getWfExecId()
            .equals(other.getWfExecId())) return false;
      }
      if (phase_ != other.phase_) return false;
      if (!getWorkflowId()
          .equals(other.getWorkflowId())) return false;
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
      if (hasWfExecId()) {
        hash = (37 * hash) + WF_EXEC_ID_FIELD_NUMBER;
        hash = (53 * hash) + getWfExecId().hashCode();
      }
      hash = (37 * hash) + PHASE_FIELD_NUMBER;
      hash = (53 * hash) + phase_;
      hash = (37 * hash) + WORKFLOW_ID_FIELD_NUMBER;
      hash = (53 * hash) + getWorkflowId().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static nebulaidl.plugins.WaitableOuterClass.Waitable parseFrom(
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
    public static Builder newBuilder(nebulaidl.plugins.WaitableOuterClass.Waitable prototype) {
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
     * Represents an Execution that was launched and could be waited on.
     * </pre>
     *
     * Protobuf type {@code nebulaidl.plugins.Waitable}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:nebulaidl.plugins.Waitable)
        nebulaidl.plugins.WaitableOuterClass.WaitableOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return nebulaidl.plugins.WaitableOuterClass.internal_static_nebulaidl_plugins_Waitable_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return nebulaidl.plugins.WaitableOuterClass.internal_static_nebulaidl_plugins_Waitable_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                nebulaidl.plugins.WaitableOuterClass.Waitable.class, nebulaidl.plugins.WaitableOuterClass.Waitable.Builder.class);
      }

      // Construct using nebulaidl.plugins.WaitableOuterClass.Waitable.newBuilder()
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
        if (wfExecIdBuilder_ == null) {
          wfExecId_ = null;
        } else {
          wfExecId_ = null;
          wfExecIdBuilder_ = null;
        }
        phase_ = 0;

        workflowId_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return nebulaidl.plugins.WaitableOuterClass.internal_static_nebulaidl_plugins_Waitable_descriptor;
      }

      @java.lang.Override
      public nebulaidl.plugins.WaitableOuterClass.Waitable getDefaultInstanceForType() {
        return nebulaidl.plugins.WaitableOuterClass.Waitable.getDefaultInstance();
      }

      @java.lang.Override
      public nebulaidl.plugins.WaitableOuterClass.Waitable build() {
        nebulaidl.plugins.WaitableOuterClass.Waitable result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public nebulaidl.plugins.WaitableOuterClass.Waitable buildPartial() {
        nebulaidl.plugins.WaitableOuterClass.Waitable result = new nebulaidl.plugins.WaitableOuterClass.Waitable(this);
        if (wfExecIdBuilder_ == null) {
          result.wfExecId_ = wfExecId_;
        } else {
          result.wfExecId_ = wfExecIdBuilder_.build();
        }
        result.phase_ = phase_;
        result.workflowId_ = workflowId_;
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
        if (other instanceof nebulaidl.plugins.WaitableOuterClass.Waitable) {
          return mergeFrom((nebulaidl.plugins.WaitableOuterClass.Waitable)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(nebulaidl.plugins.WaitableOuterClass.Waitable other) {
        if (other == nebulaidl.plugins.WaitableOuterClass.Waitable.getDefaultInstance()) return this;
        if (other.hasWfExecId()) {
          mergeWfExecId(other.getWfExecId());
        }
        if (other.phase_ != 0) {
          setPhaseValue(other.getPhaseValue());
        }
        if (!other.getWorkflowId().isEmpty()) {
          workflowId_ = other.workflowId_;
          onChanged();
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
        nebulaidl.plugins.WaitableOuterClass.Waitable parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (nebulaidl.plugins.WaitableOuterClass.Waitable) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier wfExecId_;
      private com.google.protobuf.SingleFieldBuilderV3<
          nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder> wfExecIdBuilder_;
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public boolean hasWfExecId() {
        return wfExecIdBuilder_ != null || wfExecId_ != null;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier getWfExecId() {
        if (wfExecIdBuilder_ == null) {
          return wfExecId_ == null ? nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.getDefaultInstance() : wfExecId_;
        } else {
          return wfExecIdBuilder_.getMessage();
        }
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public Builder setWfExecId(nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier value) {
        if (wfExecIdBuilder_ == null) {
          if (value == null) {
            throw new NullPointerException();
          }
          wfExecId_ = value;
          onChanged();
        } else {
          wfExecIdBuilder_.setMessage(value);
        }

        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public Builder setWfExecId(
          nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder builderForValue) {
        if (wfExecIdBuilder_ == null) {
          wfExecId_ = builderForValue.build();
          onChanged();
        } else {
          wfExecIdBuilder_.setMessage(builderForValue.build());
        }

        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public Builder mergeWfExecId(nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier value) {
        if (wfExecIdBuilder_ == null) {
          if (wfExecId_ != null) {
            wfExecId_ =
              nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.newBuilder(wfExecId_).mergeFrom(value).buildPartial();
          } else {
            wfExecId_ = value;
          }
          onChanged();
        } else {
          wfExecIdBuilder_.mergeFrom(value);
        }

        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public Builder clearWfExecId() {
        if (wfExecIdBuilder_ == null) {
          wfExecId_ = null;
          onChanged();
        } else {
          wfExecId_ = null;
          wfExecIdBuilder_ = null;
        }

        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder getWfExecIdBuilder() {
        
        onChanged();
        return getWfExecIdFieldBuilder().getBuilder();
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      public nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder getWfExecIdOrBuilder() {
        if (wfExecIdBuilder_ != null) {
          return wfExecIdBuilder_.getMessageOrBuilder();
        } else {
          return wfExecId_ == null ?
              nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.getDefaultInstance() : wfExecId_;
        }
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;</code>
       */
      private com.google.protobuf.SingleFieldBuilderV3<
          nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder>
          getWfExecIdFieldBuilder() {
        if (wfExecIdBuilder_ == null) {
          wfExecIdBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
              nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifier.Builder, nebulaidl.core.IdentifierOuterClass.WorkflowExecutionIdentifierOrBuilder>(
                  getWfExecId(),
                  getParentForChildren(),
                  isClean());
          wfExecId_ = null;
        }
        return wfExecIdBuilder_;
      }

      private int phase_ = 0;
      /**
       * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
       */
      public int getPhaseValue() {
        return phase_;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
       */
      public Builder setPhaseValue(int value) {
        phase_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
       */
      public nebulaidl.core.Execution.WorkflowExecution.Phase getPhase() {
        @SuppressWarnings("deprecation")
        nebulaidl.core.Execution.WorkflowExecution.Phase result = nebulaidl.core.Execution.WorkflowExecution.Phase.valueOf(phase_);
        return result == null ? nebulaidl.core.Execution.WorkflowExecution.Phase.UNRECOGNIZED : result;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
       */
      public Builder setPhase(nebulaidl.core.Execution.WorkflowExecution.Phase value) {
        if (value == null) {
          throw new NullPointerException();
        }
        
        phase_ = value.getNumber();
        onChanged();
        return this;
      }
      /**
       * <code>.nebulaidl.core.WorkflowExecution.Phase phase = 2;</code>
       */
      public Builder clearPhase() {
        
        phase_ = 0;
        onChanged();
        return this;
      }

      private java.lang.Object workflowId_ = "";
      /**
       * <code>string workflow_id = 3;</code>
       */
      public java.lang.String getWorkflowId() {
        java.lang.Object ref = workflowId_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          workflowId_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string workflow_id = 3;</code>
       */
      public com.google.protobuf.ByteString
          getWorkflowIdBytes() {
        java.lang.Object ref = workflowId_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          workflowId_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string workflow_id = 3;</code>
       */
      public Builder setWorkflowId(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        workflowId_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string workflow_id = 3;</code>
       */
      public Builder clearWorkflowId() {
        
        workflowId_ = getDefaultInstance().getWorkflowId();
        onChanged();
        return this;
      }
      /**
       * <code>string workflow_id = 3;</code>
       */
      public Builder setWorkflowIdBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        workflowId_ = value;
        onChanged();
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


      // @@protoc_insertion_point(builder_scope:nebulaidl.plugins.Waitable)
    }

    // @@protoc_insertion_point(class_scope:nebulaidl.plugins.Waitable)
    private static final nebulaidl.plugins.WaitableOuterClass.Waitable DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new nebulaidl.plugins.WaitableOuterClass.Waitable();
    }

    public static nebulaidl.plugins.WaitableOuterClass.Waitable getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<Waitable>
        PARSER = new com.google.protobuf.AbstractParser<Waitable>() {
      @java.lang.Override
      public Waitable parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new Waitable(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<Waitable> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<Waitable> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public nebulaidl.plugins.WaitableOuterClass.Waitable getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_nebulaidl_plugins_Waitable_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_nebulaidl_plugins_Waitable_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\037nebulaidl/plugins/waitable.proto\022\020nebula" +
      "idl.plugins\032\035nebulaidl/core/execution.pro" +
      "to\032\036nebulaidl/core/identifier.proto\"\226\001\n\010W" +
      "aitable\022>\n\nwf_exec_id\030\001 \001(\0132*.nebulaidl.c" +
      "ore.WorkflowExecutionIdentifier\0225\n\005phase" +
      "\030\002 \001(\0162&.nebulaidl.core.WorkflowExecution" +
      ".Phase\022\023\n\013workflow_id\030\003 \001(\tB9Z7github.co" +
      "m/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/p" +
      "luginsb\006proto3"
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
          nebulaidl.core.Execution.getDescriptor(),
          nebulaidl.core.IdentifierOuterClass.getDescriptor(),
        }, assigner);
    internal_static_nebulaidl_plugins_Waitable_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_nebulaidl_plugins_Waitable_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_nebulaidl_plugins_Waitable_descriptor,
        new java.lang.String[] { "WfExecId", "Phase", "WorkflowId", });
    nebulaidl.core.Execution.getDescriptor();
    nebulaidl.core.IdentifierOuterClass.getDescriptor();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
