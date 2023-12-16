/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// AdminMatchableResource : Defines a resource that can be configured by customizable Project-, ProjectDomain- or WorkflowAttributes based on matching tags.   - TASK_RESOURCE: Applies to customizable task resource requests and limits.  - CLUSTER_RESOURCE: Applies to configuring templated kubernetes cluster resources.  - EXECUTION_QUEUE: Configures task and dynamic task execution queue assignment.  - EXECUTION_CLUSTER_LABEL: Configures the K8s cluster label to be used for execution to be run  - QUALITY_OF_SERVICE_SPECIFICATION: Configures default quality of service when undefined in an execution spec.  - PLUGIN_OVERRIDE: Selects configurable plugin implementation behavior for a given task type.  - WORKFLOW_EXECUTION_CONFIG: Adds defaults for customizable workflow-execution specifications and overrides.  - CLUSTER_ASSIGNMENT: Controls how to select an available cluster on which this execution should run.
type AdminMatchableResource string

// List of adminMatchableResource
const (
	AdminMatchableResourceTASK_RESOURCE                    AdminMatchableResource = "TASK_RESOURCE"
	AdminMatchableResourceCLUSTER_RESOURCE                 AdminMatchableResource = "CLUSTER_RESOURCE"
	AdminMatchableResourceEXECUTION_QUEUE                  AdminMatchableResource = "EXECUTION_QUEUE"
	AdminMatchableResourceEXECUTION_CLUSTER_LABEL          AdminMatchableResource = "EXECUTION_CLUSTER_LABEL"
	AdminMatchableResourceQUALITY_OF_SERVICE_SPECIFICATION AdminMatchableResource = "QUALITY_OF_SERVICE_SPECIFICATION"
	AdminMatchableResourcePLUGIN_OVERRIDE                  AdminMatchableResource = "PLUGIN_OVERRIDE"
	AdminMatchableResourceWORKFLOW_EXECUTION_CONFIG        AdminMatchableResource = "WORKFLOW_EXECUTION_CONFIG"
	AdminMatchableResourceCLUSTER_ASSIGNMENT               AdminMatchableResource = "CLUSTER_ASSIGNMENT"
)
