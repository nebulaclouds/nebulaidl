/*
 * nebulaidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nebulaadmin

// Defines a strong type to allow type checking between interfaces.
type CoreLiteralType struct {
	// A simple type that can be compared one-to-one with another.
	Simple *CoreSimpleType `json:"simple,omitempty"`
	// A complex type that requires matching of inner fields.
	Schema *CoreSchemaType `json:"schema,omitempty"`
	// Defines the type of the value of a collection. Only homogeneous collections are allowed.
	CollectionType *CoreLiteralType `json:"collection_type,omitempty"`
	// Defines the type of the value of a map type. The type of the key is always a string.
	MapValueType *CoreLiteralType `json:"map_value_type,omitempty"`
	// A blob might have specialized implementation details depending on associated metadata.
	Blob *CoreBlobType `json:"blob,omitempty"`
	// Defines an enum with pre-defined string values.
	EnumType              *CoreEnumType              `json:"enum_type,omitempty"`
	StructuredDatasetType *CoreStructuredDatasetType `json:"structured_dataset_type,omitempty"`
	// Defines an union type with pre-defined LiteralTypes.
	UnionType *CoreUnionType `json:"union_type,omitempty"`
	// This field contains type metadata that is descriptive of the type, but is NOT considered in type-checking.  This might be used by consumers to identify special behavior or display extended information for the type.
	Metadata *ProtobufStruct `json:"metadata,omitempty"`
	// This field contains arbitrary data that might have special semantic meaning for the client but does not effect internal nebula behavior.
	Annotation *CoreTypeAnnotation `json:"annotation,omitempty"`
	// Hints to improve type matching.
	Structure *CoreTypeStructure `json:"structure,omitempty"`
}
