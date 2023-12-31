# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/plugins/sagemaker/hyperparameter_tuning_job.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.plugins.sagemaker import parameter_ranges_pb2 as nebulaidl_dot_plugins_dot_sagemaker_dot_parameter__ranges__pb2
from nebulaidl.plugins.sagemaker import training_job_pb2 as nebulaidl_dot_plugins_dot_sagemaker_dot_training__job__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n;nebulaidl/plugins/sagemaker/hyperparameter_tuning_job.proto\x12\x1bnebulaidl.plugins.sagemaker\x1a\x32nebulaidl/plugins/sagemaker/parameter_ranges.proto\x1a.nebulaidl/plugins/sagemaker/training_job.proto\"\xe1\x01\n\x17HyperparameterTuningJob\x12K\n\x0ctraining_job\x18\x01 \x01(\x0b\x32(.nebulaidl.plugins.sagemaker.TrainingJobR\x0btrainingJob\x12<\n\x1bmax_number_of_training_jobs\x18\x02 \x01(\x03R\x17maxNumberOfTrainingJobs\x12;\n\x1amax_parallel_training_jobs\x18\x03 \x01(\x03R\x17maxParallelTrainingJobs\"H\n!HyperparameterTuningObjectiveType\"#\n\x05Value\x12\x0c\n\x08MINIMIZE\x10\x00\x12\x0c\n\x08MAXIMIZE\x10\x01\"\xad\x01\n\x1dHyperparameterTuningObjective\x12k\n\x0eobjective_type\x18\x01 \x01(\x0e\x32\x44.nebulaidl.plugins.sagemaker.HyperparameterTuningObjectiveType.ValueR\robjectiveType\x12\x1f\n\x0bmetric_name\x18\x02 \x01(\tR\nmetricName\"A\n\x1cHyperparameterTuningStrategy\"!\n\x05Value\x12\x0c\n\x08\x42\x41YESIAN\x10\x00\x12\n\n\x06RANDOM\x10\x01\":\n\x1cTrainingJobEarlyStoppingType\"\x1a\n\x05Value\x12\x07\n\x03OFF\x10\x00\x12\x08\n\x04\x41UTO\x10\x01\"\xdd\x03\n\x1dHyperparameterTuningJobConfig\x12\x61\n\x15hyperparameter_ranges\x18\x01 \x01(\x0b\x32,.nebulaidl.plugins.sagemaker.ParameterRangesR\x14hyperparameterRanges\x12h\n\x0ftuning_strategy\x18\x02 \x01(\x0e\x32?.nebulaidl.plugins.sagemaker.HyperparameterTuningStrategy.ValueR\x0etuningStrategy\x12\x65\n\x10tuning_objective\x18\x03 \x01(\x0b\x32:.nebulaidl.plugins.sagemaker.HyperparameterTuningObjectiveR\x0ftuningObjective\x12\x87\x01\n training_job_early_stopping_type\x18\x04 \x01(\x0e\x32?.nebulaidl.plugins.sagemaker.TrainingJobEarlyStoppingType.ValueR\x1ctrainingJobEarlyStoppingTypeB\x8c\x02\n\x1f\x63om.nebulaidl.plugins.sagemakerB\x1cHyperparameterTuningJobProtoP\x01Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\xa2\x02\x03NPS\xaa\x02\x1bNebulaidl.Plugins.Sagemaker\xca\x02\x1bNebulaidl\\Plugins\\Sagemaker\xe2\x02\'Nebulaidl\\Plugins\\Sagemaker\\GPBMetadata\xea\x02\x1dNebulaidl::Plugins::Sagemakerb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.plugins.sagemaker.hyperparameter_tuning_job_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\037com.nebulaidl.plugins.sagemakerB\034HyperparameterTuningJobProtoP\001Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/plugins\242\002\003NPS\252\002\033Nebulaidl.Plugins.Sagemaker\312\002\033Nebulaidl\\Plugins\\Sagemaker\342\002\'Nebulaidl\\Plugins\\Sagemaker\\GPBMetadata\352\002\035Nebulaidl::Plugins::Sagemaker'
  _globals['_HYPERPARAMETERTUNINGJOB']._serialized_start=193
  _globals['_HYPERPARAMETERTUNINGJOB']._serialized_end=418
  _globals['_HYPERPARAMETERTUNINGOBJECTIVETYPE']._serialized_start=420
  _globals['_HYPERPARAMETERTUNINGOBJECTIVETYPE']._serialized_end=492
  _globals['_HYPERPARAMETERTUNINGOBJECTIVETYPE_VALUE']._serialized_start=457
  _globals['_HYPERPARAMETERTUNINGOBJECTIVETYPE_VALUE']._serialized_end=492
  _globals['_HYPERPARAMETERTUNINGOBJECTIVE']._serialized_start=495
  _globals['_HYPERPARAMETERTUNINGOBJECTIVE']._serialized_end=668
  _globals['_HYPERPARAMETERTUNINGSTRATEGY']._serialized_start=670
  _globals['_HYPERPARAMETERTUNINGSTRATEGY']._serialized_end=735
  _globals['_HYPERPARAMETERTUNINGSTRATEGY_VALUE']._serialized_start=702
  _globals['_HYPERPARAMETERTUNINGSTRATEGY_VALUE']._serialized_end=735
  _globals['_TRAININGJOBEARLYSTOPPINGTYPE']._serialized_start=737
  _globals['_TRAININGJOBEARLYSTOPPINGTYPE']._serialized_end=795
  _globals['_TRAININGJOBEARLYSTOPPINGTYPE_VALUE']._serialized_start=769
  _globals['_TRAININGJOBEARLYSTOPPINGTYPE_VALUE']._serialized_end=795
  _globals['_HYPERPARAMETERTUNINGJOBCONFIG']._serialized_start=798
  _globals['_HYPERPARAMETERTUNINGJOBCONFIG']._serialized_end=1275
# @@protoc_insertion_point(module_scope)
