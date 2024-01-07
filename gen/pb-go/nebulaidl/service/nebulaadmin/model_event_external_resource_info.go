/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// This message contains metadata about external resources produced or used by a specific task execution.
type EventExternalResourceInfo struct {
	// Identifier for an external resource created by this task execution, for example Qubole query ID or presto query ids.
	ExternalId string `json:"external_id,omitempty"`
	// A unique index for the external resource with respect to all external resources for this task. Although the identifier may change between task reporting events or retries, this will remain the same to enable aggregating information from multiple reports.
	Index int64 `json:"index,omitempty"`
	RetryAttempt int64 `json:"retry_attempt,omitempty"`
	Phase *CoreTaskExecutionPhase `json:"phase,omitempty"`
	// Captures the status of caching for this external resource execution.
	CacheStatus *CoreCatalogCacheStatus `json:"cache_status,omitempty"`
	Logs []CoreTaskLog `json:"logs,omitempty"`
}
