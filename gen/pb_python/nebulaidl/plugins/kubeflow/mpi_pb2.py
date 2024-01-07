# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/plugins/kubeflow/mpi.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import tasks_pb2 as nebulaidl_dot_core_dot_tasks__pb2
from nebulaidl.plugins.kubeflow import common_pb2 as nebulaidl_dot_plugins_dot_kubeflow_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n$nebulaidl/plugins/kubeflow/mpi.proto\x12\x1anebulaidl.plugins.kubeflow\x1a\x1anebulaidl/core/tasks.proto\x1a\'nebulaidl/plugins/kubeflow/common.proto\"\xcc\x02\n\x1a\x44istributedMPITrainingTask\x12\x66\n\x0fworker_replicas\x18\x01 \x01(\x0b\x32=.nebulaidl.plugins.kubeflow.DistributedMPITrainingReplicaSpecR\x0eworkerReplicas\x12j\n\x11launcher_replicas\x18\x02 \x01(\x0b\x32=.nebulaidl.plugins.kubeflow.DistributedMPITrainingReplicaSpecR\x10launcherReplicas\x12\x44\n\nrun_policy\x18\x03 \x01(\x0b\x32%.nebulaidl.plugins.kubeflow.RunPolicyR\trunPolicy\x12\x14\n\x05slots\x18\x04 \x01(\x05R\x05slots\"\xfa\x01\n!DistributedMPITrainingReplicaSpec\x12\x1a\n\x08replicas\x18\x01 \x01(\x05R\x08replicas\x12\x14\n\x05image\x18\x02 \x01(\tR\x05image\x12\x37\n\tresources\x18\x03 \x01(\x0b\x32\x19.nebulaidl.core.ResourcesR\tresources\x12P\n\x0erestart_policy\x18\x04 \x01(\x0e\x32).nebulaidl.plugins.kubeflow.RestartPolicyR\rrestartPolicy\x12\x18\n\x07\x63ommand\x18\x05 \x03(\tR\x07\x63ommandB\xf3\x01\n\x1e\x63om.nebulaidl.plugins.kubeflowB\x08MpiProtoP\x01Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\xa2\x02\x03NPK\xaa\x02\x1aNebulaidl.Plugins.Kubeflow\xca\x02\x1aNebulaidl\\Plugins\\Kubeflow\xe2\x02&Nebulaidl\\Plugins\\Kubeflow\\GPBMetadata\xea\x02\x1cNebulaidl::Plugins::Kubeflowb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.plugins.kubeflow.mpi_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\036com.nebulaidl.plugins.kubeflowB\010MpiProtoP\001Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\242\002\003NPK\252\002\032Nebulaidl.Plugins.Kubeflow\312\002\032Nebulaidl\\Plugins\\Kubeflow\342\002&Nebulaidl\\Plugins\\Kubeflow\\GPBMetadata\352\002\034Nebulaidl::Plugins::Kubeflow'
  _globals['_DISTRIBUTEDMPITRAININGTASK']._serialized_start=138
  _globals['_DISTRIBUTEDMPITRAININGTASK']._serialized_end=470
  _globals['_DISTRIBUTEDMPITRAININGREPLICASPEC']._serialized_start=473
  _globals['_DISTRIBUTEDMPITRAININGREPLICASPEC']._serialized_end=723
# @@protoc_insertion_point(module_scope)
