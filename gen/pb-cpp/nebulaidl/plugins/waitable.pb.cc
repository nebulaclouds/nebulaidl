// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/plugins/waitable.proto

#include "nebulaidl/plugins/waitable.pb.h"

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

extern PROTOBUF_INTERNAL_EXPORT_nebulaidl_2fcore_2fidentifier_2eproto ::google::protobuf::internal::SCCInfo<0> scc_info_WorkflowExecutionIdentifier_nebulaidl_2fcore_2fidentifier_2eproto;
namespace nebulaidl {
namespace plugins {
class WaitableDefaultTypeInternal {
 public:
  ::google::protobuf::internal::ExplicitlyConstructed<Waitable> _instance;
} _Waitable_default_instance_;
}  // namespace plugins
}  // namespace nebulaidl
static void InitDefaultsWaitable_nebulaidl_2fplugins_2fwaitable_2eproto() {
  GOOGLE_PROTOBUF_VERIFY_VERSION;

  {
    void* ptr = &::nebulaidl::plugins::_Waitable_default_instance_;
    new (ptr) ::nebulaidl::plugins::Waitable();
    ::google::protobuf::internal::OnShutdownDestroyMessage(ptr);
  }
  ::nebulaidl::plugins::Waitable::InitAsDefaultInstance();
}

::google::protobuf::internal::SCCInfo<1> scc_info_Waitable_nebulaidl_2fplugins_2fwaitable_2eproto =
    {{ATOMIC_VAR_INIT(::google::protobuf::internal::SCCInfoBase::kUninitialized), 1, InitDefaultsWaitable_nebulaidl_2fplugins_2fwaitable_2eproto}, {
      &scc_info_WorkflowExecutionIdentifier_nebulaidl_2fcore_2fidentifier_2eproto.base,}};

void InitDefaults_nebulaidl_2fplugins_2fwaitable_2eproto() {
  ::google::protobuf::internal::InitSCC(&scc_info_Waitable_nebulaidl_2fplugins_2fwaitable_2eproto.base);
}

::google::protobuf::Metadata file_level_metadata_nebulaidl_2fplugins_2fwaitable_2eproto[1];
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_nebulaidl_2fplugins_2fwaitable_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_nebulaidl_2fplugins_2fwaitable_2eproto = nullptr;

const ::google::protobuf::uint32 TableStruct_nebulaidl_2fplugins_2fwaitable_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::Waitable, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::Waitable, wf_exec_id_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::Waitable, phase_),
  PROTOBUF_FIELD_OFFSET(::nebulaidl::plugins::Waitable, workflow_id_),
};
static const ::google::protobuf::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::nebulaidl::plugins::Waitable)},
};

static ::google::protobuf::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::google::protobuf::Message*>(&::nebulaidl::plugins::_Waitable_default_instance_),
};

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_nebulaidl_2fplugins_2fwaitable_2eproto = {
  {}, AddDescriptors_nebulaidl_2fplugins_2fwaitable_2eproto, "nebulaidl/plugins/waitable.proto", schemas,
  file_default_instances, TableStruct_nebulaidl_2fplugins_2fwaitable_2eproto::offsets,
  file_level_metadata_nebulaidl_2fplugins_2fwaitable_2eproto, 1, file_level_enum_descriptors_nebulaidl_2fplugins_2fwaitable_2eproto, file_level_service_descriptors_nebulaidl_2fplugins_2fwaitable_2eproto,
};

const char descriptor_table_protodef_nebulaidl_2fplugins_2fwaitable_2eproto[] =
  "\n nebulaidl/plugins/waitable.proto\022\021nebu"
  "laidl.plugins\032\036nebulaidl/core/execution."
  "proto\032\037nebulaidl/core/identifier.proto\"\230"
  "\001\n\010Waitable\022\?\n\nwf_exec_id\030\001 \001(\0132+.nebula"
  "idl.core.WorkflowExecutionIdentifier\0226\n\005"
  "phase\030\002 \001(\0162\'.nebulaidl.core.WorkflowExe"
  "cution.Phase\022\023\n\013workflow_id\030\003 \001(\tB\?Z=git"
  "hub.com/nebulaclouds/nebulaidl/gen/pb-go"
  "/nebulaidl/pluginsb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_nebulaidl_2fplugins_2fwaitable_2eproto = {
  false, InitDefaults_nebulaidl_2fplugins_2fwaitable_2eproto, 
  descriptor_table_protodef_nebulaidl_2fplugins_2fwaitable_2eproto,
  "nebulaidl/plugins/waitable.proto", &assign_descriptors_table_nebulaidl_2fplugins_2fwaitable_2eproto, 346,
};

