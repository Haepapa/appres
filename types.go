package appres

// AttributeType defines the configuration for creating attributes in Appwrite collections.
// It contains all the necessary fields to specify the type, constraints, and behavior
// of an attribute when creating it in a collection.
//
// Supported attribute types:
//   - "string": Text attributes with size, encryption, and array support
//   - "email": Email validation attributes with array support  
//   - "integer": Integer attributes with min/max constraints and array support
//   - "datetime": Date and time attributes with array support
//   - "boolean": Boolean (true/false) attributes with array support
//   - "relationship": Relationship attributes linking collections
//   - "url": URL validation attributes with array support
//
// Example usage:
//
//	attr := AttributeType{
//		Type:     "string",
//		Name:     "username",
//		Size:     50,
//		Required: true,
//		Default:  "",
//		Array:    false,
//		Encrypt:  false,
//	}
//
//	// Integer attribute example:
//	intAttr := AttributeType{
//		Type:     "integer",
//		Name:     "age",
//		Required: true,
//		Min:      0,
//		Max:      120,
//		Default:  "18",
//		Array:    false,
//	}
type AttributeType struct {
	// Type specifies the attribute type. Supported values: "string", "email", "integer", "datetime", "boolean"
	Type string
	
	// Name is the key/identifier for the attribute in the collection
	Name string
	
	// Size defines the maximum length for string and email attributes
	Size int
	
	// Required determines whether this attribute must have a value
	Required bool
	
	// Unique bool - not implemented in sdk (https://github.com/appwrite/sdk-for-go/blob/main/databases/databases.go)
	
	// Default is the default value assigned to the attribute if no value is provided
	Default interface{}
	
	// Array indicates whether the attribute can store multiple values as an array
	Array bool
	
	// Encrypt determines whether the attribute value should be encrypted at rest
	// Note: Only available for string attributes
	Encrypt bool
	
	// Min is the minimum value for integer attributes (optional)
	// If not set (0), no minimum constraint will be applied
	Min interface{}
	
	// Max is the maximum value for integer attributes (optional)
	// If not set (0), no maximum constraint will be applied
	Max interface{}

	// The ID of the collection this relationship attribute links to.
	RelatedCollectionID string

	// The type of relationship
	// must be one of; `oneToOne`, `oneToMany`, `manyToOne`, `manyToMany`.
	// Reference documentation: https://appwrite.io/docs/products/databases/relationships#types
	RelationshipType string

	// Enable two-way directionality
	// false: One-way - The relationship is only visible to one side of the relation. This is similar to a tree data structure.
	// true:  Two-way - The relationship is visible to both sides of the relationship. This is similar to a graph data structure.
	// Reference documentation: https://appwrite.io/docs/products/databases/relationships#directionality
	TwoWay bool

	// The key/identifier used to name the two-way relationship on the related collection side.
	TwoWayKey string

	// On delete constraint behaviour for relationship attributes
	// must be one of; `restrict`, `cascade`, `setnull`.
	// Restrict: If a row has at least one related row, it cannot be deleted.
	// Cascade:	If a row has related rows, when it is deleted, the related rows are also deleted.
	// Set null: If a row has related rows, when it is deleted, the related rows are kept with their relationship column set to null.
	// Reference documentation: https://appwrite.io/docs/products/databases/relationships#on-delete
	OnDelete string
}

// BucketType defines the configuration for creating storage buckets in Appwrite.
// It contains all the necessary fields to specify bucket behavior, security, and constraints.
//
// Example usage:
//
//	bucket := BucketType{
//		Name:         "user-uploads",
//		Enabled:      true,
//		FileSecurity: true,
//		MaxFileSize:  10000000, // 10MB
//		Permissions:  []string{"read(\"any\")"},
//		Compression:  "gzip",
//		Encryption:   true,
//		Antivirus:    true,
//	}
type BucketType struct {
	// Name is the bucket identifier
	Name string

	// Permissions is an array of permission strings (e.g. "read(\"any\")")
	Permissions []string

	// FileSecurity enables file-level security permissions
	FileSecurity bool

	// Enabled determines if the bucket is accessible to users
	Enabled bool

	// MaxFileSize is the maximum file size allowed in bytes (max: 30MB)
	MaxFileSize int

	// AllowedFileExtensions limits file types (max: 100 extensions)
	AllowedFileExtensions []string

	// Compression algorithm: "none", "gzip", or "zstd"
	Compression string

	// Encryption enables file encryption at rest
	Encryption bool

	// Antivirus enables virus scanning for uploaded files  
	Antivirus bool
}
