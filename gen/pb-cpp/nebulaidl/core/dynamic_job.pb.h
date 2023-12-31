// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/core/dynamic_job.proto

#ifndef PROTOBUF_INCLUDED_nebulaidl_2fcore_2fdynamic_5fjob_2eproto
#define PROTOBUF_INCLUDED_nebulaidl_2fcore_2fdynamic_5fjob_2eproto

#include <limits>
#include <string>

#include <google/protobuf/port_def.inc>
#if PROTOBUF_VERSION < 3007000
#error This file was generated by a newer version of protoc which is
#error incompatible with your Protocol Buffer headers. Please update
#error your headers.
#endif
#if 3007000 < PROTOBUF_MIN_PROTOC_VERSION
#error This file was generated by an older version of protoc which is
#error incompatible with your Protocol Buffer headers. Please
#error regenerate this file with a newer version of protoc.
#endif

#include <google/protobuf/port_undef.inc>
#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/arena.h>
#include <google/protobuf/arenastring.h>
#include <google/protobuf/generated_message_table_driven.h>
#include <google/protobuf/generated_message_util.h>
#include <google/protobuf/inlined_string_field.h>
#include <google/protobuf/metadata.h>
#include <google/protobuf/message.h>
#include <google/protobuf/repeated_field.h>  // IWYU pragma: export
#include <google/protobuf/extension_set.h>  // IWYU pragma: export
#include <google/protobuf/unknown_field_set.h>
#include "nebulaidl/core/tasks.pb.h"
#include "nebulaidl/core/workflow.pb.h"
#include "nebulaidl/core/literals.pb.h"
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>
#define PROTOBUF_INTERNAL_EXPORT_nebulaidl_2fcore_2fdynamic_5fjob_2eproto

