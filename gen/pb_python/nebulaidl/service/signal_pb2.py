# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/service/signal.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from nebulaidl.admin import signal_pb2 as nebulaidl_dot_admin_dot_signal__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1enebulaidl/service/signal.proto\x12\x11nebulaidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1cnebulaidl/admin/signal.proto2\xa0\x03\n\rSignalService\x12Y\n\x11GetOrCreateSignal\x12).nebulaidl.admin.SignalGetOrCreateRequest\x1a\x17.nebulaidl.admin.Signal\"\x00\x12\xc3\x01\n\x0bListSignals\x12\".nebulaidl.admin.SignalListRequest\x1a\x1b.nebulaidl.admin.SignalList\"s\x82\xd3\xe4\x93\x02m\x12k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\x12n\n\tSetSignal\x12!.nebulaidl.admin.SignalSetRequest\x1a\".nebulaidl.admin.SignalSetResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/api/v1/signalsB\xc8\x01\n\x15\x63om.nebulaidl.serviceB\x0bSignalProtoP\x01Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/service\xa2\x02\x03NSX\xaa\x02\x11Nebulaidl.Service\xca\x02\x11Nebulaidl\\Service\xe2\x02\x1dNebulaidl\\Service\\GPBMetadata\xea\x02\x12Nebulaidl::Serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.service.signal_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025com.nebulaidl.serviceB\013SignalProtoP\001Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/service\242\002\003NSX\252\002\021Nebulaidl.Service\312\002\021Nebulaidl\\Service\342\002\035Nebulaidl\\Service\\GPBMetadata\352\002\022Nebulaidl::Service'
  _SIGNALSERVICE.methods_by_name['ListSignals']._options = None
  _SIGNALSERVICE.methods_by_name['ListSignals']._serialized_options = b'\202\323\344\223\002m\022k/api/v1/signals/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}'
  _SIGNALSERVICE.methods_by_name['SetSignal']._options = None
  _SIGNALSERVICE.methods_by_name['SetSignal']._serialized_options = b'\202\323\344\223\002\024:\001*\"\017/api/v1/signals'
  _globals['_SIGNALSERVICE']._serialized_start=114
  _globals['_SIGNALSERVICE']._serialized_end=530
# @@protoc_insertion_point(module_scope)
