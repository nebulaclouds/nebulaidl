// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/mpi.proto

#include "nebulaidl/plugins/mpi.pb.h"

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
class DistributedMPITrainingTaskDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<DistributedMPITrainingTask> _instance;
} _DistributedMPITrainingTask_default_instance_;
}  // namespace plugins
}  // namespace nebulaidl
static void InitDefaultsDistributedMPITrainingTask_nebulaidl_2fplugins_2fmpi_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::nebulaidl::plugins::_DistributedMPITrainingTask_default_instance_;
    new (ptr) ::nebulaidl::plugins::DistributedMPITrainingTask();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::nebulaidl::plugins::DistributedMPITrainingTask::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_DistributedMPITrainingTask_nebulaidl_2fplugins_2fmpi_2eproto =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsDistributedMPITrainingTask_nebulaidl_2fplugins_2fmpi_2eproto}, {}};

void InitDefaults_nebulaidl_2fplugins_2fmpi_2eproto() {
  ::google::protobuf::internal::InitSCC(&scc_info_DistributedMPITrainingTask_nebulaidl_2fplugins_2fmpi_2eproto.base);
}

::google::protobuf::Metadata file_level_metadata_nebulaidl_2fplugins_2fmpi_2eproto[1];
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_nebulaidl_2fplugins_2fmpi_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_nebulaidl_2fplugins_2fmpi_2eproto = nullptr;

const ::google::protobuf::uint32 TableStruct_nebulaidl_2fplugins_2fmpi_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedMPITrainingTask, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedMPITrainingTask, num_workers_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedMPITrainingTask, num_launcher_replicas_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::DistributedMPITrainingTask, slots_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::nebulaidl::plugins::DistributedMPITrainingTask)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::nebulaidl::plugins::_DistributedMPITrainingTask_default_instance_),
};

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_nebulaidl_2fplugins_2fmpi_2eproto = {
  {}, AddDescriptors_nebulaidl_2fplugins_2fmpi_2eproto, "nebulaidl/plugins/mpi.proto", schemas,
  file_default_instances, TableStruct_nebulaidl_2fplugins_2fmpi_2eproto::offsets,
  file_level_metadata_nebulaidl_2fplugins_2fmpi_2eproto, 1, file_level_enum_descriptors_nebulaidl_2fplugins_2fmpi_2eproto, file_level_service_descriptors_nebulaidl_2fplugins_2fmpi_2eproto,
};

const char descriptor_table_protodef_nebulaidl_2fplugins_2fmpi_2eproto[] =
  "\n\033nebulaidl/plugins/mpi.proto\022\021nebulaidl"
  ".plugins\"_\n\032DistributedMPITrainingTask\022\023"
  "\n\013num_workers\030\001 \001(\005\022\035\n\025num_launcher_repl"
  "icas\030\002 \001(\005\022\r\n\005slots\030\003 \001(\005B\?Z=github.com/"
  "nebulaclouds/nebulaidl/gen/pb-go/nebulai"
  "dl/pluginsb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_nebulaidl_2fplugins_2fmpi_2eproto = {
  false, InitDefaults_nebulaidl_2fplugins_2fmpi_2eproto, 
  descriptor_table_protodef_nebulaidl_2fplugins_2fmpi_2eproto,
  "nebulaidl/plugins/mpi.proto", &assign_descriptors_table_nebulaidl_2fplugins_2fmpi_2eproto, 218,
};

