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
	Default string
	
	// Array indicates whether the attribute can store multiple values as an array
	Array bool
	
	// Encrypt determines whether the attribute value should be encrypted at rest
	// Note: Only available for string attributes
	Encrypt bool
	
	// Min is the minimum value for integer attributes (optional)
	// If not set (0), no minimum constraint will be applied
	Min float64
	
	// Max is the maximum value for integer attributes (optional)
	// If not set (0), no maximum constraint will be applied
	Max float64
}
