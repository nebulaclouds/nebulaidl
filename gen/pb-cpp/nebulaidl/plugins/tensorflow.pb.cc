// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/tensorflow.proto

#include "nebulaidl/plugins/tensorflow.pb.h"

#include <algorithm>

#include <google/protobuf/stubs/common.h>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite_inl.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

namespace nebulaidl {
namespace plugins {
class DistributedTensorflowTrainingTaskDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<DistributedTensorflowTrainingTask> _instance;
} _DistributedTensorflowTrainingTask_default_instance_;
}  // namespace plugins
}  // namespace nebulaidl
static void InitDefaultsDistributedTensorflowTrainingTask_nebulaidl_2fplugins_2ftensorflow_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::nebulaidl::plugins::_DistributedTensorflowTrainingTask_default_instance_;
    new (ptr) ::nebulaidl::plugins::DistributedTensorflowTrainingTask();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::nebulaidl::plugins::DistributedTensorflowTrainingTask::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_DistributedTensorflowTrainingTask_nebulaidl_2fplugins_2ftensorflow_2eproto =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsDistributedTensorflowTrainingTask_nebulaidl_2fplugins_2ftensorflow_2eproto}, {}};

void InitDefaults_nebulaidl_2fplugins_2ftensorflow_2eproto() {
  ::google::protobuf::internal::InitSCC(&scc_info_DistributedTensorflowTrainingTask_nebulaidl_2fplugins_2ftensorflow_2eproto.base);
}

::google::protobuf::Metadata file_level_metadata_nebulaidl_2fplugins_2ftensorflow_2eproto[1];
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_nebulaidl_2fplugins_2ftensorflow_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_nebulaidl_2fplugins_2ftensorflow_2eproto = nullptr;

const ::google::protobuf::uint32 TableStruct_nebulaidl_2fplugins_2ftensorflow_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedTensorflowTrainingTask, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedTensorflowTrainingTask, workers_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedTensorflowTrainingTask, ps_replicas_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedTensorflowTrainingTask, chief_replicas_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::nebulaidl::plugins::DistributedTensorflowTrainingTask)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::nebulaidl::plugins::_DistributedTensorflowTrainingTask_default_instance_),
};

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_nebulaidl_2fplugins_2ftensorflow_2eproto = {
  {}, AddDescriptors_nebulaidl_2fplugins_2ftensorflow_2eproto, "nebulaidl/plugins/tensorflow.proto", schemas,
  file_default_instances, TableStruct_nebulaidl_2fplugins_2ftensorflow_2eproto::offsets,
  file_level_metadata_nebulaidl_2fplugins_2ftensorflow_2eproto, 1, file_level_enum_descriptors_nebulaidl_2fplugins_2ftensorflow_2eproto, file_level_service_descriptors_nebulaidl_2fplugins_2ftensorflow_2eproto,
};

const char descriptor_table_protodef_nebulaidl_2fplugins_2ftensorflow_2eproto[] =
  "\n\"nebulaidl/plugins/tensorflow.proto\022\021ne"
  "bulaidl.plugins\"a\n!DistributedTensorflow"
  "TrainingTask\022\017\n\007workers\030\001 \001(\005\022\023\n\013ps_repl"
  "icas\030\002 \001(\005\022\026\n\016chief_replicas\030\003 \001(\005B\?Z=gi"
  "thub.com/nebulaclouds/nebulaidl/gen/pb-g"
  "o/nebulaidl/pluginsb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_nebulaidl_2fplugins_2ftensorflow_2eproto = {
  false, InitDefaults_nebulaidl_2fplugins_2ftensorflow_2eproto, 
  descriptor_table_protodef_nebulaidl_2fplugins_2ftensorflow_2eproto,
  "nebulaidl/plugins/tensorflow.proto", &assign_descriptors_table_nebulaidl_2fplugins_2ftensorflow_2eproto, 227,
};