void AddDescriptors_nebulaidl_2fplugins_2fmpi_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[1] =
  {
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_nebulaidl_2fplugins_2fmpi_2eproto, deps, 0);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_nebulaidl_2fplugins_2fmpi_2eproto = []() { AddDescriptors_nebulaidl_2fplugins_2fmpi_2eproto(); return true; }();
namespace nebulaidl {
namespace plugins {

// ===================================================================

void DistributedMPITrainingTask::InitAsDefaultInstance() {
}
class DistributedMPITrainingTask::HasBitSetters {
 public:
};

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int DistributedMPITrainingTask::kNumWorkersFieldNumber;
const int DistributedMPITrainingTask::kNumLauncherReplicasFieldNumber;
const int DistributedMPITrainingTask::kSlotsFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

DistributedMPITrainingTask::DistributedMPITrainingTask()
  : ::google::protobuf::Message(), _internal_metadata_(nullptr) {
  SharedCtor();
  // @@protoc_insertion_point(constructor:nebulaidl.plugins.DistributedMPITrainingTask)
}
DistributedMPITrainingTask::DistributedMPITrainingTask(const DistributedMPITrainingTask& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(nullptr) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::memcpy(&num_workers_, &from.num_workers_,
    static_cast<size_t>(reinterpret_cast<char*>(&slots_) -
    reinterpret_cast<char*>(&num_workers_)) + sizeof(slots_));
  // @@protoc_insertion_point(copy_constructor:nebulaidl.plugins.DistributedMPITrainingTask)
}

void DistributedMPITrainingTask::SharedCtor() {
  ::memset(&num_workers_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&slots_) -
      reinterpret_cast<char*>(&num_workers_)) + sizeof(slots_));
}

DistributedMPITrainingTask::~DistributedMPITrainingTask() {
  // @@protoc_insertion_point(destructor:nebulaidl.plugins.DistributedMPITrainingTask)
  SharedDtor();
}

void DistributedMPITrainingTask::SharedDtor() {
}

void DistributedMPITrainingTask::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const DistributedMPITrainingTask& DistributedMPITrainingTask::default_instance() {
  ::google::protobuf::internal::InitSCC(&::scc_info_DistributedMPITrainingTask_nebulaidl_2fplugins_2fmpi_2eproto.base);
  return *internal_default_instance();
}


void DistributedMPITrainingTask::Clear() {
// @@protoc_insertion_point(message_clear_start:nebulaidl.plugins.DistributedMPITrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  ::memset(&num_workers_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&slots_) -
      reinterpret_cast<char*>(&num_workers_)) + sizeof(slots_));
  _internal_metadata_.Clear();
}

