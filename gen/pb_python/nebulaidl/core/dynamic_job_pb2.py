# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/dynamic_job.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import tasks_pb2 as nebulaidl_dot_core_dot_tasks__pb2
from nebulaidl.core import workflow_pb2 as nebulaidl_dot_core_dot_workflow__pb2
from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1f\x66lyteidl/core/dynamic_job.proto\x12\rnebulaidl.core\x1a\x19\x66lyteidl/core/tasks.proto\x1a\x1c\x66lyteidl/core/workflow.proto\x1a\x1c\x66lyteidl/core/literals.proto\"\x8a\x02\n\x0e\x44ynamicJobSpec\x12)\n\x05nodes\x18\x01 \x03(\x0b\x32\x13.nebulaidl.core.NodeR\x05nodes\x12#\n\rmin_successes\x18\x02 \x01(\x03R\x0cminSuccesses\x12\x30\n\x07outputs\x18\x03 \x03(\x0b\x32\x16.nebulaidl.core.BindingR\x07outputs\x12\x31\n\x05tasks\x18\x04 \x03(\x0b\x32\x1b.nebulaidl.core.TaskTemplateR\x05tasks\x12\x43\n\x0csubworkflows\x18\x05 \x03(\x0b\x32\x1f.nebulaidl.core.WorkflowTemplateR\x0csubworkflowsB\xaf\x01\n\x11\x63om.nebulaidl.coreB\x0f\x44ynamicJobProtoP\x01Z4github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03\x46\x43X\xaa\x02\rNebulaidl.Core\xca\x02\rNebulaidl\\Core\xe2\x02\x19\x46lyteidl\\Core\\GPBMetadata\xea\x02\x0e\x46lyteidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.dynamic_job_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\021com.nebulaidl.coreB\017DynamicJobProtoP\001Z4github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003FCX\252\002\rNebulaidl.Core\312\002\rNebulaidl\\Core\342\002\031Nebulaidl\\Core\\GPBMetadata\352\002\016Nebulaidl::Core'
  _globals['_DYNAMICJOBSPEC']._serialized_start=138
  _globals['_DYNAMICJOBSPEC']._serialized_end=404
# @@protoc_insertion_point(module_scope)