void AddDescriptors_nebulaidl_2fplugins_2ftensorflow_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[1] =
  {
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_nebulaidl_2fplugins_2ftensorflow_2eproto, deps, 0);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_nebulaidl_2fplugins_2ftensorflow_2eproto = []() { AddDescriptors_nebulaidl_2fplugins_2ftensorflow_2eproto(); return true; }();
namespace nebulaidl {
namespace plugins {

// ===================================================================

void DistributedTensorflowTrainingTask::InitAsDefaultInstance() {
}
class DistributedTensorflowTrainingTask::HasBitSetters {
 public:
};

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int DistributedTensorflowTrainingTask::kWorkersFieldNumber;
const int DistributedTensorflowTrainingTask::kPsReplicasFieldNumber;
const int DistributedTensorflowTrainingTask::kChiefReplicasFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

DistributedTensorflowTrainingTask::DistributedTensorflowTrainingTask()
  : ::google::protobuf::Message(), _internal_metadata_(nullptr) {
  SharedCtor();
  // @@protoc_insertion_point(constructor:nebulaidl.plugins.DistributedTensorflowTrainingTask)
}
DistributedTensorflowTrainingTask::DistributedTensorflowTrainingTask(const DistributedTensorflowTrainingTask& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(nullptr) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::memcpy(&workers_, &from.workers_,
    static_cast<size_t>(reinterpret_cast<char*>(&chief_replicas_) -
    reinterpret_cast<char*>(&workers_)) + sizeof(chief_replicas_));
  // @@protoc_insertion_point(copy_constructor:nebulaidl.plugins.DistributedTensorflowTrainingTask)
}

void DistributedTensorflowTrainingTask::SharedCtor() {
  ::memset(&workers_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&chief_replicas_) -
      reinterpret_cast<char*>(&workers_)) + sizeof(chief_replicas_));
}

DistributedTensorflowTrainingTask::~DistributedTensorflowTrainingTask() {
  // @@protoc_insertion_point(destructor:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  SharedDtor();
}

void DistributedTensorflowTrainingTask::SharedDtor() {
}

void DistributedTensorflowTrainingTask::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const DistributedTensorflowTrainingTask& DistributedTensorflowTrainingTask::default_instance() {
  ::google::protobuf::internal::InitSCC(&::scc_info_DistributedTensorflowTrainingTask_nebulaidl_2fplugins_2ftensorflow_2eproto.base);
  return *internal_default_instance();
}


void DistributedTensorflowTrainingTask::Clear() {
// @@protoc_insertion_point(message_clear_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  ::memset(&workers_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&chief_replicas_) -
      reinterpret_cast<char*>(&workers_)) + sizeof(chief_replicas_));
  _internal_metadata_.Clear();
}

#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
const char* DistributedTensorflowTrainingTask::_InternalParse(const char* begin, const char* end, void* object,
                  ::google::protobuf::internal::ParseContext* ctx) {
  auto msg = static_cast<DistributedTensorflowTrainingTask*>(object);
  ::google::protobuf::int32 size; (void)size;
  int depth; (void)depth;
  ::google::protobuf::uint32 tag;
  ::google::protobuf::internal::ParseFunc parser_till_end; (void)parser_till_end;
  auto ptr = begin;
  while (ptr < end) {
    ptr = ::google::protobuf::io::Parse32(ptr, &tag);
    GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
    switch (tag >> 3) {
      // int32 workers = 1;
      case 1: {
        if (static_cast<::google::protobuf::uint8>(tag) != 8) goto handle_unusual;
        msg->set_workers(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // int32 ps_replicas = 2;
      case 2: {
        if (static_cast<::google::protobuf::uint8>(tag) != 16) goto handle_unusual;
        msg->set_ps_replicas(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // int32 chief_replicas = 3;
      case 3: {
        if (static_cast<::google::protobuf::uint8>(tag) != 24) goto handle_unusual;
        msg->set_chief_replicas(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->EndGroup(tag);
          return ptr;
        }
        auto res = UnknownFieldParse(tag, {_InternalParse, msg},
          ptr, end, msg->_internal_metadata_.mutable_unknown_fields(), ctx);
        ptr = res.first;
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr != nullptr);
        if (res.second) return ptr;
      }
    }  // switch
  }  // while
  return ptr;
}
#else  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
bool DistributedTensorflowTrainingTask::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!PROTOBUF_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // int32 workers = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (8 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &workers_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // int32 ps_replicas = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (16 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &ps_replicas_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // int32 chief_replicas = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (24 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &chief_replicas_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      default: {
      handle_unusual:
        if (tag == 0) {
          goto success;
        }
        DO_(::google::protobuf::internal::WireFormat::SkipField(
              input, tag, _internal_metadata_.mutable_unknown_fields()));
        break;
      }
    }
  }