void AddDescriptors_nebulaidl_2fplugins_2fwaitable_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[2] =
  {
    ::AddDescriptors_nebulaidl_2fcore_2fexecution_2eproto,
    ::AddDescriptors_nebulaidl_2fcore_2fidentifier_2eproto,
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_nebulaidl_2fplugins_2fwaitable_2eproto, deps, 2);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_nebulaidl_2fplugins_2fwaitable_2eproto = []() { AddDescriptors_nebulaidl_2fplugins_2fwaitable_2eproto(); return true; }();
namespace nebulaidl {
namespace plugins {

// ===================================================================

void Waitable::InitAsDefaultInstance() {
  ::nebulaidl::plugins::_Waitable_default_instance_._instance.get_mutable()->wf_exec_id_ = const_cast< ::nebulaidl::core::WorkflowExecutionIdentifier*>(
      ::nebulaidl::core::WorkflowExecutionIdentifier::internal_default_instance());
}
class Waitable::HasBitSetters {
 public:
  static const ::nebulaidl::core::WorkflowExecutionIdentifier& wf_exec_id(const Waitable* msg);
};

const ::nebulaidl::core::WorkflowExecutionIdentifier&
Waitable::HasBitSetters::wf_exec_id(const Waitable* msg) {
  return *msg->wf_exec_id_;
}
void Waitable::clear_wf_exec_id() {
  if (GetArenaNoVirtual() == nullptr && wf_exec_id_ != nullptr) {
    delete wf_exec_id_;
  }
  wf_exec_id_ = nullptr;
}
#if !defined(_MSC_VER) || _MSC_VER >= 1900
const int Waitable::kWfExecIdFieldNumber;
const int Waitable::kPhaseFieldNumber;
const int Waitable::kWorkflowIdFieldNumber;
#endif  // !defined(_MSC_VER) || _MSC_VER >= 1900

Waitable::Waitable()
  : ::google::protobuf::Message(), _internal_metadata_(nullptr) {
  SharedCtor();
  // @@protoc_insertion_point(constructor:nebulaidl.plugins.Waitable)
}
Waitable::Waitable(const Waitable& from)
  : ::google::protobuf::Message(),
      _internal_metadata_(nullptr) {
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  workflow_id_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (from.workflow_id().size() > 0) {
    workflow_id_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.workflow_id_);
  }
  if (from.has_wf_exec_id()) {
    wf_exec_id_ = new ::nebulaidl::core::WorkflowExecutionIdentifier(*from.wf_exec_id_);
  } else {
    wf_exec_id_ = nullptr;
  }
  phase_ = from.phase_;
  // @@protoc_insertion_point(copy_constructor:nebulaidl.plugins.Waitable)
}

void Waitable::SharedCtor() {
  ::google::protobuf::internal::InitSCC(
      &scc_info_Waitable_nebulaidl_2fplugins_2fwaitable_2eproto.base);
  workflow_id_.UnsafeSetDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  ::memset(&wf_exec_id_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&phase_) -
      reinterpret_cast<char*>(&wf_exec_id_)) + sizeof(phase_));
}

Waitable::~Waitable() {
  // @@protoc_insertion_point(destructor:nebulaidl.plugins.Waitable)
  SharedDtor();
}

void Waitable::SharedDtor() {
  workflow_id_.DestroyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (this != internal_default_instance()) delete wf_exec_id_;
}

void Waitable::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}
const Waitable& Waitable::default_instance() {
  ::google::protobuf::internal::InitSCC(&::scc_info_Waitable_nebulaidl_2fplugins_2fwaitable_2eproto.base);
  return *internal_default_instance();
}


void Waitable::Clear() {
// @@protoc_insertion_point(message_clear_start:nebulaidl.plugins.Waitable)
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  workflow_id_.ClearToEmptyNoArena(&::google::protobuf::internal::GetEmptyStringAlreadyInited());
  if (GetArenaNoVirtual() == nullptr && wf_exec_id_ != nullptr) {
    delete wf_exec_id_;
  }
  wf_exec_id_ = nullptr;
  phase_ = 0;
  _internal_metadata_.Clear();
}

