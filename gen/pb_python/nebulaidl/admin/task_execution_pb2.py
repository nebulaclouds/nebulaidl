# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/admin/task_execution.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.admin import common_pb2 as nebulaidl_dot_admin_dot_common__pb2
from nebulaidl.core import execution_pb2 as nebulaidl_dot_core_dot_execution__pb2
from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2
from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2
from nebulaidl.event import event_pb2 as nebulaidl_dot_event_dot_event__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$nebulaidl/admin/task_execution.proto\x12\x0fnebulaidl.admin\x1a\x1cnebulaidl/admin/common.proto\x1a\x1enebulaidl/core/execution.proto\x1a\x1fnebulaidl/core/identifier.proto\x1a\x1dnebulaidl/core/literals.proto\x1a\x1bnebulaidl/event/event.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1cgoogle/protobuf/struct.proto\"R\n\x17TaskExecutionGetRequest\x12\x37\n\x02id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierR\x02id\"\xe5\x01\n\x18TaskExecutionListRequest\x12S\n\x11node_execution_id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierR\x0fnodeExecutionId\x12\x14\n\x05limit\x18\x02 \x01(\rR\x05limit\x12\x14\n\x05token\x18\x03 \x01(\tR\x05token\x12\x18\n\x07\x66ilters\x18\x04 \x01(\tR\x07\x66ilters\x12.\n\x07sort_by\x18\x05 \x01(\x0b\x32\x15.nebulaidl.admin.SortR\x06sortBy\"\xc3\x01\n\rTaskExecution\x12\x37\n\x02id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierR\x02id\x12\x1b\n\tinput_uri\x18\x02 \x01(\tR\x08inputUri\x12?\n\x07\x63losure\x18\x03 \x01(\x0b\x32%.nebulaidl.admin.TaskExecutionClosureR\x07\x63losure\x12\x1b\n\tis_parent\x18\x04 \x01(\x08R\x08isParent\"r\n\x11TaskExecutionList\x12G\n\x0ftask_executions\x18\x01 \x03(\x0b\x32\x1e.nebulaidl.admin.TaskExecutionR\x0etaskExecutions\x12\x14\n\x05token\x18\x02 \x01(\tR\x05token\"\xa2\x06\n\x14TaskExecutionClosure\x12#\n\noutput_uri\x18\x01 \x01(\tB\x02\x18\x01H\x00R\toutputUri\x12\x36\n\x05\x65rror\x18\x02 \x01(\x0b\x32\x1e.nebulaidl.core.ExecutionErrorH\x00R\x05\x65rror\x12\x41\n\x0boutput_data\x18\x0c \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapB\x02\x18\x01H\x00R\noutputData\x12\x39\n\x05phase\x18\x03 \x01(\x0e\x32#.nebulaidl.core.TaskExecution.PhaseR\x05phase\x12+\n\x04logs\x18\x04 \x03(\x0b\x32\x17.nebulaidl.core.TaskLogR\x04logs\x12\x39\n\nstarted_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tstartedAt\x12\x35\n\x08\x64uration\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x64uration\x12\x39\n\ncreated_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x39\n\nupdated_at\x18\x08 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\tupdatedAt\x12\x38\n\x0b\x63ustom_info\x18\t \x01(\x0b\x32\x17.google.protobuf.StructR\ncustomInfo\x12\x16\n\x06reason\x18\n \x01(\tR\x06reason\x12\x1b\n\ttask_type\x18\x0b \x01(\tR\x08taskType\x12\x42\n\x08metadata\x18\x10 \x01(\x0b\x32&.nebulaidl.event.TaskExecutionMetadataR\x08metadata\x12#\n\revent_version\x18\x11 \x01(\x05R\x0c\x65ventVersion\x12\x31\n\x07reasons\x18\x12 \x03(\x0b\x32\x17.nebulaidl.admin.ReasonR\x07reasonsB\x0f\n\routput_result\"_\n\x06Reason\x12;\n\x0boccurred_at\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\noccurredAt\x12\x18\n\x07message\x18\x02 \x01(\tR\x07message\"V\n\x1bTaskExecutionGetDataRequest\x12\x37\n\x02id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierR\x02id\"\xc6\x02\n\x1cTaskExecutionGetDataResponse\x12\x34\n\x06inputs\x18\x01 \x01(\x0b\x32\x18.nebulaidl.admin.UrlBlobB\x02\x18\x01R\x06inputs\x12\x36\n\x07outputs\x18\x02 \x01(\x0b\x32\x18.nebulaidl.admin.UrlBlobB\x02\x18\x01R\x07outputs\x12;\n\x0b\x66ull_inputs\x18\x03 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapR\nfullInputs\x12=\n\x0c\x66ull_outputs\x18\x04 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapR\x0b\x66ullOutputs\x12<\n\x0bnebula_urls\x18\x05 \x01(\x0b\x32\x1b.nebulaidl.admin.NebulaURLsR\nnebulaUrlsB\xc3\x01\n\x13\x63om.nebulaidl.adminB\x12TaskExecutionProtoP\x01Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/admin\xa2\x02\x03NAX\xaa\x02\x0fNebulaidl.Admin\xca\x02\x0fNebulaidl\\Admin\xe2\x02\x1bNebulaidl\\Admin\\GPBMetadata\xea\x02\x10Nebulaidl::Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.admin.task_execution_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023com.nebulaidl.adminB\022TaskExecutionProtoP\001Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/admin\242\002\003NAX\252\002\017Nebulaidl.Admin\312\002\017Nebulaidl\\Admin\342\002\033Nebulaidl\\Admin\\GPBMetadata\352\002\020Nebulaidl::Admin'
  _TASKEXECUTIONCLOSURE.fields_by_name['output_uri']._options = None
  _TASKEXECUTIONCLOSURE.fields_by_name['output_uri']._serialized_options = b'\030\001'
  _TASKEXECUTIONCLOSURE.fields_by_name['output_data']._options = None
  _TASKEXECUTIONCLOSURE.fields_by_name['output_data']._serialized_options = b'\030\001'
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._options = None
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['inputs']._serialized_options = b'\030\001'
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._options = None
  _TASKEXECUTIONGETDATARESPONSE.fields_by_name['outputs']._serialized_options = b'\030\001'
  _globals['_TASKEXECUTIONGETREQUEST']._serialized_start=307
  _globals['_TASKEXECUTIONGETREQUEST']._serialized_end=389
  _globals['_TASKEXECUTIONLISTREQUEST']._serialized_start=392
  _globals['_TASKEXECUTIONLISTREQUEST']._serialized_end=621
  _globals['_TASKEXECUTION']._serialized_start=624
  _globals['_TASKEXECUTION']._serialized_end=819
  _globals['_TASKEXECUTIONLIST']._serialized_start=821
  _globals['_TASKEXECUTIONLIST']._serialized_end=935
  _globals['_TASKEXECUTIONCLOSURE']._serialized_start=938
  _globals['_TASKEXECUTIONCLOSURE']._serialized_end=1740
  _globals['_REASON']._serialized_start=1742
  _globals['_REASON']._serialized_end=1837
  _globals['_TASKEXECUTIONGETDATAREQUEST']._serialized_start=1839
  _globals['_TASKEXECUTIONGETDATAREQUEST']._serialized_end=1925
  _globals['_TASKEXECUTIONGETDATARESPONSE']._serialized_start=1928
  _globals['_TASKEXECUTIONGETDATARESPONSE']._serialized_end=2254
# @@protoc_insertion_point(module_scope)
