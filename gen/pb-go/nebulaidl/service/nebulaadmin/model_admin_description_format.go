/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// AdminDescriptionFormat : - DESCRIPTION_FORMAT_RST: python default documentation - comments is rst
type AdminDescriptionFormat string

// List of adminDescriptionFormat
const (
	AdminDescriptionFormatUNKNOWN  AdminDescriptionFormat = "DESCRIPTION_FORMAT_UNKNOWN"
	AdminDescriptionFormatMARKDOWN AdminDescriptionFormat = "DESCRIPTION_FORMAT_MARKDOWN"
	AdminDescriptionFormatHTML     AdminDescriptionFormat = "DESCRIPTION_FORMAT_HTML"
	AdminDescriptionFormatRST      AdminDescriptionFormat = "DESCRIPTION_FORMAT_RST"
)