#if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
const char* Waitable::_InternalParse(const char* begin, const char* end, void* object,
                  ::google::protobuf::internal::ParseContext* ctx) {
  auto msg = static_cast<Waitable*>(object);
  ::google::protobuf::int32 size; (void)size;
  int depth; (void)depth;
  ::google::protobuf::uint32 tag;
  ::google::protobuf::internal::ParseFunc parser_till_end; (void)parser_till_end;
  auto ptr = begin;
  while (ptr < end) {
    ptr = ::google::protobuf::io::Parse32(ptr, &tag);
    GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
    switch (tag >> 3) {
      // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
      case 1: {
        if (static_cast<::google::protobuf::uint8>(tag) != 10) goto handle_unusual;
        ptr = ::google::protobuf::io::ReadSize(ptr, &size);
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        parser_till_end = ::nebulaidl::core::WorkflowExecutionIdentifier::_InternalParse;
        object = msg->mutable_wf_exec_id();
        if (size > end - ptr) goto len_delim_till_end;
        ptr += size;
        GOOGLE_PROTOBUF_PARSER_ASSERT(ctx->ParseExactRange(
            {parser_till_end, object}, ptr - size, ptr));
        break;
      }
      // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
      case 2: {
        if (static_cast<::google::protobuf::uint8>(tag) != 16) goto handle_unusual;
        ::google::protobuf::uint64 val = ::google::protobuf::internal::ReadVarint(&ptr);
        msg->set_phase(static_cast<::nebulaidl::core::WorkflowExecution_Phase>(val));
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        break;
      }
      // string workflow_id = 3;
      case 3: {
        if (static_cast<::google::protobuf::uint8>(tag) != 26) goto handle_unusual;
        ptr = ::google::protobuf::io::ReadSize(ptr, &size);
        GOOGLE_PROTOBUF_PARSER_ASSERT(ptr);
        ctx->extra_parse_data().SetFieldName("nebulaidl.plugins.Waitable.workflow_id");
        object = msg->mutable_workflow_id();
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
bool Waitable::MergePartialFromCodedStream(
    ::google::protobuf::io::CodedInputStream* input) {
#define DO_(EXPRESSION) if (!PROTOBUF_PREDICT_TRUE(EXPRESSION)) goto failure
  ::google::protobuf::uint32 tag;
  // @@protoc_insertion_point(parse_start:nebulaidl.plugins.Waitable)
  for (;;) {
    ::std::pair<::google::protobuf::uint32, bool> p = input->ReadTagWithCutoffNoLastTag(127u);
    tag = p.first;
    if (!p.second) goto handle_unusual;
    switch (::google::protobuf::internal::WireFormatLite::GetTagFieldNumber(tag)) {
      // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
      case 1: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (10 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadMessage(
               input, mutable_wf_exec_id()));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
      case 2: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (16 & 0xFF)) {
          int value = 0;
          DO_((::google::protobuf::internal::WireFormatLite::ReadPrimitive<
                   int, ::google::protobuf::internal::WireFormatLite::TYPE_ENUM>(
                 input, &value)));
          set_phase(static_cast< ::nebulaidl::core::WorkflowExecution_Phase >(value));
        } else {
          goto handle_unusual;
        }
        break;
      }

      // string workflow_id = 3;
      case 3: {
        if (static_cast< ::google::protobuf::uint8>(tag) == (26 & 0xFF)) {
          DO_(::google::protobuf::internal::WireFormatLite::ReadString(
                input, this->mutable_workflow_id()));
          DO_(::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
            this->workflow_id().data(), static_cast<int>(this->workflow_id().length()),
            ::google::protobuf::internal::WireFormatLite::PARSE,
            "nebulaidl.plugins.Waitable.workflow_id"));
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
  // @@protoc_insertion_point(parse_success:nebulaidl.plugins.Waitable)
  return true;
failure:
  // @@protoc_insertion_point(parse_failure:nebulaidl.plugins.Waitable)
  return false;
#undef DO_
}
#endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER

void Waitable::SerializeWithCachedSizes(
    ::google::protobuf::io::CodedOutputStream* output) const {
  // @@protoc_insertion_point(serialize_start:nebulaidl.plugins.Waitable)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
  if (this->has_wf_exec_id()) {
    ::google::protobuf::internal::WireFormatLite::WriteMessageMaybeToArray(
      1, HasBitSetters::wf_exec_id(this), output);
  }

  // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
  if (this->phase() != 0) {
    ::google::protobuf::internal::WireFormatLite::WriteEnum(
      2, this->phase(), output);
  }

  // string workflow_id = 3;
  if (this->workflow_id().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->workflow_id().data(), static_cast<int>(this->workflow_id().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "nebulaidl.plugins.Waitable.workflow_id");
    ::google::protobuf::internal::WireFormatLite::WriteStringMaybeAliased(
      3, this->workflow_id(), output);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    ::google::protobuf::internal::WireFormat::SerializeUnknownFields(
        _internal_metadata_.unknown_fields(), output);
  }
  // @@protoc_insertion_point(serialize_end:nebulaidl.plugins.Waitable)
}

::google::protobuf::uint8* Waitable::InternalSerializeWithCachedSizesToArray(
    ::google::protobuf::uint8* target) const {
  // @@protoc_insertion_point(serialize_to_array_start:nebulaidl.plugins.Waitable)
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
  if (this->has_wf_exec_id()) {
    target = ::google::protobuf::internal::WireFormatLite::
      InternalWriteMessageToArray(
        1, HasBitSetters::wf_exec_id(this), target);
  }

  // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
  if (this->phase() != 0) {
    target = ::google::protobuf::internal::WireFormatLite::WriteEnumToArray(
      2, this->phase(), target);
  }

  // string workflow_id = 3;
  if (this->workflow_id().size() > 0) {
    ::google::protobuf::internal::WireFormatLite::VerifyUtf8String(
      this->workflow_id().data(), static_cast<int>(this->workflow_id().length()),
      ::google::protobuf::internal::WireFormatLite::SERIALIZE,
      "nebulaidl.plugins.Waitable.workflow_id");
    target =
      ::google::protobuf::internal::WireFormatLite::WriteStringToArray(
        3, this->workflow_id(), target);
  }

  if (_internal_metadata_.have_unknown_fields()) {
    target = ::google::protobuf::internal::WireFormat::SerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields(), target);
  }
  // @@protoc_insertion_point(serialize_to_array_end:nebulaidl.plugins.Waitable)
  return target;
}

size_t Waitable::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:nebulaidl.plugins.Waitable)
  size_t total_size = 0;

  if (_internal_metadata_.have_unknown_fields()) {
    total_size +=
      ::google::protobuf::internal::WireFormat::ComputeUnknownFieldsSize(
        _internal_metadata_.unknown_fields());
  }
  ::google::protobuf::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string workflow_id = 3;
  if (this->workflow_id().size() > 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::StringSize(
        this->workflow_id());
  }

  // .nebulaidl.core.WorkflowExecutionIdentifier wf_exec_id = 1;
  if (this->has_wf_exec_id()) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::MessageSize(
        *wf_exec_id_);
  }

  // .nebulaidl.core.WorkflowExecution.Phase phase = 2;
  if (this->phase() != 0) {
    total_size += 1 +
      ::google::protobuf::internal::WireFormatLite::EnumSize(this->phase());
  }

  int cached_size = ::google::protobuf::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void Waitable::MergeFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:nebulaidl.plugins.Waitable)
  GOOGLE_DCHECK_NE(&from, this);
  const Waitable* source =
      ::google::protobuf::DynamicCastToGenerated<Waitable>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:nebulaidl.plugins.Waitable)
    ::google::protobuf::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:nebulaidl.plugins.Waitable)
    MergeFrom(*source);
  }
}

