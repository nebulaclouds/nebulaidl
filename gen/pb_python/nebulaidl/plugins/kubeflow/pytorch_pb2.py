# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/plugins/kubeflow/pytorch.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import tasks_pb2 as nebulaidl_dot_core_dot_tasks__pb2
from nebulaidl.plugins.kubeflow import common_pb2 as nebulaidl_dot_plugins_dot_kubeflow_dot_common__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n(nebulaidl/plugins/kubeflow/pytorch.proto\x12\x1anebulaidl.plugins.kubeflow\x1a\x1anebulaidl/core/tasks.proto\x1a\'nebulaidl/plugins/kubeflow/common.proto\"\xc1\x01\n\rElasticConfig\x12!\n\x0crdzv_backend\x18\x01 \x01(\tR\x0brdzvBackend\x12!\n\x0cmin_replicas\x18\x02 \x01(\x05R\x0bminReplicas\x12!\n\x0cmax_replicas\x18\x03 \x01(\x05R\x0bmaxReplicas\x12$\n\x0enproc_per_node\x18\x04 \x01(\x05R\x0cnprocPerNode\x12!\n\x0cmax_restarts\x18\x05 \x01(\x05R\x0bmaxRestarts\"\x90\x03\n\x1e\x44istributedPyTorchTrainingTask\x12j\n\x0fworker_replicas\x18\x01 \x01(\x0b\x32\x41.nebulaidl.plugins.kubeflow.DistributedPyTorchTrainingReplicaSpecR\x0eworkerReplicas\x12j\n\x0fmaster_replicas\x18\x02 \x01(\x0b\x32\x41.nebulaidl.plugins.kubeflow.DistributedPyTorchTrainingReplicaSpecR\x0emasterReplicas\x12\x44\n\nrun_policy\x18\x03 \x01(\x0b\x32%.nebulaidl.plugins.kubeflow.RunPolicyR\trunPolicy\x12P\n\x0e\x65lastic_config\x18\x04 \x01(\x0b\x32).nebulaidl.plugins.kubeflow.ElasticConfigR\relasticConfig\"\xe4\x01\n%DistributedPyTorchTrainingReplicaSpec\x12\x1a\n\x08replicas\x18\x01 \x01(\x05R\x08replicas\x12\x14\n\x05image\x18\x02 \x01(\tR\x05image\x12\x37\n\tresources\x18\x03 \x01(\x0b\x32\x19.nebulaidl.core.ResourcesR\tresources\x12P\n\x0erestart_policy\x18\x04 \x01(\x0e\x32).nebulaidl.plugins.kubeflow.RestartPolicyR\rrestartPolicyB\xf7\x01\n\x1e\x63om.nebulaidl.plugins.kubeflowB\x0cPytorchProtoP\x01Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\xa2\x02\x03NPK\xaa\x02\x1aNebulaidl.Plugins.Kubeflow\xca\x02\x1aNebulaidl\\Plugins\\Kubeflow\xe2\x02&Nebulaidl\\Plugins\\Kubeflow\\GPBMetadata\xea\x02\x1cNebulaidl::Plugins::Kubeflowb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.plugins.kubeflow.pytorch_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\036com.nebulaidl.plugins.kubeflowB\014PytorchProtoP\001Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\242\002\003NPK\252\002\032Nebulaidl.Plugins.Kubeflow\312\002\032Nebulaidl\\Plugins\\Kubeflow\342\002&Nebulaidl\\Plugins\\Kubeflow\\GPBMetadata\352\002\034Nebulaidl::Plugins::Kubeflow'
  _globals['_ELASTICCONFIG']._serialized_start=142
  _globals['_ELASTICCONFIG']._serialized_end=335
  _globals['_DISTRIBUTEDPYTORCHTRAININGTASK']._serialized_start=338
  _globals['_DISTRIBUTEDPYTORCHTRAININGTASK']._serialized_end=738
  _globals['_DISTRIBUTEDPYTORCHTRAININGREPLICASPEC']._serialized_start=741
  _globals['_DISTRIBUTEDPYTORCHTRAININGREPLICASPEC']._serialized_end=969
# @@protoc_insertion_point(module_scope)
