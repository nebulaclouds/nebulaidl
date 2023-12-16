// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: nebulaidl/service/admin.proto

#include "nebulaidl/service/admin.pb.h"

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
namespace service {
}  // namespace service
}  // namespace nebulaidl
void InitDefaults_nebulaidl_2fservice_2fadmin_2eproto() {
}

constexpr ::google::protobuf::Metadata* file_level_metadata_nebulaidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::EnumDescriptor const** file_level_enum_descriptors_nebulaidl_2fservice_2fadmin_2eproto = nullptr;
constexpr ::google::protobuf::ServiceDescriptor const** file_level_service_descriptors_nebulaidl_2fservice_2fadmin_2eproto = nullptr;
const ::google::protobuf::uint32 TableStruct_nebulaidl_2fservice_2fadmin_2eproto::offsets[1] = {};
static constexpr ::google::protobuf::internal::MigrationSchema* schemas = nullptr;
static constexpr ::google::protobuf::Message* const* file_default_instances = nullptr;

::google::protobuf::internal::AssignDescriptorsTable assign_descriptors_table_nebulaidl_2fservice_2fadmin_2eproto = {
  {}, AddDescriptors_nebulaidl_2fservice_2fadmin_2eproto, "nebulaidl/service/admin.proto", schemas,
  file_default_instances, TableStruct_nebulaidl_2fservice_2fadmin_2eproto::offsets,
  file_level_metadata_nebulaidl_2fservice_2fadmin_2eproto, 0, file_level_enum_descriptors_nebulaidl_2fservice_2fadmin_2eproto, file_level_service_descriptors_nebulaidl_2fservice_2fadmin_2eproto,
};

