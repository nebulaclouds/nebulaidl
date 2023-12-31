# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/event/event.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2
from nebulaidl.core import compiler_pb2 as nebulaidl_dot_core_dot_compiler__pb2
from nebulaidl.core import execution_pb2 as nebulaidl_dot_core_dot_execution__pb2
from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2
from nebulaidl.core import catalog_pb2 as nebulaidl_dot_core_dot_catalog__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1bnebulaidl/event/event.proto\x12\x0fnebulaidl.event\x1a\x1dnebulaidl/core/literals.proto\x1a\x1dnebulaidl/core/compiler.proto\x1a\x1enebulaidl/core/execution.proto\x1a\x1fnebulaidl/core/identifier.proto\x1a\x1cnebulaidl/core/catalog.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\"\xae\x03\n\x16WorkflowExecutionEvent\x12N\n\x0c\x65xecution_id\x18\x01 \x01(\x0b\x32+.nebulaidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId\x12\x1f\n\x0bproducer_id\x18\x02 \x01(\tR\nproducerId\x12=\n\x05phase\x18\x03 \x01(\x0e\x32\'.nebulaidl.core.WorkflowExecution.PhaseR\x05phase\x12;\n\x0boccurred_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\noccurredAt\x12\x1f\n\noutput_uri\x18\x05 \x01(\tH\x00R\toutputUri\x12\x36\n\x05\x65rror\x18\x06 \x01(\x0b\x32\x1e.nebulaidl.core.ExecutionErrorH\x00R\x05\x65rror\x12=\n\x0boutput_data\x18\x07 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x00R\noutputDataB\x0f\n\routput_result\"\x98\t\n\x12NodeExecutionEvent\x12\x37\n\x02id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierR\x02id\x12\x1f\n\x0bproducer_id\x18\x02 \x01(\tR\nproducerId\x12\x39\n\x05phase\x18\x03 \x01(\x0e\x32#.nebulaidl.core.NodeExecution.PhaseR\x05phase\x12;\n\x0boccurred_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\noccurredAt\x12\x1d\n\tinput_uri\x18\x05 \x01(\tH\x00R\x08inputUri\x12;\n\ninput_data\x18\x14 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x00R\tinputData\x12\x1f\n\noutput_uri\x18\x06 \x01(\tH\x01R\toutputUri\x12\x36\n\x05\x65rror\x18\x07 \x01(\x0b\x32\x1e.nebulaidl.core.ExecutionErrorH\x01R\x05\x65rror\x12=\n\x0boutput_data\x18\x0f \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x01R\noutputData\x12]\n\x16workflow_node_metadata\x18\x08 \x01(\x0b\x32%.nebulaidl.event.WorkflowNodeMetadataH\x02R\x14workflowNodeMetadata\x12Q\n\x12task_node_metadata\x18\x0e \x01(\x0b\x32!.nebulaidl.event.TaskNodeMetadataH\x02R\x10taskNodeMetadata\x12^\n\x14parent_task_metadata\x18\t \x01(\x0b\x32,.nebulaidl.event.ParentTaskExecutionMetadataR\x12parentTaskMetadata\x12^\n\x14parent_node_metadata\x18\n \x01(\x0b\x32,.nebulaidl.event.ParentNodeExecutionMetadataR\x12parentNodeMetadata\x12\x1f\n\x0bretry_group\x18\x0b \x01(\tR\nretryGroup\x12 \n\x0cspec_node_id\x18\x0c \x01(\tR\nspecNodeId\x12\x1b\n\tnode_name\x18\r \x01(\tR\x08nodeName\x12#\n\revent_version\x18\x10 \x01(\x05R\x0c\x65ventVersion\x12\x1b\n\tis_parent\x18\x11 \x01(\x08R\x08isParent\x12\x1d\n\nis_dynamic\x18\x12 \x01(\x08R\tisDynamic\x12\x19\n\x08\x64\x65\x63k_uri\x18\x13 \x01(\tR\x07\x64\x65\x63kUri\x12;\n\x0breported_at\x18\x15 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nreportedAtB\r\n\x0binput_valueB\x0f\n\routput_resultB\x11\n\x0ftarget_metadata\"f\n\x14WorkflowNodeMetadata\x12N\n\x0c\x65xecution_id\x18\x01 \x01(\x0b\x32+.nebulaidl.core.WorkflowExecutionIdentifierR\x0b\x65xecutionId\"\xf5\x02\n\x10TaskNodeMetadata\x12\x45\n\x0c\x63\x61\x63he_status\x18\x01 \x01(\x0e\x32\".nebulaidl.core.CatalogCacheStatusR\x0b\x63\x61\x63heStatus\x12@\n\x0b\x63\x61talog_key\x18\x02 \x01(\x0b\x32\x1f.nebulaidl.core.CatalogMetadataR\ncatalogKey\x12X\n\x12reservation_status\x18\x03 \x01(\x0e\x32).nebulaidl.core.CatalogReservation.StatusR\x11reservationStatus\x12%\n\x0e\x63heckpoint_uri\x18\x04 \x01(\tR\rcheckpointUri\x12W\n\x10\x64ynamic_workflow\x18\x10 \x01(\x0b\x32,.nebulaidl.event.DynamicWorkflowNodeMetadataR\x0f\x64ynamicWorkflow\"\xd0\x01\n\x1b\x44ynamicWorkflowNodeMetadata\x12*\n\x02id\x18\x01 \x01(\x0b\x32\x1a.nebulaidl.core.IdentifierR\x02id\x12T\n\x11\x63ompiled_workflow\x18\x02 \x01(\x0b\x32\'.nebulaidl.core.CompiledWorkflowClosureR\x10\x63ompiledWorkflow\x12/\n\x14\x64ynamic_job_spec_uri\x18\x03 \x01(\tR\x11\x64ynamicJobSpecUri\"V\n\x1bParentTaskExecutionMetadata\x12\x37\n\x02id\x18\x01 \x01(\x0b\x32\'.nebulaidl.core.TaskExecutionIdentifierR\x02id\"6\n\x1bParentNodeExecutionMetadata\x12\x17\n\x07node_id\x18\x01 \x01(\tR\x06nodeId\"b\n\x0b\x45ventReason\x12\x16\n\x06reason\x18\x01 \x01(\tR\x06reason\x12;\n\x0boccurred_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\noccurredAt\"\xa0\x08\n\x12TaskExecutionEvent\x12\x33\n\x07task_id\x18\x01 \x01(\x0b\x32\x1a.nebulaidl.core.IdentifierR\x06taskId\x12`\n\x18parent_node_execution_id\x18\x02 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierR\x15parentNodeExecutionId\x12#\n\rretry_attempt\x18\x03 \x01(\rR\x0cretryAttempt\x12\x39\n\x05phase\x18\x04 \x01(\x0e\x32#.nebulaidl.core.TaskExecution.PhaseR\x05phase\x12\x1f\n\x0bproducer_id\x18\x05 \x01(\tR\nproducerId\x12+\n\x04logs\x18\x06 \x03(\x0b\x32\x17.nebulaidl.core.TaskLogR\x04logs\x12;\n\x0boccurred_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\noccurredAt\x12\x1d\n\tinput_uri\x18\x08 \x01(\tH\x00R\x08inputUri\x12;\n\ninput_data\x18\x13 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x00R\tinputData\x12\x1f\n\noutput_uri\x18\t \x01(\tH\x01R\toutputUri\x12\x36\n\x05\x65rror\x18\n \x01(\x0b\x32\x1e.nebulaidl.core.ExecutionErrorH\x01R\x05\x65rror\x12=\n\x0boutput_data\x18\x11 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x01R\noutputData\x12\x38\n\x0b\x63ustom_info\x18\x0b \x01(\x0b\x32\x17.google.protobuf.StructR\ncustomInfo\x12#\n\rphase_version\x18\x0c \x01(\rR\x0cphaseVersion\x12\x1a\n\x06reason\x18\r \x01(\tB\x02\x18\x01R\x06reason\x12\x36\n\x07reasons\x18\x15 \x03(\x0b\x32\x1c.nebulaidl.event.EventReasonR\x07reasons\x12\x1b\n\ttask_type\x18\x0e \x01(\tR\x08taskType\x12\x42\n\x08metadata\x18\x10 \x01(\x0b\x32&.nebulaidl.event.TaskExecutionMetadataR\x08metadata\x12#\n\revent_version\x18\x12 \x01(\x05R\x0c\x65ventVersion\x12;\n\x0breported_at\x18\x14 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\nreportedAtB\r\n\x0binput_valueB\x0f\n\routput_result\"\xa1\x02\n\x14\x45xternalResourceInfo\x12\x1f\n\x0b\x65xternal_id\x18\x01 \x01(\tR\nexternalId\x12\x14\n\x05index\x18\x02 \x01(\rR\x05index\x12#\n\rretry_attempt\x18\x03 \x01(\rR\x0cretryAttempt\x12\x39\n\x05phase\x18\x04 \x01(\x0e\x32#.nebulaidl.core.TaskExecution.PhaseR\x05phase\x12\x45\n\x0c\x63\x61\x63he_status\x18\x05 \x01(\x0e\x32\".nebulaidl.core.CatalogCacheStatusR\x0b\x63\x61\x63heStatus\x12+\n\x04logs\x18\x06 \x03(\x0b\x32\x17.nebulaidl.core.TaskLogR\x04logs\"[\n\x10ResourcePoolInfo\x12)\n\x10\x61llocation_token\x18\x01 \x01(\tR\x0f\x61llocationToken\x12\x1c\n\tnamespace\x18\x02 \x01(\tR\tnamespace\"\xa0\x03\n\x15TaskExecutionMetadata\x12%\n\x0egenerated_name\x18\x01 \x01(\tR\rgeneratedName\x12T\n\x12\x65xternal_resources\x18\x02 \x03(\x0b\x32%.nebulaidl.event.ExternalResourceInfoR\x11\x65xternalResources\x12O\n\x12resource_pool_info\x18\x03 \x03(\x0b\x32!.nebulaidl.event.ResourcePoolInfoR\x10resourcePoolInfo\x12+\n\x11plugin_identifier\x18\x04 \x01(\tR\x10pluginIdentifier\x12[\n\x0einstance_class\x18\x10 \x01(\x0e\x32\x34.nebulaidl.event.TaskExecutionMetadata.InstanceClassR\rinstanceClass\"/\n\rInstanceClass\x12\x0b\n\x07\x44\x45\x46\x41ULT\x10\x00\x12\x11\n\rINTERRUPTIBLE\x10\x01\x42\xbb\x01\n\x13\x63om.nebulaidl.eventB\nEventProtoP\x01Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/event\xa2\x02\x03NEX\xaa\x02\x0fNebulaidl.Event\xca\x02\x0fNebulaidl\\Event\xe2\x02\x1bNebulaidl\\Event\\GPBMetadata\xea\x02\x10Nebulaidl::Eventb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.event.event_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\023com.nebulaidl.eventB\nEventProtoP\001Z;github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/event\242\002\003NEX\252\002\017Nebulaidl.Event\312\002\017Nebulaidl\\Event\342\002\033Nebulaidl\\Event\\GPBMetadata\352\002\020Nebulaidl::Event'
  _TASKEXECUTIONEVENT.fields_by_name['reason']._options = None
  _TASKEXECUTIONEVENT.fields_by_name['reason']._serialized_options = b'\030\001'
  _globals['_WORKFLOWEXECUTIONEVENT']._serialized_start=269
  _globals['_WORKFLOWEXECUTIONEVENT']._serialized_end=699
  _globals['_NODEEXECUTIONEVENT']._serialized_start=702
  _globals['_NODEEXECUTIONEVENT']._serialized_end=1878
  _globals['_WORKFLOWNODEMETADATA']._serialized_start=1880
  _globals['_WORKFLOWNODEMETADATA']._serialized_end=1982
  _globals['_TASKNODEMETADATA']._serialized_start=1985
  _globals['_TASKNODEMETADATA']._serialized_end=2358
  _globals['_DYNAMICWORKFLOWNODEMETADATA']._serialized_start=2361
  _globals['_DYNAMICWORKFLOWNODEMETADATA']._serialized_end=2569
  _globals['_PARENTTASKEXECUTIONMETADATA']._serialized_start=2571
  _globals['_PARENTTASKEXECUTIONMETADATA']._serialized_end=2657
  _globals['_PARENTNODEEXECUTIONMETADATA']._serialized_start=2659
  _globals['_PARENTNODEEXECUTIONMETADATA']._serialized_end=2713
  _globals['_EVENTREASON']._serialized_start=2715
  _globals['_EVENTREASON']._serialized_end=2813
  _globals['_TASKEXECUTIONEVENT']._serialized_start=2816
  _globals['_TASKEXECUTIONEVENT']._serialized_end=3872
  _globals['_EXTERNALRESOURCEINFO']._serialized_start=3875
  _globals['_EXTERNALRESOURCEINFO']._serialized_end=4164
  _globals['_RESOURCEPOOLINFO']._serialized_start=4166
  _globals['_RESOURCEPOOLINFO']._serialized_end=4257
  _globals['_TASKEXECUTIONMETADATA']._serialized_start=4260
  _globals['_TASKEXECUTIONMETADATA']._serialized_end=4676
  _globals['_TASKEXECUTIONMETADATA_INSTANCECLASS']._serialized_start=4629
  _globals['_TASKEXECUTIONMETADATA_INSTANCECLASS']._serialized_end=4676
# @@protoc_insertion_point(module_scope)
