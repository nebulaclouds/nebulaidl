# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: nebulaidl/core/condition.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from nebulaidl.core import literals_pb2 as nebulaidl_dot_core_dot_literals__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1enebulaidl/core/condition.proto\x12\x0enebulaidl.core\x1a\x1dnebulaidl/core/literals.proto\"\x92\x02\n\x14\x43omparisonExpression\x12I\n\x08operator\x18\x01 \x01(\x0e\x32-.nebulaidl.core.ComparisonExpression.OperatorR\x08operator\x12\x36\n\nleft_value\x18\x02 \x01(\x0b\x32\x17.nebulaidl.core.OperandR\tleftValue\x12\x38\n\x0bright_value\x18\x03 \x01(\x0b\x32\x17.nebulaidl.core.OperandR\nrightValue\"=\n\x08Operator\x12\x06\n\x02\x45Q\x10\x00\x12\x07\n\x03NEQ\x10\x01\x12\x06\n\x02GT\x10\x02\x12\x07\n\x03GTE\x10\x03\x12\x06\n\x02LT\x10\x04\x12\x07\n\x03LTE\x10\x05\"\x95\x01\n\x07Operand\x12=\n\tprimitive\x18\x01 \x01(\x0b\x32\x19.nebulaidl.core.PrimitiveB\x02\x18\x01H\x00R\tprimitive\x12\x12\n\x03var\x18\x02 \x01(\tH\x00R\x03var\x12\x30\n\x06scalar\x18\x03 \x01(\x0b\x32\x16.nebulaidl.core.ScalarH\x00R\x06scalarB\x05\n\x03val\"\xae\x01\n\x11\x42ooleanExpression\x12I\n\x0b\x63onjunction\x18\x01 \x01(\x0b\x32%.nebulaidl.core.ConjunctionExpressionH\x00R\x0b\x63onjunction\x12\x46\n\ncomparison\x18\x02 \x01(\x0b\x32$.nebulaidl.core.ComparisonExpressionH\x00R\ncomparisonB\x06\n\x04\x65xpr\"\xa8\x02\n\x15\x43onjunctionExpression\x12Q\n\x08operator\x18\x01 \x01(\x0e\x32\x35.nebulaidl.core.ConjunctionExpression.LogicalOperatorR\x08operator\x12J\n\x0fleft_expression\x18\x02 \x01(\x0b\x32!.nebulaidl.core.BooleanExpressionR\x0eleftExpression\x12L\n\x10right_expression\x18\x03 \x01(\x0b\x32!.nebulaidl.core.BooleanExpressionR\x0frightExpression\"\"\n\x0fLogicalOperator\x12\x07\n\x03\x41ND\x10\x00\x12\x06\n\x02OR\x10\x01\x42\xb9\x01\n\x12\x63om.nebulaidl.coreB\x0e\x43onditionProtoP\x01Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\xa2\x02\x03NCX\xaa\x02\x0eNebulaidl.Core\xca\x02\x0eNebulaidl\\Core\xe2\x02\x1aNebulaidl\\Core\\GPBMetadata\xea\x02\x0fNebulaidl::Coreb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'nebulaidl.core.condition_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'\n\022com.nebulaidl.coreB\016ConditionProtoP\001Z:github.com/nebulaclouds/nebulaidl/gen/pb-go/nebulaidl/core\242\002\003NCX\252\002\016Nebulaidl.Core\312\002\016Nebulaidl\\Core\342\002\032Nebulaidl\\Core\\GPBMetadata\352\002\017Nebulaidl::Core'
  _OPERAND.fields_by_name['primitive']._options = None
  _OPERAND.fields_by_name['primitive']._serialized_options = b'\030\001'
  _globals['_COMPARISONEXPRESSION']._serialized_start=82
  _globals['_COMPARISONEXPRESSION']._serialized_end=356
  _globals['_COMPARISONEXPRESSION_OPERATOR']._serialized_start=295
  _globals['_COMPARISONEXPRESSION_OPERATOR']._serialized_end=356
  _globals['_OPERAND']._serialized_start=359
  _globals['_OPERAND']._serialized_end=508
  _globals['_BOOLEANEXPRESSION']._serialized_start=511
  _globals['_BOOLEANEXPRESSION']._serialized_end=685
  _globals['_CONJUNCTIONEXPRESSION']._serialized_start=688
  _globals['_CONJUNCTIONEXPRESSION']._serialized_end=984
  _globals['_CONJUNCTIONEXPRESSION_LOGICALOPERATOR']._serialized_start=950
  _globals['_CONJUNCTIONEXPRESSION_LOGICALOPERATOR']._serialized_end=984
# @@protoc_insertion_point(module_scope)
