# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/metrics.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1cnebulaidl/core/metrics.proto\x12\x0enebulaidl.core\x1a\x1fnebulaidl/core/identifier.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xa7\x03\n\x04Span\x12\x39\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartTime\x12\x35\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x65ndTime\x12N\n\x0bworkflow_id\x18\x03 \x01(\x0b\x32+.nebulaidl.core.WorkflowExecutionIdentifierH\x00R\nworkflowId\x12\x42\n\x07node_id\x18\x04 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierH\x00R\x06nodeId\x12\x42\n\x07task_id\x18\x05 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierH\x00R\x06taskId\x12#\n\x0coperation_id\x18\x06 \x01(\tH\x00R\x0boperationId\x12*\n\x05spans\x18\x07 \x03(\x0b\x32\x14.nebulaidl.core.SpanR\x05spansB\x04\n\x02idB\xb7\x01\n\x12\x63om.nebulaidl.coreB\x0cMetricsProtoP\x01Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03NCX\xaa\x02\x0eNebulaidl.Core\xca\x02\x0eNebulaidl\\Core\xe2\x02\x1aNebulaidl\\Core\\GPBMetadata\xea\x02\x0fNebulaidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.metrics_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.coreB\014MetricsProtoP\001Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003NCX\252\002\016Nebulaidl.Core\312\002\016Nebulaidl\\Core\342\002\032Nebulaidl\\Core\\GPBMetadata\352\002\017Nebulaidl::Core'
  _globals['_SPAN']._serialized_start=115
  _globals['_SPAN']._serialized_end=538
# @@protoc_insertion_point(module_scope)
