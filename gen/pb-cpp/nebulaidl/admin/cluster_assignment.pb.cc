// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/admin/cluster_assignment.proto

#include "nebulaidl/admin/cluster_assignment.pb.h"

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
namespace admin {
class ClusterAssignmentDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<ClusterAssignment> _instance;
} _ClusterAssignment_default_instance_;
}  // namespace admin
}  // namespace nebulaidl
static void InitDefaultsClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::nebulaidl::admin::_ClusterAssignment_default_instance_;
    new (ptr) ::nebulaidl::admin::ClusterAssignment();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::nebulaidl::admin::ClusterAssignment::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<0> scc_info_ClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 0, InitDefaultsClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto}, {}};

void InitDefaults_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto() {
  ::google::protobuf::internal::InitSCC(&scc_info_ClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto.base);
}

::google::protobuf::Metadata file_level_metadata_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto[1];
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto = nullptr;

const ::google::protobuf::uint32 TableStruct_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::admin::ClusterAssignment, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::admin::ClusterAssignment, cluster_pool_name_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::nebulaidl::admin::ClusterAssignment)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::nebulaidl::admin::_ClusterAssignment_default_instance_),
};

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto = {
  {}, AddDescriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, "nebulaidl/admin/cluster_assignment.proto", schemas,
  file_default_instances, TableStruct_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto::offsets,
  file_level_metadata_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, 1, file_level_enum_descriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, file_level_service_descriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto,
};

const char descriptor_table_protodef_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto[] =
  "\n(nebulaidl/admin/cluster_assignment.pro"
  "to\022\017nebulaidl.admin\":\n\021ClusterAssignment"
  "\022\031\n\021cluster_pool_name\030\003 \001(\tJ\004\010\001\020\002J\004\010\002\020\003B"
  "=Z;github.com/nebulaclouds/nebulaidl/gen"
  "/pb-go/nebulaidl/adminb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto = {
  false, InitDefaults_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, 
  descriptor_table_protodef_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto,
  "nebulaidl/admin/cluster_assignment.proto", &assign_descriptors_table_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, 190,
};

void AddDescriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[1] =
  {
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto, deps, 0);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto = []() { AddDescriptors_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto(); return true; }();
namespace nebulaidl {
namespace admin {

// ===================================================================

void ClusterAssignment::InitAsDefaultInstance() {
}
class ClusterAssignment::HasBitSetters {
 public:
};

#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int ClusterAssignment::kClusterPoolNameFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

ClusterAssignment::ClusterAssignment()
  : ::google::protobuf::Message(), _internal_metadata_(nullptr) {
  SharedCtor();
  // @@protoc_insertion_point(constructor:nebulaidl.admin.ClusterAssignment)
}
ClusterAssignment::ClusterAssignment(const ClusterAssignment& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(nullptr) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  cluster_pool_name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.cluster_pool_name().size() > 0) {
    cluster_pool_name_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.cluster_pool_name_);
  }
  // @@protoc_insertion_point(copy_constructor:nebulaidl.admin.ClusterAssignment)
}