const char descriptor_table_protodef_nebulaidl_2fservice_2fadmin_2eproto[] =
  "\n\034nebulaidl/service/admin.proto\022\020nebulaidl"
  ".service\032\034google/api/annotations.proto\032\034"
  "nebulaidl/admin/project.proto\032.nebulaidl/a"
  "dmin/project_domain_attributes.proto\032\'fl"
  "yteidl/admin/project_attributes.proto\032\031f"
  "lyteidl/admin/task.proto\032\035nebulaidl/admin"
  "/workflow.proto\032(nebulaidl/admin/workflow"
  "_attributes.proto\032 nebulaidl/admin/launch"
  "_plan.proto\032\032nebulaidl/admin/event.proto\032"
  "\036nebulaidl/admin/execution.proto\032\'nebulaid"
  "l/admin/matchable_resource.proto\032#nebulai"
  "dl/admin/node_execution.proto\032#nebulaidl/"
  "admin/task_execution.proto\032\034nebulaidl/adm"
  "in/version.proto\032\033nebulaidl/admin/common."
  "proto\032\'nebulaidl/admin/description_entity"
  ".proto2\204N\n\014AdminService\022m\n\nCreateTask\022!."
  "nebulaidl.admin.TaskCreateRequest\032\".nebula"
  "idl.admin.TaskCreateResponse\"\030\202\323\344\223\002\022\"\r/a"
  "pi/v1/tasks:\001*\022\210\001\n\007GetTask\022 .nebulaidl.ad"
  "min.ObjectGetRequest\032\024.nebulaidl.admin.Ta"
  "sk\"E\202\323\344\223\002\?\022=/api/v1/tasks/{id.project}/{"
  "id.domain}/{id.name}/{id.version}\022\227\001\n\013Li"
  "stTaskIds\0220.nebulaidl.admin.NamedEntityId"
  "entifierListRequest\032).nebulaidl.admin.Nam"
  "edEntityIdentifierList\"+\202\323\344\223\002%\022#/api/v1/"
  "task_ids/{project}/{domain}\022\256\001\n\tListTask"
  "s\022#.nebulaidl.admin.ResourceListRequest\032\030"
  ".nebulaidl.admin.TaskList\"b\202\323\344\223\002\\\0220/api/v"
  "1/tasks/{id.project}/{id.domain}/{id.nam"
  "e}Z(\022&/api/v1/tasks/{id.project}/{id.dom"
  "ain}\022}\n\016CreateWorkflow\022%.nebulaidl.admin."
  "WorkflowCreateRequest\032&.nebulaidl.admin.W"
  "orkflowCreateResponse\"\034\202\323\344\223\002\026\"\021/api/v1/w"
  "orkflows:\001*\022\224\001\n\013GetWorkflow\022 .nebulaidl.a"
  "dmin.ObjectGetRequest\032\030.nebulaidl.admin.W"
  "orkflow\"I\202\323\344\223\002C\022A/api/v1/workflows/{id.p"
  "roject}/{id.domain}/{id.name}/{id.versio"
  "n}\022\237\001\n\017ListWorkflowIds\0220.nebulaidl.admin."
  "NamedEntityIdentifierListRequest\032).nebula"
  "idl.admin.NamedEntityIdentifierList\"/\202\323\344"
  "\223\002)\022\'/api/v1/workflow_ids/{project}/{dom"
  "ain}\022\276\001\n\rListWorkflows\022#.nebulaidl.admin."
  "ResourceListRequest\032\034.nebulaidl.admin.Wor"
  "kflowList\"j\202\323\344\223\002d\0224/api/v1/workflows/{id"
  ".project}/{id.domain}/{id.name}Z,\022*/api/"
  "v1/workflows/{id.project}/{id.domain}\022\206\001"
  "\n\020CreateLaunchPlan\022\'.nebulaidl.admin.Laun"
  "chPlanCreateRequest\032(.nebulaidl.admin.Lau"
  "nchPlanCreateResponse\"\037\202\323\344\223\002\031\"\024/api/v1/l"
  "aunch_plans:\001*\022\233\001\n\rGetLaunchPlan\022 .nebula"
  "idl.admin.ObjectGetRequest\032\032.nebulaidl.ad"
  "min.LaunchPlan\"L\202\323\344\223\002F\022D/api/v1/launch_p"
  "lans/{id.project}/{id.domain}/{id.name}/"
  "{id.version}\022\242\001\n\023GetActiveLaunchPlan\022\'.f"
  "lyteidl.admin.ActiveLaunchPlanRequest\032\032."
  "nebulaidl.admin.LaunchPlan\"F\202\323\344\223\002@\022>/api/"
  "v1/active_launch_plans/{id.project}/{id."
  "domain}/{id.name}\022\234\001\n\025ListActiveLaunchPl"
  "ans\022+.nebulaidl.admin.ActiveLaunchPlanLis"
  "tRequest\032\036.nebulaidl.admin.LaunchPlanList"
  "\"6\202\323\344\223\0020\022./api/v1/active_launch_plans/{p"
  "roject}/{domain}\022\244\001\n\021ListLaunchPlanIds\0220"
  ".nebulaidl.admin.NamedEntityIdentifierLis"
  "tRequest\032).nebulaidl.admin.NamedEntityIde"
  "ntifierList\"2\202\323\344\223\002,\022*/api/v1/launch_plan"
  "_ids/{project}/{domain}\022\310\001\n\017ListLaunchPl"
  "ans\022#.nebulaidl.admin.ResourceListRequest"
  "\032\036.nebulaidl.admin.LaunchPlanList\"p\202\323\344\223\002j"
  "\0227/api/v1/launch_plans/{id.project}/{id."
  "domain}/{id.name}Z/\022-/api/v1/launch_plan"
  "s/{id.project}/{id.domain}\022\266\001\n\020UpdateLau"
  "nchPlan\022\'.nebulaidl.admin.LaunchPlanUpdat"
  "eRequest\032(.nebulaidl.admin.LaunchPlanUpda"
  "teResponse\"O\202\323\344\223\002I\032D/api/v1/launch_plans"
  "/{id.project}/{id.domain}/{id.name}/{id."
  "version}:\001*\022\201\001\n\017CreateExecution\022&.nebulai"
  "dl.admin.ExecutionCreateRequest\032\'.nebulai"
  "dl.admin.ExecutionCreateResponse\"\035\202\323\344\223\002\027"
  "\"\022/api/v1/executions:\001*\022\216\001\n\021RelaunchExec"
  "ution\022(.nebulaidl.admin.ExecutionRelaunch"
  "Request\032\'.nebulaidl.admin.ExecutionCreate"
  "Response\"&\202\323\344\223\002 \"\033/api/v1/executions/rel"
  "aunch:\001*\022\213\001\n\020RecoverExecution\022\'.nebulaidl"
  ".admin.ExecutionRecoverRequest\032\'.nebulaid"
  "l.admin.ExecutionCreateResponse\"%\202\323\344\223\002\037\""
  "\032/api/v1/executions/recover:\001*\022\225\001\n\014GetEx"
  "ecution\022+.nebulaidl.admin.WorkflowExecuti"
  "onGetRequest\032\031.nebulaidl.admin.Execution\""
  "=\202\323\344\223\0027\0225/api/v1/executions/{id.project}"
  "/{id.domain}/{id.name}\022\244\001\n\017UpdateExecuti"
  "on\022&.nebulaidl.admin.ExecutionUpdateReque"
  "st\032\'.nebulaidl.admin.ExecutionUpdateRespo"
  "nse\"@\202\323\344\223\002:\0325/api/v1/executions/{id.proj"
  "ect}/{id.domain}/{id.name}:\001*\022\271\001\n\020GetExe"
  "cutionData\022/.nebulaidl.admin.WorkflowExec"
  "utionGetDataRequest\0320.nebulaidl.admin.Wor"
  "kflowExecutionGetDataResponse\"B\202\323\344\223\002<\022:/"
  "api/v1/data/executions/{id.project}/{id."
  "domain}/{id.name}\022\211\001\n\016ListExecutions\022#.f"
  "lyteidl.admin.ResourceListRequest\032\035.flyt"
  "eidl.admin.ExecutionList\"3\202\323\344\223\002-\022+/api/v"
  "1/executions/{id.project}/{id.domain}\022\255\001"
  "\n\022TerminateExecution\022).nebulaidl.admin.Ex"
  "ecutionTerminateRequest\032*.nebulaidl.admin"
  ".ExecutionTerminateResponse\"@\202\323\344\223\002:*5/ap"
  "i/v1/executions/{id.project}/{id.domain}"
  "/{id.name}:\001*\022\322\001\n\020GetNodeExecution\022\'.fly"
  "teidl.admin.NodeExecutionGetRequest\032\035.fl"
  "yteidl.admin.NodeExecution\"v\202\323\344\223\002p\022n/api"
  "/v1/node_executions/{id.execution_id.pro"
  "ject}/{id.execution_id.domain}/{id.execu"
  "tion_id.name}/{id.node_id}\022\336\001\n\022ListNodeE"
  "xecutions\022(.nebulaidl.admin.NodeExecution"
  "ListRequest\032!.nebulaidl.admin.NodeExecuti"
  "onList\"{\202\323\344\223\002u\022s/api/v1/node_executions/"
  "{workflow_execution_id.project}/{workflo"
  "w_execution_id.domain}/{workflow_executi"
  "on_id.name}\022\245\004\n\031ListNodeExecutionsForTas"
  "k\022/.nebulaidl.admin.NodeExecutionForTaskL"
  "istRequest\032!.nebulaidl.admin.NodeExecutio"
  "nList\"\263\003\202\323\344\223\002\254\003\022\251\003/api/v1/children/task_"
  "executions/{task_execution_id.node_execu"
  "tion_id.execution_id.project}/{task_exec"
  "ution_id.node_execution_id.execution_id."
  "domain}/{task_execution_id.node_executio"
  "n_id.execution_id.name}/{task_execution_"
  "id.node_execution_id.node_id}/{task_exec"
  "ution_id.task_id.project}/{task_executio"
  "n_id.task_id.domain}/{task_execution_id."
  "task_id.name}/{task_execution_id.task_id"
  ".version}/{task_execution_id.retry_attem"
  "pt}\022\356\001\n\024GetNodeExecutionData\022+.nebulaidl."
  "admin.NodeExecutionGetDataRequest\032,.flyt"
  "eidl.admin.NodeExecutionGetDataResponse\""
  "{\202\323\344\223\002u\022s/api/v1/data/node_executions/{i"
  "d.execution_id.project}/{id.execution_id"
  ".domain}/{id.execution_id.name}/{id.node"
  "_id}\022\177\n\017RegisterProject\022&.nebulaidl.admin"
  ".ProjectRegisterRequest\032\'.nebulaidl.admin"
  ".ProjectRegisterResponse\"\033\202\323\344\223\002\025\"\020/api/v"
  "1/projects:\001*\022q\n\rUpdateProject\022\027.nebulaid"
  "l.admin.Project\032%.nebulaidl.admin.Project"
  "UpdateResponse\" \202\323\344\223\002\032\032\025/api/v1/projects"
  "/{id}:\001*\022f\n\014ListProjects\022\".nebulaidl.admi"
  "n.ProjectListRequest\032\030.nebulaidl.admin.Pr"
  "ojects\"\030\202\323\344\223\002\022\022\020/api/v1/projects\022\231\001\n\023Cre"
  "ateWorkflowEvent\022-.nebulaidl.admin.Workfl"
  "owExecutionEventRequest\032..nebulaidl.admin"
  ".WorkflowExecutionEventResponse\"#\202\323\344\223\002\035\""
  "\030/api/v1/events/workflows:\001*\022\211\001\n\017CreateN"
  "odeEvent\022).nebulaidl.admin.NodeExecutionE"
  "ventRequest\032*.nebulaidl.admin.NodeExecuti"
  "onEventResponse\"\037\202\323\344\223\002\031\"\024/api/v1/events/"
  "nodes:\001*\022\211\001\n\017CreateTaskEvent\022).nebulaidl."
  "admin.TaskExecutionEventRequest\032*.nebulai"
  "dl.admin.TaskExecutionEventResponse\"\037\202\323\344"
  "\223\002\031\"\024/api/v1/events/tasks:\001*\022\200\003\n\020GetTask"
  "Execution\022\'.nebulaidl.admin.TaskExecution"
  "GetRequest\032\035.nebulaidl.admin.TaskExecutio"
  "n\"\243\002\202\323\344\223\002\234\002\022\231\002/api/v1/task_executions/{i"
  "d.node_execution_id.execution_id.project"
  "}/{id.node_execution_id.execution_id.dom"
  "ain}/{id.node_execution_id.execution_id."
  "name}/{id.node_execution_id.node_id}/{id"
  ".task_id.project}/{id.task_id.domain}/{i"
  "d.task_id.name}/{id.task_id.version}/{id"
  ".retry_attempt}\022\230\002\n\022ListTaskExecutions\022("
  ".nebulaidl.admin.TaskExecutionListRequest"
  "\032!.nebulaidl.admin.TaskExecutionList\"\264\001\202\323"
  "\344\223\002\255\001\022\252\001/api/v1/task_executions/{node_ex"
  "ecution_id.execution_id.project}/{node_e"
  "xecution_id.execution_id.domain}/{node_e"
  "xecution_id.execution_id.name}/{node_exe"
  "cution_id.node_id}\022\234\003\n\024GetTaskExecutionD"
  "ata\022+.nebulaidl.admin.TaskExecutionGetDat"
  "aRequest\032,.nebulaidl.admin.TaskExecutionG"
  "etDataResponse\"\250\002\202\323\344\223\002\241\002\022\236\002/api/v1/data/"
  "task_executions/{id.node_execution_id.ex"
  "ecution_id.project}/{id.node_execution_i"
  "d.execution_id.domain}/{id.node_executio"
  "n_id.execution_id.name}/{id.node_executi"
  "on_id.node_id}/{id.task_id.project}/{id."
  "task_id.domain}/{id.task_id.name}/{id.ta"
  "sk_id.version}/{id.retry_attempt}\022\343\001\n\035Up"
  "dateProjectDomainAttributes\0224.nebulaidl.a"
  "dmin.ProjectDomainAttributesUpdateReques"
  "t\0325.nebulaidl.admin.ProjectDomainAttribut"
  "esUpdateResponse\"U\202\323\344\223\002O\032J/api/v1/projec"
  "t_domain_attributes/{attributes.project}"
  "/{attributes.domain}:\001*\022\301\001\n\032GetProjectDo"
  "mainAttributes\0221.nebulaidl.admin.ProjectD"
  "omainAttributesGetRequest\0322.nebulaidl.adm"
  "in.ProjectDomainAttributesGetResponse\"<\202"
  "\323\344\223\0026\0224/api/v1/project_domain_attributes"
  "/{project}/{domain}\022\315\001\n\035DeleteProjectDom"
  "ainAttributes\0224.nebulaidl.admin.ProjectDo"
  "mainAttributesDeleteRequest\0325.nebulaidl.a"
  "dmin.ProjectDomainAttributesDeleteRespon"
  "se\"\?\202\323\344\223\0029*4/api/v1/project_domain_attri"
  "butes/{project}/{domain}:\001*\022\266\001\n\027UpdatePr"
  "ojectAttributes\022..nebulaidl.admin.Project"
  "AttributesUpdateRequest\032/.nebulaidl.admin"
  ".ProjectAttributesUpdateResponse\":\202\323\344\223\0024"
  "\032//api/v1/project_attributes/{attributes"
  ".project}:\001*\022\237\001\n\024GetProjectAttributes\022+."
  "nebulaidl.admin.ProjectAttributesGetReque"
  "st\032,.nebulaidl.admin.ProjectAttributesGet"
  "Response\",\202\323\344\223\002&\022$/api/v1/project_attrib"
  "utes/{project}\022\253\001\n\027DeleteProjectAttribut"
  "es\022..nebulaidl.admin.ProjectAttributesDel"
  "eteRequest\032/.nebulaidl.admin.ProjectAttri"
  "butesDeleteResponse\"/\202\323\344\223\002)*$/api/v1/pro"
  "ject_attributes/{project}:\001*\022\344\001\n\030UpdateW"
  "orkflowAttributes\022/.nebulaidl.admin.Workf"
  "lowAttributesUpdateRequest\0320.nebulaidl.ad"
  "min.WorkflowAttributesUpdateResponse\"e\202\323"
  "\344\223\002_\032Z/api/v1/workflow_attributes/{attri"
  "butes.project}/{attributes.domain}/{attr"
  "ibutes.workflow}:\001*\022\267\001\n\025GetWorkflowAttri"
  "butes\022,.nebulaidl.admin.WorkflowAttribute"
  "sGetRequest\032-.nebulaidl.admin.WorkflowAtt"
  "ributesGetResponse\"A\202\323\344\223\002;\0229/api/v1/work"
  "flow_attributes/{project}/{domain}/{work"
  "flow}\022\303\001\n\030DeleteWorkflowAttributes\022/.fly"
  "teidl.admin.WorkflowAttributesDeleteRequ"
  "est\0320.nebulaidl.admin.WorkflowAttributesD"
  "eleteResponse\"D\202\323\344\223\002>*9/api/v1/workflow_"
  "attributes/{project}/{domain}/{workflow}"
  ":\001*\022\240\001\n\027ListMatchableAttributes\022..nebulai"
  "dl.admin.ListMatchableAttributesRequest\032"
  "/.nebulaidl.admin.ListMatchableAttributes"
  "Response\"$\202\323\344\223\002\036\022\034/api/v1/matchable_attr"
  "ibutes\022\237\001\n\021ListNamedEntities\022&.nebulaidl."
  "admin.NamedEntityListRequest\032\037.nebulaidl."
  "admin.NamedEntityList\"A\202\323\344\223\002;\0229/api/v1/n"
  "amed_entities/{resource_type}/{project}/"
  "{domain}\022\247\001\n\016GetNamedEntity\022%.nebulaidl.a"
  "dmin.NamedEntityGetRequest\032\033.nebulaidl.ad"
  "min.NamedEntity\"Q\202\323\344\223\002K\022I/api/v1/named_e"
  "ntities/{resource_type}/{id.project}/{id"
  ".domain}/{id.name}\022\276\001\n\021UpdateNamedEntity"
  "\022(.nebulaidl.admin.NamedEntityUpdateReque"
  "st\032).nebulaidl.admin.NamedEntityUpdateRes"
  "ponse\"T\202\323\344\223\002N\032I/api/v1/named_entities/{r"
  "esource_type}/{id.project}/{id.domain}/{"
  "id.name}:\001*\022l\n\nGetVersion\022!.nebulaidl.adm"
  "in.GetVersionRequest\032\".nebulaidl.admin.Ge"
  "tVersionResponse\"\027\202\323\344\223\002\021\022\017/api/v1/versio"
  "n\022\304\001\n\024GetDescriptionEntity\022 .nebulaidl.ad"
  "min.ObjectGetRequest\032!.nebulaidl.admin.De"
  "scriptionEntity\"g\202\323\344\223\002a\022_/api/v1/descrip"
  "tion_entities/{id.resource_type}/{id.pro"
  "ject}/{id.domain}/{id.name}/{id.version}"
  "\022\222\002\n\027ListDescriptionEntities\022,.nebulaidl."
  "admin.DescriptionEntityListRequest\032%.fly"
  "teidl.admin.DescriptionEntityList\"\241\001\202\323\344\223"
  "\002\232\001\022O/api/v1/description_entities/{resou"
  "rce_type}/{id.project}/{id.domain}/{id.n"
  "ame}ZG\022E/api/v1/description_entities/{re"
  "source_type}/{id.project}/{id.domain}\022\305\001"
  "\n\023GetExecutionMetrics\0222.nebulaidl.admin.W"
  "orkflowExecutionGetMetricsRequest\0323.flyt"
  "eidl.admin.WorkflowExecutionGetMetricsRe"
  "sponse\"E\202\323\344\223\002\?\022=/api/v1/metrics/executio"
  "ns/{id.project}/{id.domain}/{id.name}B9Z"
  "7github.com/nebulaclouds/nebulaidl/gen/pb-go/"
  "nebulaidl/serviceb\006proto3"
  ;
