/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Additional metadata around a named entity.
type AdminNamedEntityMetadata struct {
	Description string `json:"description,omitempty"`
	// Shared state across all version of the entity At this point in time, only workflow entities can have their state archived.
	State *AdminNamedEntityState `json:"state,omitempty"`
}
