/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// A generic key value pair.
type CoreKeyValuePair struct {
	// required.
	Key string `json:"key,omitempty"`
	// +optional.
	Value string `json:"value,omitempty"`
}