::google::protobuf::internal::DescriptorTable descriptor_table_nebulaidl_2fservice_2fadmin_2eproto = {
  false, InitDefaults_nebulaidl_2fservice_2fadmin_2eproto,
  descriptor_table_protodef_nebulaidl_2fservice_2fadmin_2eproto,
  "nebulaidl/service/admin.proto", &assign_descriptors_table_nebulaidl_2fservice_2fadmin_2eproto, 10664,
};

void AddDescriptors_nebulaidl_2fservice_2fadmin_2eproto() {
  static constexpr ::google::protobuf::internal::InitFunc deps[16] =
  {
    ::AddDescriptors_google_2fapi_2fannotations_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fproject_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fproject_5fdomain_5fattributes_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fproject_5fattributes_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2ftask_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fworkflow_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fworkflow_5fattributes_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2flaunch_5fplan_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fevent_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fexecution_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fmatchable_5fresource_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fnode_5fexecution_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2ftask_5fexecution_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fversion_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fcommon_2eproto,
    ::AddDescriptors_nebulaidl_2fadmin_2fdescription_5fentity_2eproto,
  };
 ::google::protobuf::internal::AddDescriptors(&descriptor_table_nebulaidl_2fservice_2fadmin_2eproto, deps, 16);
}

// Force running AddDescriptors() at dynamic initialization time.
static bool dynamic_init_dummy_nebulaidl_2fservice_2fadmin_2eproto = []() { AddDescriptors_nebulaidl_2fservice_2fadmin_2eproto(); return true; }();
namespace nebulaidl {
namespace service {

// @@protoc_insertion_point(namespace_scope)
}  // namespace service
}  // namespace nebulaidl
namespace google {
namespace protobuf {
}  // namespace protobuf
}  // namespace google

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
