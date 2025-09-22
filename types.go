package appres

// AttributeType defines the configuration for creating attributes in Appwrite collections.
// It contains all the necessary fields to specify the type, constraints, and behavior
// of an attribute when creating it in a collection.
//
// Supported attribute types:
//   - "string": Text attributes with configurable size limits
//   - "email": Email validation attributes
//   - "integer": Integer attributes with configurable min/max constraints
//   - "datetime": Date and time attributes
//   - "boolean": Boolean (true/false) attributes
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

	RelatedCollectionID string

	RelationshipType string

	TwoWay bool

	TwoWayKey string

	OnDelete string
}

type BucketType struct {
	// bucket name
	Name        string

	// An array of permission strings.
	// e.g. read("any") grant read access to role "any"
	Permissions []string

	// When file security is enabled, users will be able to access files for which they have been granted either File or Bucket permissions.
	// If file security is disabled, users can access files only if they have Bucket permissions. 
	FileSecurity bool

	// Is bucket enabled? When set to 'disabled', users cannot access the files in this bucket but Server SDKs with and API key can still access the bucket. 
	// No files are lost when this is toggled.
	Enabled     bool

	// Maximum file size allowed in bytes. Maximum allowed value is 30MB.
	MaxFileSize int

	// Allowed file extensions. Maximum of 100 extensions are allowed, each 64 characters long.
	AllowedFileExtensions []string

	// Compression algorithm choosen for compression. 
	// Can be one of:
	// - none
	// - gzip
	// - zstd
	// For file size above 20MB compression is skipped even if it's enabled
	Compression string

	// Is encryption enabled? For file size above 20MB encryption is skipped even if it's enabled
	Encryption  bool

	// Is virus scanning enabled? For file size above 20MB AntiVirus scanning is skipped even if it's enabled
	Antivirus   bool
}