success:
  // @@protoc_insertion_point(parse_success:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  return false;
#undef DO_
}
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER

void DistributedTensorflowTrainingTask::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 workers = 1;
  if (this->workers() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(1, this->workers(), output);
  }

  // int32 ps_replicas = 2;
  if (this->ps_replicas() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(2, this->ps_replicas(), output);
  }

  // int32 chief_replicas = 3;
  if (this->chief_replicas() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(3, this->chief_replicas(), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:nebulaidl.plugins.DistributedTensorflowTrainingTask)
}

::google::protobuf::uint8* DistributedTensorflowTrainingTask::InternalSerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 workers = 1;
  if (this->workers() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(1, this->workers(), target);
  }

  // int32 ps_replicas = 2;
  if (this->ps_replicas() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(2, this->ps_replicas(), target);
  }

  // int32 chief_replicas = 3;
  if (this->chief_replicas() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(3, this->chief_replicas(), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  return target;
}

size_t DistributedTensorflowTrainingTask::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // int32 workers = 1;
  if (this->workers() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->workers());
  }

  // int32 ps_replicas = 2;
  if (this->ps_replicas() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->ps_replicas());
  }

  // int32 chief_replicas = 3;
  if (this->chief_replicas() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->chief_replicas());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void DistributedTensorflowTrainingTask::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  GOOGLE_DCHECK_NE(&from, this);
  const DistributedTensorflowTrainingTask* source =
      ::google::protobuf::DynamicCastToGenerated<DistributedTensorflowTrainingTask>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:nebulaidl.plugins.DistributedTensorflowTrainingTask)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:nebulaidl.plugins.DistributedTensorflowTrainingTask)
    MergeFrom(*source);
  }
}

void DistributedTensorflowTrainingTask::MergeFrom(const DistributedTensorflowTrainingTask& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.workers() != 0) {
    set_workers(from.workers());
  }
  if (from.ps_replicas() != 0) {
    set_ps_replicas(from.ps_replicas());
  }
  if (from.chief_replicas() != 0) {
    set_chief_replicas(from.chief_replicas());
  }
}

void DistributedTensorflowTrainingTask::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void DistributedTensorflowTrainingTask::CopyFrom(const DistributedTensorflowTrainingTask& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:nebulaidl.plugins.DistributedTensorflowTrainingTask)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool DistributedTensorflowTrainingTask::IsInitialized() const {
  return true;
}

void DistributedTensorflowTrainingTask::Swap(DistributedTensorflowTrainingTask* other) {
  if (other == this) return;
  InternalSwap(other);
}
void DistributedTensorflowTrainingTask::InternalSwap(DistributedTensorflowTrainingTask* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(workers_, other->workers_);
  swap(ps_replicas_, other->ps_replicas_);
  swap(chief_replicas_, other->chief_replicas_);
}

::google::protobuf::Metadata DistributedTensorflowTrainingTask::GetMetadata() const {
  ::google::protobuf::internal::AssignDescriptors(&::assign_descriptors_table_nebulaidl_2fplugins_2ftensorflow_2eproto);
  return ::file_level_metadata_nebulaidl_2fplugins_2ftensorflow_2eproto[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace plugins
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> PROTOBUF_NOINLINE ::nebulaidl::plugins::DistributedTensorflowTrainingTask* Arena::CreateMaybeMessage< ::nebulaidl::plugins::DistributedTensorflowTrainingTask >(Arena* arena) {
  return Arena::CreateInternal< ::nebulaidl::plugins::DistributedTensorflowTrainingTask >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
