// Package appres provides utilities for creating and managing Appwrite resources programmatically.
// It simplifies the process of creating databases, collections, and attributes in your Appwrite backend.
//
// Usage:
//
//	app.Utils()
//	db, err := app.CreateDatabase("my-database")
//	if err != nil {
//		log.Fatal(err)
//	}
//	col, err := app.CreateCollection(db.Id, "my-collection")
//	if err != nil {
//		log.Fatal(err)
//	}
package appres

import (
	"fmt"
	"log"
	"time"

	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"

	"github.com/Haepapa/appres/helper"
)

// AppwriteDatabase is the global database client instance used by all database operations.
// It is initialised by calling Utils() and should not be accessed directly.
var (
	AppwriteDatabase *databases.Databases
)

// Utils initialises the Appwrite client with configuration from environment variables.
// It loads environment variables from the .env.local file and creates a new Appwrite client
// with the configured endpoint, project ID, and API key.
//
// This function must be called before using any other functions in this package.
// It will terminate the program if the .env.local file cannot be loaded.
//
// Environment variables required:
//   - NEXT_PUBLIC_APPWRITE_ENDPOINT: The Appwrite server endpoint URL
//   - NEXT_PUBLIC_APPWRITE_PROJECT: The Appwrite project ID
//   - APPWRITE_API_KEY_RESDEF: The API key with appropriate permissions
//
// Example:
//
//	app.Utils()
//	// Now you can use other functions such as CreateDatabase, CreateCollection, etc.
func Utils(){
	helper.Envvars()
	client := appwrite.NewClient(
		appwrite.WithEndpoint(helper.AppwriteEndpointURL),
		appwrite.WithProject(helper.AppwriteProjectID),
		appwrite.WithKey(helper.AppwriteRESDEFAPIKey),
	)
	AppwriteDatabase = appwrite.NewDatabases(client)
}

// CreateDatabase creates a new database with the specified name or returns the existing one if it already exists.
// It first checks if a database with the given name already exists to avoid duplicates.
//
// The function automatically generates a unique ID for new databases and logs the creation process.
//
// Parameters:
//   - name: The name of the database to create
//
// Returns:
//   - *models.Database: Pointer to the created or existing database
//   - error: Any error that occurred during the operation
//
// Example:
//
//	db, err := app.CreateDatabase("my-app-database")
//	if err != nil {
//		log.Fatal("Failed to create database:", err)
//	}
//	fmt.Printf("Database created with ID: %s\n", db.Id)
func CreateDatabase(name string) (*models.Database, error) {
	// List all databases
	databases, err := AppwriteDatabase.List()
	if err != nil {
		log.Println("Error listing databases:", err)
		return nil, err
	}
	for _, db := range databases.Databases {
		if db.Name == name {
			log.Println("Database already exists with id:", db.Id)
			return &db, nil
		}
	}
	// Create a database
	db, err := AppwriteDatabase.Create(id.Unique(), name)
	if err != nil {
		log.Println("Error creating database:", err)
		return nil, err
	}
	log.Println("Database created with id:", db.Id)
	return db, nil
}
// CreateCollection creates a new collection in the specified database or returns the existing one if it already exists.
// It first checks if a collection with the given name already exists in the database to avoid duplicates.
//
// The function automatically generates a unique ID for new collections and logs the creation process.
//
// Parameters:
//   - dbId: The ID of the database where the collection should be created
//   - name: The name of the collection to create
//
// Returns:
//   - *models.Collection: Pointer to the created or existing collection
//   - error: Any error that occurred during the operation
//
// Example:
//
//	col, err := app.CreateCollection(db.Id, "users")
//	if err != nil {
//		log.Fatal("Failed to create collection:", err)
//	}
//	fmt.Printf("Collection created with ID: %s\n", col.Id)
func CreateCollection(dbId string, name string) (*models.Collection, error) {
	// List all collections in database
	collections, err := AppwriteDatabase.ListCollections(dbId)
	if err != nil {
		log.Println("Error listing collections:", err)
		return nil, err
	}
	for _, col := range collections.Collections {
		if col.Name == name {
			log.Println("Collection already exists with id:", col.Id)
			return &col, nil
		}
	}
	// Create a collection
	col, err := AppwriteDatabase.CreateCollection(dbId, id.Unique(), name)
	if err != nil {
		log.Println("Error creating collection:", err)
		return nil, err
	}
	log.Println("Collection created with id:", col.Id)
	return col, nil
}

