# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/identifier.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1fnebulaidl/core/identifier.proto\x12\x0enebulaidl.core\"\xaf\x01\n\nIdentifier\x12\x41\n\rresource_type\x18\x01 \x01(\x0e\x32\x1c.nebulaidl.core.ResourceTypeR\x0cresourceType\x12\x18\n\x07project\x18\x02 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x03 \x01(\tR\x06\x64omain\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\x12\x18\n\x07version\x18\x05 \x01(\tR\x07version\"c\n\x1bWorkflowExecutionIdentifier\x12\x18\n\x07project\x18\x01 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x12\n\x04name\x18\x04 \x01(\tR\x04name\"\x82\x01\n\x17NodeExecutionIdentifier\x12\x17\n\x07node_id\x18\x01 \x01(\tR\x06nodeId\x12N\n\x0c\x65xecution_id\x18\x02 \x01(\x0b\x32+.nebulaidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId\"\xc8\x01\n\x17TaskExecutionIdentifier\x12\x33\n\x07task_id\x18\x01 \x01(\x0b\x32\x1a.nebulaidl.core.IdentifierR\x06taskId\x12S\n\x11node_execution_id\x18\x02 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierR\x0fnodeExecutionId\x12#\n\rretry_attempt\x18\x03 \x01(\rR\x0cretryAttempt\"\x7f\n\x10SignalIdentifier\x12\x1b\n\tsignal_id\x18\x01 \x01(\tR\x08signalId\x12N\n\x0c\x65xecution_id\x18\x02 \x01(\x0b\x32+.nebulaidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId*U\n\x0cResourceType\x12\x0f\n\x0bUNSPECIFIED\x10\x00\x12\x08\n\x04TASK\x10\x01\x12\x0c\n\x08WORKFLOW\x10\x02\x12\x0f\n\x0bLAUNCH_PLAN\x10\x03\x12\x0b\n\x07\x44\x41TASET\x10\x04\x42\xba\x01\n\x12\x63om.nebulaidl.coreB\x0fIdentifierProtoP\x01Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03NCX\xaa\x02\x0eNebulaidl.Core\xca\x02\x0eNebulaidl\\Core\xe2\x02\x1aNebulaidl\\Core\\GPBMetadata\xea\x02\x0fNebulaidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.identifier_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.coreB\017IdentifierProtoP\001Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003NCX\252\002\016Nebulaidl.Core\312\002\016Nebulaidl\\Core\342\002\032Nebulaidl\\Core\\GPBMetadata\352\002\017Nebulaidl::Core'
  _globals['_RESOURCETYPE']._serialized_start=795
  _globals['_RESOURCETYPE']._serialized_end=880
  _globals['_IDENTIFIER']._serialized_start=52
  _globals['_IDENTIFIER']._serialized_end=227
  _globals['_WORKFLOWEXECUTIONIDENTIFIER']._serialized_start=229
  _globals['_WORKFLOWEXECUTIONIDENTIFIER']._serialized_end=328
  _globals['_NODEEXECUTIONIDENTIFIER']._serialized_start=331
  _globals['_NODEEXECUTIONIDENTIFIER']._serialized_end=461
  _globals['_TASKEXECUTIONIDENTIFIER']._serialized_start=464
  _globals['_TASKEXECUTIONIDENTIFIER']._serialized_end=664
  _globals['_SIGNALIDENTIFIER']._serialized_start=666
  _globals['_SIGNALIDENTIFIER']._serialized_end=793
# @@protoc_insertion_point(module_scope)
