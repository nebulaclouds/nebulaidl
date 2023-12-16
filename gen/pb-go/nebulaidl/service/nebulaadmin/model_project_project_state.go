/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// ProjectProjectState : The state of the project is used to control its visibility in the UI and validity.   - ACTIVE: By default, all projects are considered active.  - ARCHIVED: Archived projects are no longer visible in the UI and no longer valid.  - SYSTEM_GENERATED: System generated projects that aren't explicitly created or managed by a user.
type ProjectProjectState string

// List of ProjectProjectState
const (
	ProjectProjectStateACTIVE           ProjectProjectState = "ACTIVE"
	ProjectProjectStateARCHIVED         ProjectProjectState = "ARCHIVED"
	ProjectProjectStateSYSTEM_GENERATED ProjectProjectState = "SYSTEM_GENERATED"
)
