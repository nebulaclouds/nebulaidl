/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// CoreSimpleType : Define a set of simple types.
type CoreSimpleType string

// List of coreSimpleType
const (
	CoreSimpleTypeNONE     CoreSimpleType = "NONE"
	CoreSimpleTypeINTEGER  CoreSimpleType = "INTEGER"
	CoreSimpleTypeFLOAT    CoreSimpleType = "FLOAT"
	CoreSimpleTypeSTRING_  CoreSimpleType = "STRING"
	CoreSimpleTypeBOOLEAN  CoreSimpleType = "BOOLEAN"
	CoreSimpleTypeDATETIME CoreSimpleType = "DATETIME"
	CoreSimpleTypeDURATION CoreSimpleType = "DURATION"
	CoreSimpleTypeBINARY   CoreSimpleType = "BINARY"
	CoreSimpleTypeERROR_   CoreSimpleType = "ERROR"
	CoreSimpleTypeSTRUCT_  CoreSimpleType = "STRUCT"
)