#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
const char* DistributedMPITrainingTask::_InternalParse(const char* begin, const char* end, void* object,
                  ::google::protobuf::internal::ParseContext* ctx) {
  auto msg = static_cast<DistributedMPITrainingTask*>(object);
  ::google::protobuf::int32 size; (void)size;
  int depth; (void)depth;
  ::google::protobuf::uint32 tag;
  ::google::protobuf::internal::ParseFunc parser_till_end; (void)parser_till_end;
  auto ptr = begin;
  while (ptr < end) {
    ptr = ::google::protobuf::io::Parse32(ptr, &tag);
    GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
    switch (tag >> 3) {
      // int32 num_workers = 1;
      case 1: {
        if (static_cast<::google::protobuf::uint8>(tag) != 8) goto handle_unusual;
        msg->set_num_workers(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // int32 num_launcher_replicas = 2;
      case 2: {
        if (static_cast<::google::protobuf::uint8>(tag) != 16) goto handle_unusual;
        msg->set_num_launcher_replicas(::google::protobuf::internal::ReadVarint(&ptr));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // int32 slots = 3;
      case 3: {
        if (static_cast<::google::protobuf::uint8>(tag) != 24) goto handle_unusual;
        msg->set_slots(::google::protobuf::internal::ReadVarint(&ptr));
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
bool DistributedMPITrainingTask::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!PROTOBUF_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:nebulaidl.plugins.DistributedMPITrainingTask)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // int32 num_workers = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (8 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &num_workers_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // int32 num_launcher_replicas = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (16 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &num_launcher_replicas_)));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // int32 slots = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (24 & 0xFF)) {

          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   ::google::protobuf::int32, ::google::protobuf::internal::WireFormatLite::TYPE_INT32>(
                 input, &slots_)));
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
  // @@protoc_insertion_point(parse_success:nebulaidl.plugins.DistributedMPITrainingTask)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:nebulaidl.plugins.DistributedMPITrainingTask)
  return false;
#undef DO_
}
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER

void DistributedMPITrainingTask::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:nebulaidl.plugins.DistributedMPITrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 num_workers = 1;
  if (this->num_workers() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(1, this->num_workers(), output);
  }

  // int32 num_launcher_replicas = 2;
  if (this->num_launcher_replicas() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(2, this->num_launcher_replicas(), output);
  }

  // int32 slots = 3;
  if (this->slots() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteInt32(3, this->slots(), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:nebulaidl.plugins.DistributedMPITrainingTask)
}

::google::protobuf::uint8* DistributedMPITrainingTask::InternalSerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:nebulaidl.plugins.DistributedMPITrainingTask)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // int32 num_workers = 1;
  if (this->num_workers() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(1, this->num_workers(), target);
  }

  // int32 num_launcher_replicas = 2;
  if (this->num_launcher_replicas() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(2, this->num_launcher_replicas(), target);
  }

  // int32 slots = 3;
  if (this->slots() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteInt32ToArray(3, this->slots(), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:nebulaidl.plugins.DistributedMPITrainingTask)
  return target;
}

size_t DistributedMPITrainingTask::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:nebulaidl.plugins.DistributedMPITrainingTask)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // int32 num_workers = 1;
  if (this->num_workers() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->num_workers());
  }

  // int32 num_launcher_replicas = 2;
  if (this->num_launcher_replicas() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->num_launcher_replicas());
  }

  // int32 slots = 3;
  if (this->slots() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::Int32Size(
        this->slots());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void DistributedMPITrainingTask::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:nebulaidl.plugins.DistributedMPITrainingTask)
  GOOGLE_DCHECK_NE(&from, this);
  const DistributedMPITrainingTask* source =
      ::google::protobuf::DynamicCastToGenerated<DistributedMPITrainingTask>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:nebulaidl.plugins.DistributedMPITrainingTask)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:nebulaidl.plugins.DistributedMPITrainingTask)
    MergeFrom(*source);
  }
}

void DistributedMPITrainingTask::MergeFrom(const DistributedMPITrainingTask& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:nebulaidl.plugins.DistributedMPITrainingTask)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.num_workers() != 0) {
    set_num_workers(from.num_workers());
  }
  if (from.num_launcher_replicas() != 0) {
    set_num_launcher_replicas(from.num_launcher_replicas());
  }
  if (from.slots() != 0) {
    set_slots(from.slots());
  }
}

void DistributedMPITrainingTask::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:nebulaidl.plugins.DistributedMPITrainingTask)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void DistributedMPITrainingTask::CopyFrom(const DistributedMPITrainingTask& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:nebulaidl.plugins.DistributedMPITrainingTask)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool DistributedMPITrainingTask::IsInitialized() const {
  return true;
}

void DistributedMPITrainingTask::Swap(DistributedMPITrainingTask* other) {
  if (other == this) return;
  InternalSwap(other);
}
void DistributedMPITrainingTask::InternalSwap(DistributedMPITrainingTask* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
  swap(num_workers_, other->num_workers_);
  swap(num_launcher_replicas_, other->num_launcher_replicas_);
  swap(slots_, other->slots_);
}

::google::protobuf::Metadata DistributedMPITrainingTask::GetMetadata() const {
  ::google::protobuf::internal::AssignDescriptors(&::assign_descriptors_table_nebulaidl_2fplugins_2fmpi_2eproto);
  return ::file_level_metadata_nebulaidl_2fplugins_2fmpi_2eproto[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace plugins
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> PROTOBUF_NOINLINE ::nebulaidl::plugins::DistributedMPITrainingTask* Arena::CreateMaybeMessage< ::nebulaidl::plugins::DistributedMPITrainingTask >(Arena* arena) {
  return Arena::CreateInternal< ::nebulaidl::plugins::DistributedMPITrainingTask >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
