/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Indicates the priority of an execution.
type CoreQualityOfService struct {
	Tier *QualityOfServiceTier     `json:"tier,omitempty"`
	Spec *CoreQualityOfServiceSpec `json:"spec,omitempty"`
}