void ClusterAssignment::SharedCtor() {
  ::google::protobuf::internal::InitSCC(
      &scc_info_ClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto.base);
  cluster_pool_name_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

ClusterAssignment::~ClusterAssignment() {
  // @@protoc_insertion_point(destructor:nebulaidl.admin.ClusterAssignment)
  SharedDtor();
}

void ClusterAssignment::SharedDtor() {
  cluster_pool_name_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
}

void ClusterAssignment::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const ClusterAssignment& ClusterAssignment::default_instance() {
  ::google::protobuf::internal::InitSCC(&::scc_info_ClusterAssignment_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto.base);
  return *internal_default_instance();
}


void ClusterAssignment::Clear() {
// @@protoc_insertion_point(message_clear_start:nebulaidl.admin.ClusterAssignment)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  cluster_pool_name_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  _internal_metadata_.Clear();
}

#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
const char* ClusterAssignment::_InternalParse(const char* begin, const char* end, void* object,
                  ::google::protobuf::internal::ParseContext* ctx) {
  auto msg = static_cast<ClusterAssignment*>(object);
  ::google::protobuf::int32 size; (void)size;
  int depth; (void)depth;
  ::google::protobuf::uint32 tag;
  ::google::protobuf::internal::ParseFunc parser_till_end; (void)parser_till_end;
  auto ptr = begin;
  while (ptr < end) {
    ptr = ::google::protobuf::io::Parse32(ptr, &tag);
    GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
    switch (tag >> 3) {
      // string cluster_pool_name = 3;
      case 3: {
        if (static_cast<::google::protobuf::uint8>(tag) != 26) goto handle_unusual;
        ptr = ::google::protobuf::io::ReadSize(ptr, &size);
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        ctx->extra_parse_data().SetFieldName("nebulaidl.admin.ClusterAssignment.cluster_pool_name");
        object = msg->mutable_cluster_pool_name();
        if (size > end - ptr + ::google::protobuf::internal::ParseContext::kSlopBytes) {
          parser_till_end = ::google::protobuf::internal::GreedyStringParserUTF8;
          goto string_till_end;
        }
        GOOGLE_PROTOBUF_PARSER_ASSERT(::google::protobuf::internal::StringCheckUTF8(ptr, size, ctx));
        ::google::protobuf::internal::InlineGreedyStringParser(object, ptr, size, ctx);
        ptr += size;
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
string_till_end:
  static_cast<::std::string*>(object)->clear();
  static_cast<::std::string*>(object)->reserve(size);
  goto len_delim_till_end;
len_delim_till_end:
  return ctx->StoreAndTailCall(ptr, end, {_InternalParse, msg},
                               {parser_till_end, object}, size);
}
#else  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
bool ClusterAssignment::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!PROTOBUF_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:nebulaidl.admin.ClusterAssignment)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // string cluster_pool_name = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (26 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_cluster_pool_name()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->cluster_pool_name().data(), static_cast<int>(this->cluster_pool_name().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "nebulaidl.admin.ClusterAssignment.cluster_pool_name"));
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
  // @@protoc_insertion_point(parse_success:nebulaidl.admin.ClusterAssignment)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:nebulaidl.admin.ClusterAssignment)
  return false;
#undef DO_
}
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER

void ClusterAssignment::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:nebulaidl.admin.ClusterAssignment)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string cluster_pool_name = 3;
  if (this->cluster_pool_name().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->cluster_pool_name().data(), static_cast<int>(this->cluster_pool_name().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "nebulaidl.admin.ClusterAssignment.cluster_pool_name");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      3, this->cluster_pool_name(), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:nebulaidl.admin.ClusterAssignment)
}

::google::protobuf::uint8* ClusterAssignment::InternalSerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:nebulaidl.admin.ClusterAssignment)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string cluster_pool_name = 3;
  if (this->cluster_pool_name().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->cluster_pool_name().data(), static_cast<int>(this->cluster_pool_name().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "nebulaidl.admin.ClusterAssignment.cluster_pool_name");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        3, this->cluster_pool_name(), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:nebulaidl.admin.ClusterAssignment)
  return target;
}

size_t ClusterAssignment::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:nebulaidl.admin.ClusterAssignment)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string cluster_pool_name = 3;
  if (this->cluster_pool_name().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->cluster_pool_name());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void ClusterAssignment::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:nebulaidl.admin.ClusterAssignment)
  GOOGLE_DCHECK_NE(&from, this);
  const ClusterAssignment* source =
      ::google::protobuf::DynamicCastToGenerated<ClusterAssignment>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:nebulaidl.admin.ClusterAssignment)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:nebulaidl.admin.ClusterAssignment)
    MergeFrom(*source);
  }
}

void ClusterAssignment::MergeFrom(const ClusterAssignment& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:nebulaidl.admin.ClusterAssignment)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.cluster_pool_name().size() > 0) {

    cluster_pool_name_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.cluster_pool_name_);
  }
}

void ClusterAssignment::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:nebulaidl.admin.ClusterAssignment)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void ClusterAssignment::CopyFrom(const ClusterAssignment& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:nebulaidl.admin.ClusterAssignment)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool ClusterAssignment::IsInitialized() const {
  return true;
}

void ClusterAssignment::Swap(ClusterAssignment* other) {
  if (other == this) return;
  InternalSwap(other);
}
void ClusterAssignment::InternalSwap(ClusterAssignment* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
  cluster_pool_name_.Swap(&other->cluster_pool_name_, &::google::protobuf::internal::GetEmptyStringAlreadyInited(),
    GetArenaNoVirtual());
}

::google::protobuf::Metadata ClusterAssignment::GetMetadata() const {
  ::google::protobuf::internal::AssignDescriptors(&::assign_descriptors_table_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto);
  return ::file_level_metadata_nebulaidl_2fadmin_2fcluster_5fassignment_2eproto[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace admin
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> PROTOBUF_NOINLINE ::nebulaidl::admin::ClusterAssignment* Arena::CreateMaybeMessage< ::nebulaidl::admin::ClusterAssignment >(Arena* arena) {
  return Arena::CreateInternal< ::nebulaidl::admin::ClusterAssignment >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