void Waitable::MergeFrom(const Waitable& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:nebulaidl.plugins.Waitable)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom(from._internal_metadata_);
  ::google::protobuf::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.workflow_id().size() > 0) {

    workflow_id_.AssignWithDefault(&::google::protobuf::internal::GetEmptyStringAlreadyInited(), from.workflow_id_);
  }
  if (from.has_wf_exec_id()) {
    mutable_wf_exec_id()->::nebulaidl::core::WorkflowExecutionIdentifier::MergeFrom(from.wf_exec_id());
  }
  if (from.phase() != 0) {
    set_phase(from.phase());
  }
}

void Waitable::CopyFrom(const ::google::protobuf::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:nebulaidl.plugins.Waitable)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void Waitable::CopyFrom(const Waitable& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:nebulaidl.plugins.Waitable)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Waitable::IsInitialized() const {
  return true;
}

void Waitable::Swap(Waitable* other) {
  if (other == this) return;
  InternalSwap(other);
}
void Waitable::InternalSwap(Waitable* other) {
  using std::swap;
  _internal_metadata_.Swap(&other->_internal_metadata_);
  workflow_id_.Swap(&other->workflow_id_, &::google::protobuf::internal::GetEmptyStringAlreadyInited(),
    GetArenaNoVirtual());
  swap(wf_exec_id_, other->wf_exec_id_);
  swap(phase_, other->phase_);
}

::google::protobuf::Metadata Waitable::GetMetadata() const {
  ::google::protobuf::internal::AssignDescriptors(&::assign_descriptors_table_nebulaidl_2fplugins_2fwaitable_2eproto);
  return ::file_level_metadata_nebulaidl_2fplugins_2fwaitable_2eproto[kIndexInFileMessages];
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace plugins
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> PROTOBUF_NOINLINE ::nebulaidl::plugins::Waitable* Arena::CreateMaybeMessage< ::nebulaidl::plugins::Waitable >(Arena* arena) {
  return Arena::CreateInternal< ::nebulaidl::plugins::Waitable >(arena);
}
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
