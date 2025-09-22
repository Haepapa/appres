package appres

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/databases"
	"github.com/appwrite/sdk-for-go/storage"

	"github.com/Haepapa/appres/helper"
)

// AppwriteDatabase is the global database client instance used by all database operations.
// It is initialised by calling Utils() and should not be accessed directly.
var (
	AppwriteDatabase *databases.Databases
	AppwriteStorage   *storage.Storage
)

// Utils initialises the Appwrite client with configuration from environment variables.
// It loads environment variables from the .env.local file and creates a new Appwrite client
// with the configured endpoint, project ID, and API key.
//
// This function must be called before using any other functions in this package.
// It will terminate the program if the .env.local file cannot be loaded.
//
// Environment variables required:
//   - APPWRITE_ENDPOINT_URL: The Appwrite server endpoint URL
//   - APPWRITE_PROJECT_ID: The Appwrite project ID  
//   - APPWRITE_API_KEY_APPRES: The API key with database and storage permissions
//
// Example:
//
//	app.Utils()
//	// Now you can use other functions such as CreateDatabase, CreateCollection, etc.
func Utils() {
	helper.Envvars()
	client := appwrite.NewClient(
		appwrite.WithEndpoint(helper.AppwriteEndpointURL),
		appwrite.WithProject(helper.AppwriteProjectID),
		appwrite.WithKey(helper.AppwriteRESDEFAPIKey),
	)
	AppwriteDatabase = appwrite.NewDatabases(client)
	AppwriteStorage = appwrite.NewStorage(client)
}