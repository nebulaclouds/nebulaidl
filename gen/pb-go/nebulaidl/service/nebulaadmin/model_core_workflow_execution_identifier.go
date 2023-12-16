/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

type CoreWorkflowExecutionIdentifier struct {
	// Name of the project the resource belongs to.
	Project string `json:"project,omitempty"`
	// Name of the domain the resource belongs to. A domain can be considered as a subset within a specific project.
	Domain string `json:"domain,omitempty"`
	// User or system provided value for the resource.
	Name string `json:"name,omitempty"`
}
