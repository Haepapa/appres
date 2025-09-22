package appres

import (
	"log"

	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
)

// CreateDatabase creates a new database with the specified name or returns the existing one if it already exists.
// It first checks if a database with the given name already exists to avoid duplicates.
//
// The function automatically generates a unique ID for new databases and logs the creation process.
//
// Parameters:
//   - name: The name of the database to create
//
// Global Variables Used:
//   - AppwriteDatabase: The initialized Appwrite database client
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