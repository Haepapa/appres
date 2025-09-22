package appres

import (
	"log"

	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
)

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