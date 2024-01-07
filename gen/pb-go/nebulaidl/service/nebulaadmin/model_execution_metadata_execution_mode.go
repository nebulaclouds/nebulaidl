/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin
// ExecutionMetadataExecutionMode : The method by which this execution was launched.   - MANUAL: The default execution mode, MANUAL implies that an execution was launched by an individual.  - SCHEDULED: A schedule triggered this execution launch.  - SYSTEM: A system process was responsible for launching this execution rather an individual.  - RELAUNCH: This execution was launched with identical inputs as a previous execution.  - CHILD_WORKFLOW: This execution was triggered by another execution.  - RECOVERED: This execution was recovered from another execution.
type ExecutionMetadataExecutionMode string

// List of ExecutionMetadataExecutionMode
const (
	ExecutionMetadataExecutionModeMANUAL ExecutionMetadataExecutionMode = "MANUAL"
	ExecutionMetadataExecutionModeSCHEDULED ExecutionMetadataExecutionMode = "SCHEDULED"
	ExecutionMetadataExecutionModeSYSTEM ExecutionMetadataExecutionMode = "SYSTEM"
	ExecutionMetadataExecutionModeRELAUNCH ExecutionMetadataExecutionMode = "RELAUNCH"
	ExecutionMetadataExecutionModeCHILD_WORKFLOW ExecutionMetadataExecutionMode = "CHILD_WORKFLOW"
	ExecutionMetadataExecutionModeRECOVERED ExecutionMetadataExecutionMode = "RECOVERED"
)
