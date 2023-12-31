# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/literals.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2
from nebulaidl.core import types_pb2 as nebulaidl_dot_core_dot_types__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1dnebulaidl/core/literals.proto\x12\x0enebulaidl.core\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1cgoogle/protobuf/struct.proto\x1a\x1anebulaidl/core/types.proto\"\x87\x02\n\tPrimitive\x12\x1a\n\x07integer\x18\x01 \x01(\x03H\x00R\x07integer\x12!\n\x0b\x66loat_value\x18\x02 \x01(\x01H\x00R\nfloatValue\x12#\n\x0cstring_value\x18\x03 \x01(\tH\x00R\x0bstringValue\x12\x1a\n\x07\x62oolean\x18\x04 \x01(\x08H\x00R\x07\x62oolean\x12\x38\n\x08\x64\x61tetime\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.TimestampH\x00R\x08\x64\x61tetime\x12\x37\n\x08\x64uration\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00R\x08\x64urationB\x07\n\x05value\"\x06\n\x04Void\"R\n\x04\x42lob\x12\x38\n\x08metadata\x18\x01 \x01(\x0b\x32\x1c.nebulaidl.core.BlobMetadataR\x08metadata\x12\x10\n\x03uri\x18\x03 \x01(\tR\x03uri\"<\n\x0c\x42lobMetadata\x12,\n\x04type\x18\x01 \x01(\x0b\x32\x18.nebulaidl.core.BlobTypeR\x04type\"0\n\x06\x42inary\x12\x14\n\x05value\x18\x01 \x01(\x0cR\x05value\x12\x10\n\x03tag\x18\x02 \x01(\tR\x03tag\"J\n\x06Schema\x12\x10\n\x03uri\x18\x01 \x01(\tR\x03uri\x12.\n\x04type\x18\x03 \x01(\x0b\x32\x1a.nebulaidl.core.SchemaTypeR\x04type\"g\n\x05Union\x12-\n\x05value\x18\x01 \x01(\x0b\x32\x17.nebulaidl.core.LiteralR\x05value\x12/\n\x04type\x18\x02 \x01(\x0b\x32\x1b.nebulaidl.core.LiteralTypeR\x04type\"z\n\x19StructuredDatasetMetadata\x12]\n\x17structured_dataset_type\x18\x01 \x01(\x0b\x32%.nebulaidl.core.StructuredDatasetTypeR\x15structuredDatasetType\"l\n\x11StructuredDataset\x12\x10\n\x03uri\x18\x01 \x01(\tR\x03uri\x12\x45\n\x08metadata\x18\x02 \x01(\x0b\x32).nebulaidl.core.StructuredDatasetMetadataR\x08metadata\"\xf8\x03\n\x06Scalar\x12\x39\n\tprimitive\x18\x01 \x01(\x0b\x32\x19.nebulaidl.core.PrimitiveH\x00R\tprimitive\x12*\n\x04\x62lob\x18\x02 \x01(\x0b\x32\x14.nebulaidl.core.BlobH\x00R\x04\x62lob\x12\x30\n\x06\x62inary\x18\x03 \x01(\x0b\x32\x16.nebulaidl.core.BinaryH\x00R\x06\x62inary\x12\x30\n\x06schema\x18\x04 \x01(\x0b\x32\x16.nebulaidl.core.SchemaH\x00R\x06schema\x12\x33\n\tnone_type\x18\x05 \x01(\x0b\x32\x14.nebulaidl.core.VoidH\x00R\x08noneType\x12-\n\x05\x65rror\x18\x06 \x01(\x0b\x32\x15.nebulaidl.core.ErrorH\x00R\x05\x65rror\x12\x33\n\x07generic\x18\x07 \x01(\x0b\x32\x17.google.protobuf.StructH\x00R\x07generic\x12R\n\x12structured_dataset\x18\x08 \x01(\x0b\x32!.nebulaidl.core.StructuredDatasetH\x00R\x11structuredDataset\x12-\n\x05union\x18\t \x01(\x0b\x32\x15.nebulaidl.core.UnionH\x00R\x05unionB\x07\n\x05value\"\xcd\x01\n\x07Literal\x12\x30\n\x06scalar\x18\x01 \x01(\x0b\x32\x16.nebulaidl.core.ScalarH\x00R\x06scalar\x12\x43\n\ncollection\x18\x02 \x01(\x0b\x32!.nebulaidl.core.LiteralCollectionH\x00R\ncollection\x12.\n\x03map\x18\x03 \x01(\x0b\x32\x1a.nebulaidl.core.LiteralMapH\x00R\x03map\x12\x12\n\x04hash\x18\x04 \x01(\tR\x04hashB\x07\n\x05value\"H\n\x11LiteralCollection\x12\x33\n\x08literals\x18\x01 \x03(\x0b\x32\x17.nebulaidl.core.LiteralR\x08literals\"\xa8\x01\n\nLiteralMap\x12\x44\n\x08literals\x18\x01 \x03(\x0b\x32(.nebulaidl.core.LiteralMap.LiteralsEntryR\x08literals\x1aT\n\rLiteralsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12-\n\x05value\x18\x02 \x01(\x0b\x32\x17.nebulaidl.core.LiteralR\x05value:\x02\x38\x01\"P\n\x15\x42indingDataCollection\x12\x37\n\x08\x62indings\x18\x01 \x03(\x0b\x32\x1b.nebulaidl.core.BindingDataR\x08\x62indings\"\xb4\x01\n\x0e\x42indingDataMap\x12H\n\x08\x62indings\x18\x01 \x03(\x0b\x32,.nebulaidl.core.BindingDataMap.BindingsEntryR\x08\x62indings\x1aX\n\rBindingsEntry\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x31\n\x05value\x18\x02 \x01(\x0b\x32\x1b.nebulaidl.core.BindingDataR\x05value:\x02\x38\x01\"H\n\tUnionInfo\x12;\n\ntargetType\x18\x01 \x01(\x0b\x32\x1b.nebulaidl.core.LiteralTypeR\ntargetType\"\xb3\x02\n\x0b\x42indingData\x12\x30\n\x06scalar\x18\x01 \x01(\x0b\x32\x16.nebulaidl.core.ScalarH\x00R\x06scalar\x12G\n\ncollection\x18\x02 \x01(\x0b\x32%.nebulaidl.core.BindingDataCollectionH\x00R\ncollection\x12;\n\x07promise\x18\x03 \x01(\x0b\x32\x1f.nebulaidl.core.OutputReferenceH\x00R\x07promise\x12\x32\n\x03map\x18\x04 \x01(\x0b\x32\x1e.nebulaidl.core.BindingDataMapH\x00R\x03map\x12/\n\x05union\x18\x05 \x01(\x0b\x32\x19.nebulaidl.core.UnionInfoR\x05unionB\x07\n\x05value\"R\n\x07\x42inding\x12\x10\n\x03var\x18\x01 \x01(\tR\x03var\x12\x35\n\x07\x62inding\x18\x02 \x01(\x0b\x32\x1b.nebulaidl.core.BindingDataR\x07\x62inding\"6\n\x0cKeyValuePair\x12\x10\n\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n\x05value\x18\x02 \x01(\tR\x05value\")\n\rRetryStrategy\x12\x18\n\x07retries\x18\x05 \x01(\rR\x07retriesB\xb8\x01\n\x12\x63om.nebulaidl.coreB\rLiteralsProtoP\x01Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03NCX\xaa\x02\x0eNebulaidl.Core\xca\x02\x0eNebulaidl\\Core\xe2\x02\x1aNebulaidl\\Core\\GPBMetadata\xea\x02\x0fNebulaidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.literals_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.coreB\rLiteralsProtoP\001Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003NCX\252\002\016Nebulaidl.Core\312\002\016Nebulaidl\\Core\342\002\032Nebulaidl\\Core\\GPBMetadata\352\002\017Nebulaidl::Core'
  _LITERALMAP_LITERALSENTRY._options = None
  _LITERALMAP_LITERALSENTRY._serialized_options = b'8\001'
  _BINDINGDATAMAP_BINDINGSENTRY._options = None
  _BINDINGDATAMAP_BINDINGSENTRY._serialized_options = b'8\001'
  _globals['_PRIMITIVE']._serialized_start=173
  _globals['_PRIMITIVE']._serialized_end=436
  _globals['_VOID']._serialized_start=438
  _globals['_VOID']._serialized_end=444
  _globals['_BLOB']._serialized_start=446
  _globals['_BLOB']._serialized_end=528
  _globals['_BLOBMETADATA']._serialized_start=530
  _globals['_BLOBMETADATA']._serialized_end=590
  _globals['_BINARY']._serialized_start=592
  _globals['_BINARY']._serialized_end=640
  _globals['_SCHEMA']._serialized_start=642
  _globals['_SCHEMA']._serialized_end=716
  _globals['_UNION']._serialized_start=718
  _globals['_UNION']._serialized_end=821
  _globals['_STRUCTUREDDATASETMETADATA']._serialized_start=823
  _globals['_STRUCTUREDDATASETMETADATA']._serialized_end=945
  _globals['_STRUCTUREDDATASET']._serialized_start=947
  _globals['_STRUCTUREDDATASET']._serialized_end=1055
  _globals['_SCALAR']._serialized_start=1058
  _globals['_SCALAR']._serialized_end=1562
  _globals['_LITERAL']._serialized_start=1565
  _globals['_LITERAL']._serialized_end=1770
  _globals['_LITERALCOLLECTION']._serialized_start=1772
  _globals['_LITERALCOLLECTION']._serialized_end=1844
  _globals['_LITERALMAP']._serialized_start=1847
  _globals['_LITERALMAP']._serialized_end=2015
  _globals['_LITERALMAP_LITERALSENTRY']._serialized_start=1931
  _globals['_LITERALMAP_LITERALSENTRY']._serialized_end=2015
  _globals['_BINDINGDATACOLLECTION']._serialized_start=2017
  _globals['_BINDINGDATACOLLECTION']._serialized_end=2097
  _globals['_BINDINGDATAMAP']._serialized_start=2100
  _globals['_BINDINGDATAMAP']._serialized_end=2280
  _globals['_BINDINGDATAMAP_BINDINGSENTRY']._serialized_start=2192
  _globals['_BINDINGDATAMAP_BINDINGSENTRY']._serialized_end=2280
  _globals['_UNIONINFO']._serialized_start=2282
  _globals['_UNIONINFO']._serialized_end=2354
  _globals['_BINDINGDATA']._serialized_start=2357
  _globals['_BINDINGDATA']._serialized_end=2664
  _globals['_BINDING']._serialized_start=2666
  _globals['_BINDING']._serialized_end=2748
  _globals['_KEYVALUEPAIR']._serialized_start=2750
  _globals['_KEYVALUEPAIR']._serialized_end=2804
  _globals['_RETRYSTRATEGY']._serialized_start=2806
  _globals['_RETRYSTRATEGY']._serialized_end=2847
# @@protoc_insertion_point(module_scope)
