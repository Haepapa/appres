package appres

type AttributeType struct {
	Type string
	Name string
	Size int
	Required bool
	// Unique bool - not implemented in sdk (https://github.com/appwrite/sdk-for-go/blob/main/databases/databases.go)
	Default string
	Array bool
	Encrypt bool
}