// CreateAttribute creates a new attribute in the specified collection or skips creation if it already exists.
// It first checks if an attribute with the given name already exists in the collection to avoid duplicates.
//
// The function supports creating string and email attributes with full configuration options including
// size limits, default values, array types, and encryption settings.
//
// Parameters:
//   - dbID: The ID of the database containing the collection
//   - colID: The ID of the collection where the attribute should be created
//   - att: AttributeType struct containing the attribute configuration
//
// Returns:
//   - error: Any error that occurred during the operation, or nil if successful
//
// Supported attribute types:
//   - "string": Text attributes with configurable size, defaults, arrays, and encryption
//   - "email": Email validation attributes with defaults and array support
//
// References:
//   - Appwrite Documentation: https://appwrite.io/docs/references/cloud/server-go/databases
//
// Example:
//
//	attr := app.AttributeType{
//		Type:     "string",
//		Name:     "title",
//		Size:     255,
//		Required: true,
//		Default:  "",
//		Array:    false,
//		Encrypt:  false,
//	}
//	err := app.CreateAttribute(db.Id, col.Id, attr)
//	if err != nil {
//		log.Fatal("Failed to create attribute:", err)
//	}
func CreateAttribute(dbID string, colID string, att AttributeType) (error){
	attributes, err := AppwriteDatabase.ListAttributes(dbID, colID)
	if err != nil {
		log.Println("Error listing attributes:", err)
		return err
	}
	for _, attr := range attributes.Attributes {
		if attrName, ok := attr["key"].(string); ok && attrName == att.Name {
			log.Println("Attribute already exists with key:", attr["key"])
			return nil
		}
	}
	//----------------------------------------------------------------------------------------
	// Create STRING attribute
	//----------------------------------------------------------------------------------------
	if att.Type == "string" {
		if att.Default != nil {
			if _, ok := att.Default.(string); !ok {
				return fmt.Errorf("default value for string attribute must be a string")
			}
		}
		var opts []databases.CreateStringAttributeOption
		if !att.Required && att.Default != nil {
			opts = append(opts, AppwriteDatabase.WithCreateStringAttributeDefault(att.Default.(string)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateStringAttributeArray(att.Array))
		opts = append(opts, AppwriteDatabase.WithCreateStringAttributeEncrypt(att.Encrypt))

		newAtt, err := AppwriteDatabase.CreateStringAttribute(
			dbID, 
			colID, 
			att.Name, 
			att.Size, 
			att.Required,
			opts...,
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
	//----------------------------------------------------------------------------------------
	// Create EMAIL attribute
	//----------------------------------------------------------------------------------------
	} else if att.Type == "email" {
		if att.Default != nil {
			if _, ok := att.Default.(string); !ok {
				return fmt.Errorf("default value for string attribute must be a string")
			}
		}
		var opts []databases.CreateEmailAttributeOption
		if !att.Required && att.Default != nil {
			opts = append(opts, AppwriteDatabase.WithCreateEmailAttributeDefault(att.Default.(string)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateEmailAttributeArray(att.Array))
		newAtt, err := AppwriteDatabase.CreateEmailAttribute(
			dbID,
			colID,
			att.Name,
			att.Required,
			opts...
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
	//----------------------------------------------------------------------------------------
	// Create INTEGER attribute
	//----------------------------------------------------------------------------------------
	} else if att.Type == "integer" {
		if att.Default != nil {
			_, ok := att.Default.(int)
			if !ok {
				return fmt.Errorf("default value for integer attribute must be an int")
			}
		}
		if att.Min != nil {
			_, ok := att.Min.(int)
			if !ok {
				return fmt.Errorf("min value for integer attribute must be an int")
			}
		}
		if att.Max != nil {
			_, ok := att.Max.(int)
			if !ok {
				return fmt.Errorf("max value for integer attribute must be an int")
			}
		}
		var opts []databases.CreateIntegerAttributeOption
		if !att.Required && att.Default != nil {
			opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeDefault(att.Default.(int)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeMin(att.Min.(int)))
		opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeMax(att.Max.(int)))
		opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeArray(att.Array))
		newAtt, err := AppwriteDatabase.CreateIntegerAttribute(
			dbID,
			colID,
			att.Name,
			att.Required,
			opts...,
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
	//----------------------------------------------------------------------------------------
	// Create DATETIME attribute
	//----------------------------------------------------------------------------------------
	} else if att.Type == "datetime" {
		layout := time.RFC3339
		if att.Default != nil {
			s, ok := att.Default.(string)
			if !ok {
				return fmt.Errorf("default value for datetime attribute must be a string")
			}
			if _, err := time.Parse(layout, s); err != nil {
				return fmt.Errorf("default value for datetime attribute must be a valid RFC3339 datetime string: %v", err)
			}
		}
		var opts []databases.CreateDatetimeAttributeOption
		if !att.Required && att.Default != nil {
			opts = append(opts, AppwriteDatabase.WithCreateDatetimeAttributeDefault(att.Default.(string)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateDatetimeAttributeArray(att.Array))
		newAtt, err := AppwriteDatabase.CreateDatetimeAttribute(
			dbID,
			colID,
			att.Name,
			att.Required,
			opts...,
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
	//----------------------------------------------------------------------------------------
	// Create BOOLEAN attribute
	//----------------------------------------------------------------------------------------
	} else if att.Type == "boolean" {
		if att.Default != nil {
			if _, ok := att.Default.(bool); !ok {
				return fmt.Errorf("default value for boolean attribute must be a bool")
			}
		}
		var opts []databases.CreateBooleanAttributeOption
		if !att.Required && att.Default != nil {
			opts = append(opts, AppwriteDatabase.WithCreateBooleanAttributeDefault(att.Default.(bool)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateBooleanAttributeArray(att.Array))
		newAtt, err := AppwriteDatabase.CreateBooleanAttribute(
			dbID,
			colID,
			att.Name,
			att.Required,
			opts...,
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
	}
	return fmt.Errorf("unsupported attribute type: %s", att.Type)
}
