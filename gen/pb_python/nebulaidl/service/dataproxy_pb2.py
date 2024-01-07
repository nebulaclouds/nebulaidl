# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/service/dataproxy.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from nebulaidl.core import identifier_pb2 as nebulaidl_dot_core_dot_identifier__pb2
from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n!nebulaidl/service/dataproxy.proto\x12\x11nebulaidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1fnebulaidl/core/identifier.proto\x1a\x1dnebulaidl/core/literals.proto\"\x97\x01\n\x1c\x43reateUploadLocationResponse\x12\x1d\n\nsigned_url\x18\x01 \x01(\tR\tsignedUrl\x12\x1d\n\nnative_url\x18\x02 \x01(\tR\tnativeUrl\x12\x39\n\nexpires_at\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAt\"\xeb\x01\n\x1b\x43reateUploadLocationRequest\x12\x18\n\x07project\x18\x01 \x01(\tR\x07project\x12\x16\n\x06\x64omain\x18\x02 \x01(\tR\x06\x64omain\x12\x1a\n\x08\x66ilename\x18\x03 \x01(\tR\x08\x66ilename\x12\x38\n\nexpires_in\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\texpiresIn\x12\x1f\n\x0b\x63ontent_md5\x18\x05 \x01(\x0cR\ncontentMd5\x12#\n\rfilename_root\x18\x06 \x01(\tR\x0c\x66ilenameRoot\"|\n\x1d\x43reateDownloadLocationRequest\x12\x1d\n\nnative_url\x18\x01 \x01(\tR\tnativeUrl\x12\x38\n\nexpires_in\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\texpiresIn:\x02\x18\x01\"~\n\x1e\x43reateDownloadLocationResponse\x12\x1d\n\nsigned_url\x18\x01 \x01(\tR\tsignedUrl\x12\x39\n\nexpires_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAt:\x02\x18\x01\"\xfc\x01\n\x19\x43reateDownloadLinkRequest\x12\x44\n\rartifact_type\x18\x01 \x01(\x0e\x32\x1f.nebulaidl.service.ArtifactTypeR\x0c\x61rtifactType\x12\x38\n\nexpires_in\x18\x02 \x01(\x0b\x32\x19.google.protobuf.DurationR\texpiresIn\x12U\n\x11node_execution_id\x18\x03 \x01(\x0b\x32\'.nebulaidl.core.NodeExecutionIdentifierH\x00R\x0fnodeExecutionIdB\x08\n\x06source\"\xc8\x01\n\x1a\x43reateDownloadLinkResponse\x12!\n\nsigned_url\x18\x01 \x03(\tB\x02\x18\x01R\tsignedUrl\x12=\n\nexpires_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampB\x02\x18\x01R\texpiresAt\x12H\n\x0fpre_signed_urls\x18\x03 \x01(\x0b\x32 .nebulaidl.service.PreSignedURLsR\rpreSignedUrls\"i\n\rPreSignedURLs\x12\x1d\n\nsigned_url\x18\x01 \x03(\tR\tsignedUrl\x12\x39\n\nexpires_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\texpiresAt\"/\n\x0eGetDataRequest\x12\x1d\n\nnebula_url\x18\x01 \x01(\tR\tnebulaUrl\"\xd9\x01\n\x0fGetDataResponse\x12=\n\x0bliteral_map\x18\x01 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x00R\nliteralMap\x12J\n\x0fpre_signed_urls\x18\x02 \x01(\x0b\x32 .nebulaidl.service.PreSignedURLsH\x00R\rpreSignedUrls\x12\x33\n\x07literal\x18\x03 \x01(\x0b\x32\x17.nebulaidl.core.LiteralH\x00R\x07literalB\x06\n\x04\x64\x61ta*C\n\x0c\x41rtifactType\x12\x1b\n\x17\x41RTIFACT_TYPE_UNDEFINED\x10\x00\x12\x16\n\x12\x41RTIFACT_TYPE_DECK\x10\x01\x32\xea\x04\n\x10\x44\x61taProxyService\x12\xa2\x01\n\x14\x43reateUploadLocation\x12..nebulaidl.service.CreateUploadLocationRequest\x1a/.nebulaidl.service.CreateUploadLocationResponse\")\x82\xd3\xe4\x93\x02#:\x01*\"\x1e/api/v1/dataproxy/artifact_urn\x12\xa8\x01\n\x16\x43reateDownloadLocation\x12\x30.nebulaidl.service.CreateDownloadLocationRequest\x1a\x31.nebulaidl.service.CreateDownloadLocationResponse\")\x88\x02\x01\x82\xd3\xe4\x93\x02 \x12\x1e/api/v1/dataproxy/artifact_urn\x12\x9d\x01\n\x12\x43reateDownloadLink\x12,.nebulaidl.service.CreateDownloadLinkRequest\x1a-.nebulaidl.service.CreateDownloadLinkResponse\"*\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/api/v1/dataproxy/artifact_link\x12\x66\n\x07GetData\x12!.nebulaidl.service.GetDataRequest\x1a\".nebulaidl.service.GetDataResponse\"\x14\x82\xd3\xe4\x93\x02\x0e\x12\x0c/api/v1/dataB\xcb\x01\n\x15\x63om.nebulaidl.serviceB\x0e\x44\x61taproxyProtoP\x01Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/service\xa2\x02\x03NSX\xaa\x02\x11Nebulaidl.Service\xca\x02\x11Nebulaidl\\Service\xe2\x02\x1dNebulaidl\\Service\\GPBMetadata\xea\x02\x12Nebulaidl::Serviceb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.service.dataproxy_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\025com.nebulaidl.serviceB\016DataproxyProtoP\001Z=github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/service\242\002\003NSX\252\002\021Nebulaidl.Service\312\002\021Nebulaidl\\Service\342\002\035Nebulaidl\\Service\\GPBMetadata\352\002\022Nebulaidl::Service'
  _CREATEDOWNLOADLOCATIONREQUEST._options = None
  _CREATEDOWNLOADLOCATIONREQUEST._serialized_options = b'\030\001'
  _CREATEDOWNLOADLOCATIONRESPONSE._options = None
  _CREATEDOWNLOADLOCATIONRESPONSE._serialized_options = b'\030\001'
  _CREATEDOWNLOADLINKRESPONSE.fields_by_name['signed_url']._options = None
  _CREATEDOWNLOADLINKRESPONSE.fields_by_name['signed_url']._serialized_options = b'\030\001'
  _CREATEDOWNLOADLINKRESPONSE.fields_by_name['expires_at']._options = None
  _CREATEDOWNLOADLINKRESPONSE.fields_by_name['expires_at']._serialized_options = b'\030\001'
  _DATAPROXYSERVICE.methods_by_name['CreateUploadLocation']._options = None
  _DATAPROXYSERVICE.methods_by_name['CreateUploadLocation']._serialized_options = b'\202\323\344\223\002#:\001*\"\036/api/v1/dataproxy/artifact_urn'
  _DATAPROXYSERVICE.methods_by_name['CreateDownloadLocation']._options = None
  _DATAPROXYSERVICE.methods_by_name['CreateDownloadLocation']._serialized_options = b'\210\002\001\202\323\344\223\002 \022\036/api/v1/dataproxy/artifact_urn'
  _DATAPROXYSERVICE.methods_by_name['CreateDownloadLink']._options = None
  _DATAPROXYSERVICE.methods_by_name['CreateDownloadLink']._serialized_options = b'\202\323\344\223\002$:\001*\"\037/api/v1/dataproxy/artifact_link'
  _DATAPROXYSERVICE.methods_by_name['GetData']._options = None
  _DATAPROXYSERVICE.methods_by_name['GetData']._serialized_options = b'\202\323\344\223\002\016\022\014/api/v1/data'
  _globals['_ARTIFACTTYPE']._serialized_start=1695
  _globals['_ARTIFACTTYPE']._serialized_end=1762
  _globals['_CREATEUPLOADLOCATIONRESPONSE']._serialized_start=216
  _globals['_CREATEUPLOADLOCATIONRESPONSE']._serialized_end=367
  _globals['_CREATEUPLOADLOCATIONREQUEST']._serialized_start=370
  _globals['_CREATEUPLOADLOCATIONREQUEST']._serialized_end=605
  _globals['_CREATEDOWNLOADLOCATIONREQUEST']._serialized_start=607
  _globals['_CREATEDOWNLOADLOCATIONREQUEST']._serialized_end=731
  _globals['_CREATEDOWNLOADLOCATIONRESPONSE']._serialized_start=733
  _globals['_CREATEDOWNLOADLOCATIONRESPONSE']._serialized_end=859
  _globals['_CREATEDOWNLOADLINKREQUEST']._serialized_start=862
  _globals['_CREATEDOWNLOADLINKREQUEST']._serialized_end=1114
  _globals['_CREATEDOWNLOADLINKRESPONSE']._serialized_start=1117
  _globals['_CREATEDOWNLOADLINKRESPONSE']._serialized_end=1317
  _globals['_PRESIGNEDURLS']._serialized_start=1319
  _globals['_PRESIGNEDURLS']._serialized_end=1424
  _globals['_GETDATAREQUEST']._serialized_start=1426
  _globals['_GETDATAREQUEST']._serialized_end=1473
  _globals['_GETDATARESPONSE']._serialized_start=1476
  _globals['_GETDATARESPONSE']._serialized_end=1693
  _globals['_DATAPROXYSERVICE']._serialized_start=1765
  _globals['_DATAPROXYSERVICE']._serialized_end=2383
# @@protoc_insertion_point(module_scope)