// Internal implementation detail -- do not use these members.
struct TableStruct_nebulaidl_2fcore_2fdynamic_5fjob_2eproto {
  static const ::google::protobuf::internal::ParseTableField entries[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::AuxillaryParseTableField aux[]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::ParseTable schema[1]
    PROTOBUF_SECTION_VARIABLE(protodesc_cold);
  static const ::google::protobuf::internal::FieldMetadata field_metadata[];
  static const ::google::protobuf::internal::SerializationTable serialization_table[];
  static const ::google::protobuf::uint32 offsets[];
};
void AddDescriptors_nebulaidl_2fcore_2fdynamic_5fjob_2eproto();
namespace nebulaidl {
namespace core {
class DynamicJobSpec;
class DynamicJobSpecDefaultTypeInternal;
extern DynamicJobSpecDefaultTypeInternal _DynamicJobSpec_default_instance_;
}  // namespace core
}  // namespace nebulaidl
namespace google {
namespace protobuf {
template<> ::nebulaidl::core::DynamicJobSpec* Arena::CreateMaybeMessage<::nebulaidl::core::DynamicJobSpec>(Arena*);
}  // namespace protobuf
}  // namespace google
namespace nebulaidl {
namespace core {

// ===================================================================

class DynamicJobSpec final :
    public ::google::protobuf::Message /* @@protoc_insertion_point(class_definition:nebulaidl.core.DynamicJobSpec) */ {
 public:
  DynamicJobSpec();
  virtual ~DynamicJobSpec();

  DynamicJobSpec(const DynamicJobSpec& from);

  inline DynamicJobSpec& operator=(const DynamicJobSpec& from) {
    CopyFrom(from);
    return *this;
  }
  #if LANG_CXX11
  DynamicJobSpec(DynamicJobSpec&& from) noexcept
    : DynamicJobSpec() {
    *this = ::std::move(from);
  }

  inline DynamicJobSpec& operator=(DynamicJobSpec&& from) noexcept {
    if (GetArenaNoVirtual() == from.GetArenaNoVirtual()) {
      if (this != &from) InternalSwap(&from);
    } else {
      CopyFrom(from);
    }
    return *this;
  }
  #endif
  static const ::google::protobuf::Descriptor* descriptor() {
    return default_instance().GetDescriptor();
  }
  static const DynamicJobSpec& default_instance();

  static void InitAsDefaultInstance();  // FOR INTERNAL USE ONLY
  static inline const DynamicJobSpec* internal_default_instance() {
    return reinterpret_cast<const DynamicJobSpec*>(
               &_DynamicJobSpec_default_instance_);
  }
  static constexpr int kIndexInFileMessages =
    0;

  void Swap(DynamicJobSpec* other);
  friend void swap(DynamicJobSpec& a, DynamicJobSpec& b) {
    a.Swap(&b);
  }

  // implements Message ----------------------------------------------

  inline DynamicJobSpec* New() const final {
    return CreateMaybeMessage<DynamicJobSpec>(nullptr);
  }

  DynamicJobSpec* New(::google::protobuf::Arena* arena) const final {
    return CreateMaybeMessage<DynamicJobSpec>(arena);
  }
  void CopyFrom(const ::google::protobuf::Message& from) final;
  void MergeFrom(const ::google::protobuf::Message& from) final;
  void CopyFrom(const DynamicJobSpec& from);
  void MergeFrom(const DynamicJobSpec& from);
  PROTOBUF_ATTRIBUTE_REINITIALIZES void Clear() final;
  bool IsInitialized() const final;

  size_t ByteSizeLong() const final;
  #if GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  static const char* _InternalParse(const char* begin, const char* end, void* object, ::google::protobuf::internal::ParseContext* ctx);
  ::google::protobuf::internal::ParseFunc _ParseFunc() const final { return _InternalParse; }
  #else
  bool MergePartialFromCodedStream(
      ::google::protobuf::io::CodedInputStream* input) final;
  #endif  // GOOGLE_PROTOBUF_ENABLE_EXPERIMENTAL_PARSER
  void SerializeWithCachedSizes(
      ::google::protobuf::io::CodedOutputStream* output) const final;
  ::google::protobuf::uint8* InternalSerializeWithCachedSizesToArray(
      ::google::protobuf::uint8* target) const final;
  int GetCachedSize() const final { return _cached_size_.Get(); }

  private:
  void SharedCtor();
  void SharedDtor();
  void SetCachedSize(int size) const final;
  void InternalSwap(DynamicJobSpec* other);
  private:
  inline ::google::protobuf::Arena* GetArenaNoVirtual() const {
    return nullptr;
  }
  inline void* MaybeArenaPtr() const {
    return nullptr;
  }
  public:

  ::google::protobuf::Metadata GetMetadata() const final;

  // nested types ----------------------------------------------------

  // accessors -------------------------------------------------------

  // repeated .nebulaidl.core.Node nodes = 1;
  int nodes_size() const;
  void clear_nodes();
  static const int kNodesFieldNumber = 1;
  ::nebulaidl::core::Node* mutable_nodes(int index);
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Node >*
      mutable_nodes();
  const ::nebulaidl::core::Node& nodes(int index) const;
  ::nebulaidl::core::Node* add_nodes();
  const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Node >&
      nodes() const;

  // repeated .nebulaidl.core.Binding outputs = 3;
  int outputs_size() const;
  void clear_outputs();
  static const int kOutputsFieldNumber = 3;
  ::nebulaidl::core::Binding* mutable_outputs(int index);
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Binding >*
      mutable_outputs();
  const ::nebulaidl::core::Binding& outputs(int index) const;
  ::nebulaidl::core::Binding* add_outputs();
  const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Binding >&
      outputs() const;

  // repeated .nebulaidl.core.TaskTemplate tasks = 4;
  int tasks_size() const;
  void clear_tasks();
  static const int kTasksFieldNumber = 4;
  ::nebulaidl::core::TaskTemplate* mutable_tasks(int index);
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::TaskTemplate >*
      mutable_tasks();
  const ::nebulaidl::core::TaskTemplate& tasks(int index) const;
  ::nebulaidl::core::TaskTemplate* add_tasks();
  const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::TaskTemplate >&
      tasks() const;

  // repeated .nebulaidl.core.WorkflowTemplate subworkflows = 5;
  int subworkflows_size() const;
  void clear_subworkflows();
  static const int kSubworkflowsFieldNumber = 5;
  ::nebulaidl::core::WorkflowTemplate* mutable_subworkflows(int index);
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::WorkflowTemplate >*
      mutable_subworkflows();
  const ::nebulaidl::core::WorkflowTemplate& subworkflows(int index) const;
  ::nebulaidl::core::WorkflowTemplate* add_subworkflows();
  const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::WorkflowTemplate >&
      subworkflows() const;

  // int64 min_successes = 2;
  void clear_min_successes();
  static const int kMinSuccessesFieldNumber = 2;
  ::google::protobuf::int64 min_successes() const;
  void set_min_successes(::google::protobuf::int64 value);

  // @@protoc_insertion_point(class_scope:nebulaidl.core.DynamicJobSpec)
 private:
  class HasBitSetters;

  ::google::protobuf::internal::InternalMetadataWithArena _internal_metadata_;
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Node > nodes_;
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Binding > outputs_;
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::TaskTemplate > tasks_;
  ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::WorkflowTemplate > subworkflows_;
  ::google::protobuf::int64 min_successes_;
  mutable ::google::protobuf::internal::CachedSize _cached_size_;
  friend struct ::TableStruct_nebulaidl_2fcore_2fdynamic_5fjob_2eproto;
};
// ===================================================================


// ===================================================================

#ifdef __GNUC__
  #pragma GCC diagnostic push
  #pragma GCC diagnostic ignored "-Wstrict-aliasing"
#endif  // __GNUC__
// DynamicJobSpec

// repeated .nebulaidl.core.Node nodes = 1;
inline int DynamicJobSpec::nodes_size() const {
  return nodes_.size();
}
inline ::nebulaidl::core::Node* DynamicJobSpec::mutable_nodes(int index) {
  // @@protoc_insertion_point(field_mutable:nebulaidl.core.DynamicJobSpec.nodes)
  return nodes_.Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Node >*
DynamicJobSpec::mutable_nodes() {
  // @@protoc_insertion_point(field_mutable_list:nebulaidl.core.DynamicJobSpec.nodes)
  return &nodes_;
}
inline const ::nebulaidl::core::Node& DynamicJobSpec::nodes(int index) const {
  // @@protoc_insertion_point(field_get:nebulaidl.core.DynamicJobSpec.nodes)
  return nodes_.Get(index);
}
inline ::nebulaidl::core::Node* DynamicJobSpec::add_nodes() {
  // @@protoc_insertion_point(field_add:nebulaidl.core.DynamicJobSpec.nodes)
  return nodes_.Add();
}
inline const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Node >&
DynamicJobSpec::nodes() const {
  // @@protoc_insertion_point(field_list:nebulaidl.core.DynamicJobSpec.nodes)
  return nodes_;
}

// int64 min_successes = 2;
inline void DynamicJobSpec::clear_min_successes() {
  min_successes_ = PROTOBUF_LONGLONG(0);
}
inline ::google::protobuf::int64 DynamicJobSpec::min_successes() const {
  // @@protoc_insertion_point(field_get:nebulaidl.core.DynamicJobSpec.min_successes)
  return min_successes_;
}
inline void DynamicJobSpec::set_min_successes(::google::protobuf::int64 value) {
  
  min_successes_ = value;
  // @@protoc_insertion_point(field_set:nebulaidl.core.DynamicJobSpec.min_successes)
}

// repeated .nebulaidl.core.Binding outputs = 3;
inline int DynamicJobSpec::outputs_size() const {
  return outputs_.size();
}
inline ::nebulaidl::core::Binding* DynamicJobSpec::mutable_outputs(int index) {
  // @@protoc_insertion_point(field_mutable:nebulaidl.core.DynamicJobSpec.outputs)
  return outputs_.Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Binding >*
DynamicJobSpec::mutable_outputs() {
  // @@protoc_insertion_point(field_mutable_list:nebulaidl.core.DynamicJobSpec.outputs)
  return &outputs_;
}
inline const ::nebulaidl::core::Binding& DynamicJobSpec::outputs(int index) const {
  // @@protoc_insertion_point(field_get:nebulaidl.core.DynamicJobSpec.outputs)
  return outputs_.Get(index);
}
inline ::nebulaidl::core::Binding* DynamicJobSpec::add_outputs() {
  // @@protoc_insertion_point(field_add:nebulaidl.core.DynamicJobSpec.outputs)
  return outputs_.Add();
}
inline const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::Binding >&
DynamicJobSpec::outputs() const {
  // @@protoc_insertion_point(field_list:nebulaidl.core.DynamicJobSpec.outputs)
  return outputs_;
}

// repeated .nebulaidl.core.TaskTemplate tasks = 4;
inline int DynamicJobSpec::tasks_size() const {
  return tasks_.size();
}
inline ::nebulaidl::core::TaskTemplate* DynamicJobSpec::mutable_tasks(int index) {
  // @@protoc_insertion_point(field_mutable:nebulaidl.core.DynamicJobSpec.tasks)
  return tasks_.Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::TaskTemplate >*
DynamicJobSpec::mutable_tasks() {
  // @@protoc_insertion_point(field_mutable_list:nebulaidl.core.DynamicJobSpec.tasks)
  return &tasks_;
}
inline const ::nebulaidl::core::TaskTemplate& DynamicJobSpec::tasks(int index) const {
  // @@protoc_insertion_point(field_get:nebulaidl.core.DynamicJobSpec.tasks)
  return tasks_.Get(index);
}
inline ::nebulaidl::core::TaskTemplate* DynamicJobSpec::add_tasks() {
  // @@protoc_insertion_point(field_add:nebulaidl.core.DynamicJobSpec.tasks)
  return tasks_.Add();
}
inline const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::TaskTemplate >&
DynamicJobSpec::tasks() const {
  // @@protoc_insertion_point(field_list:nebulaidl.core.DynamicJobSpec.tasks)
  return tasks_;
}

// repeated .nebulaidl.core.WorkflowTemplate subworkflows = 5;
inline int DynamicJobSpec::subworkflows_size() const {
  return subworkflows_.size();
}
inline ::nebulaidl::core::WorkflowTemplate* DynamicJobSpec::mutable_subworkflows(int index) {
  // @@protoc_insertion_point(field_mutable:nebulaidl.core.DynamicJobSpec.subworkflows)
  return subworkflows_.Mutable(index);
}
inline ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::WorkflowTemplate >*
DynamicJobSpec::mutable_subworkflows() {
  // @@protoc_insertion_point(field_mutable_list:nebulaidl.core.DynamicJobSpec.subworkflows)
  return &subworkflows_;
}
inline const ::nebulaidl::core::WorkflowTemplate& DynamicJobSpec::subworkflows(int index) const {
  // @@protoc_insertion_point(field_get:nebulaidl.core.DynamicJobSpec.subworkflows)
  return subworkflows_.Get(index);
}
inline ::nebulaidl::core::WorkflowTemplate* DynamicJobSpec::add_subworkflows() {
  // @@protoc_insertion_point(field_add:nebulaidl.core.DynamicJobSpec.subworkflows)
  return subworkflows_.Add();
}
inline const ::google::protobuf::RepeatedPtrField< ::nebulaidl::core::WorkflowTemplate >&
DynamicJobSpec::subworkflows() const {
  // @@protoc_insertion_point(field_list:nebulaidl.core.DynamicJobSpec.subworkflows)
  return subworkflows_;
}

#ifdef __GNUC__
  #pragma GCC diagnostic pop
#endif  // __GNUC__

// @@protoc_insertion_point(namespace_scope)

}  // namespace core
}  // namespace nebulaidl

// @@protoc_insertion_point(global_scope)

#include <google/protobuf/port_undef.inc>
#endif  // PROTOBUF_INCLUDED_nebulaidl_2fcore_2fdynamic_5fjob_2eproto
