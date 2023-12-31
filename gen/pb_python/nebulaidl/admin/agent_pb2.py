# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/admin/agent.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2
from nebulaidl.core import tasks_pb2 as nebulaidl_dot_core_dot_tasks__pb2
from nebulaidl.core import interface_pb2 as nebulaidl_dot_core_dot_interface__pb2
from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bnebulaidl/admin/agent.proto\x12\x0fnebulaidl.admin\x1a\x1dnebulaidl/core/literals.proto\x1a\x1anebulaidl/core/tasks.proto\x1a\x1enebulaidl/core/interface.proto\x1a\x1fnebulaidl/core/identifier.proto\"\x9c\x05\n\x15TaskExecutionMetadata\x12S\n\x11task_execution_id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierR\x0ftaskExecutionId\x12\x1c\n\tnamespace\x18\x02 \x01(\tR\tnamespace\x12J\n\x06labels\x18\x03 \x03(\x0b\x32\x32.nebulaidl.admin.TaskExecutionMetadata.LabelsEntryR\x06labels\x12Y\n\x0b\x61nnotations\x18\x04 \x03(\x0b\x32\x37.nebulaidl.admin.TaskExecutionMetadata.AnnotationsEntryR\x0b\x61nnotations\x12.\n\x13k8s_service_account\x18\x05 \x01(\tR\x11k8sServiceAccount\x12u\n\x15\x65nvironment_variables\x18\x06 \x03(\x0b\x32@.nebulaidl.admin.TaskExecutionMetadata.EnvironmentVariablesEntryR\x14\x65nvironmentVariables\x1a\x39\n\x0bLabelsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x1a>\n\x10\x41nnotationsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\x1aG\n\x19\x45nvironmentVariablesEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value:\x02\x38\x01\"\x86\x02\n\x11\x43reateTaskRequest\x12\x32\n\x06inputs\x18\x01 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapR\x06inputs\x12\x38\n\x08template\x18\x02 \x01(\x0b\x32\x1c.nebulaidl.core.TaskTemplateR\x08template\x12#\n\routput_prefix\x18\x03 \x01(\tR\x0coutputPrefix\x12^\n\x17task_execution_metadata\x18\x04 \x01(\x0b\x32&.nebulaidl.admin.TaskExecutionMetadataR\x15taskExecutionMetadata\"9\n\x12\x43reateTaskResponse\x12#\n\rresource_meta\x18\x01 \x01(\x0cR\x0cresourceMeta\"R\n\x0eGetTaskRequest\x12\x1b\n\ttask_type\x18\x01 \x01(\tR\x08taskType\x12#\n\rresource_meta\x18\x02 \x01(\x0cR\x0cresourceMeta\"H\n\x0fGetTaskResponse\x12\x35\n\x08resource\x18\x01 \x01(\x0b\x32\x19.nebulaidl.admin.ResourceR\x08resource\"n\n\x08Resource\x12,\n\x05state\x18\x01 \x01(\x0e\x32\x16.nebulaidl.admin.StateR\x05state\x12\x34\n\x07outputs\x18\x02 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapR\x07outputs\"U\n\x11\x44\x65leteTaskRequest\x12\x1b\n\ttask_type\x18\x01 \x01(\tR\x08taskType\x12#\n\rresource_meta\x18\x02 \x01(\x0cR\x0cresourceMeta\"\x14\n\x12\x44\x65leteTaskResponse*^\n\x05State\x12\x15\n\x11RETRYABLE_FAILURE\x10\x00\x12\x15\n\x11PERMANENT_FAILURE\x10\x01\x12\x0b\n\x07PENDING\x10\x02\x12\x0b\n\x07RUNNING\x10\x03\x12\r\n\tSUCCEEDED\x10\x04\x42\xbb\x01\n\x13\x63om.nebulaidl.adminB\nAgentProtoP\x01Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/admin\xa2\x02\x03NAX\xaa\x02\x0fNebulaidl.Admin\xca\x02\x0fNebulaidl\\Admin\xe2\x02\x1bNebulaidl\\Admin\\GPBMetadata\xea\x02\x10Nebulaidl::Adminb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.admin.agent_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023com.nebulaidl.adminB\nAgentProtoP\001Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/admin\242\002\003NAX\252\002\017Nebulaidl.Admin\312\002\017Nebulaidl\\Admin\342\002\033Nebulaidl\\Admin\\GPBMetadata\352\002\020Nebulaidl::Admin'
  _TASKEXECUTIONMETADATA_LABELSENTRY._options = None
  _TASKEXECUTIONMETADATA_LABELSENTRY._serialized_options = b'8\001'
  _TASKEXECUTIONMETADATA_ANNOTATIONSENTRY._options = None
  _TASKEXECUTIONMETADATA_ANNOTATIONSENTRY._serialized_options = b'8\001'
  _TASKEXECUTIONMETADATA_ENVIRONMENTVARIABLESENTRY._options = None
  _TASKEXECUTIONMETADATA_ENVIRONMENTVARIABLESENTRY._serialized_options = b'8\001'
  _globals['_STATE']._serialized_start=1546
  _globals['_STATE']._serialized_end=1640
  _globals['_TASKEXECUTIONMETADATA']._serialized_start=173
  _globals['_TASKEXECUTIONMETADATA']._serialized_end=841
  _globals['_TASKEXECUTIONMETADATA_LABELSENTRY']._serialized_start=647
  _globals['_TASKEXECUTIONMETADATA_LABELSENTRY']._serialized_end=704
  _globals['_TASKEXECUTIONMETADATA_ANNOTATIONSENTRY']._serialized_start=706
  _globals['_TASKEXECUTIONMETADATA_ANNOTATIONSENTRY']._serialized_end=768
  _globals['_TASKEXECUTIONMETADATA_ENVIRONMENTVARIABLESENTRY']._serialized_start=770
  _globals['_TASKEXECUTIONMETADATA_ENVIRONMENTVARIABLESENTRY']._serialized_end=841
  _globals['_CREATETASKREQUEST']._serialized_start=844
  _globals['_CREATETASKREQUEST']._serialized_end=1106
  _globals['_CREATETASKRESPONSE']._serialized_start=1108
  _globals['_CREATETASKRESPONSE']._serialized_end=1165
  _globals['_GETTASKREQUEST']._serialized_start=1167
  _globals['_GETTASKREQUEST']._serialized_end=1249
  _globals['_GETTASKRESPONSE']._serialized_start=1251
  _globals['_GETTASKRESPONSE']._serialized_end=1323
  _globals['_RESOURCE']._serialized_start=1325
  _globals['_RESOURCE']._serialized_end=1435
  _globals['_DELETETASKREQUEST']._serialized_start=1437
  _globals['_DELETETASKREQUEST']._serialized_end=1522
  _globals['_DELETETASKRESPONSE']._serialized_start=1524
  _globals['_DELETETASKRESPONSE']._serialized_end=1544
# @@protoc_insertion_point(module_scope)
