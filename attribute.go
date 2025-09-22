package appres

import (
	"fmt"
	"log"
	"time"

	"github.com/appwrite/sdk-for-go/databases"
)

// CreateAttribute creates a new attribute in the specified collection or skips creation if it already exists.
// It checks for duplicates to avoid errors and supports all major attribute types.
//
// Parameters:
//   - dbID: The ID of the database containing the collection
//   - colID: The ID of the collection where the attribute should be created
//   - att: AttributeType struct containing the attribute configuration
//
// Global Variables Used:
//   - AppwriteDatabase: The initialized Appwrite database client
//
// Returns:
//   - error: Any error that occurred during the operation, or nil if successful
//
// Supported types: string, email, integer, datetime, boolean, relationship, url
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
func CreateAttribute(dbID string, colID string, att AttributeType) error {
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
			opts...,
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
		if att.Min != nil {
			opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeMin(att.Min.(int)))
		}
		if att.Max != nil {
			opts = append(opts, AppwriteDatabase.WithCreateIntegerAttributeMax(att.Max.(int)))
		}
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
		//----------------------------------------------------------------------------------------
		// Create RELATIONSHIP attribute
		//----------------------------------------------------------------------------------------
	} else if att.Type == "relationship" {
		var opts []databases.CreateRelationshipAttributeOption
		opts = append(opts, AppwriteDatabase.WithCreateRelationshipAttributeTwoWay(att.TwoWay))
		if att.Name != "" {
			opts = append(opts, AppwriteDatabase.WithCreateRelationshipAttributeKey(att.Name))
		}
		if att.OnDelete != "" {
			opts = append(opts, AppwriteDatabase.WithCreateRelationshipAttributeOnDelete(att.OnDelete))
		}
		if att.TwoWayKey != "" {
			opts = append(opts, AppwriteDatabase.WithCreateRelationshipAttributeTwoWayKey(att.TwoWayKey))
		}
		newAtt, err := AppwriteDatabase.CreateRelationshipAttribute(
			dbID,
			colID,
			att.RelatedCollectionID,
			att.RelationshipType,
			opts...,
		)
		if err != nil {
			log.Println("error creating attribute:", err)
			return err
		}
		log.Println("attribute created with key:", newAtt.Key)
		return nil
		//----------------------------------------------------------------------------------------
		// Create URL attribute
		//----------------------------------------------------------------------------------------
	} else if att.Type == "url" {
		var opts []databases.CreateUrlAttributeOption
		if att.Default != nil {
			if _, ok := att.Default.(string); !ok {
				return fmt.Errorf("default value for url attribute must be a string")
			}
			opts = append(opts, AppwriteDatabase.WithCreateUrlAttributeDefault(att.Default.(string)))
		}
		opts = append(opts, AppwriteDatabase.WithCreateUrlAttributeArray(att.Array))
		newAtt, err := AppwriteDatabase.CreateUrlAttribute(
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