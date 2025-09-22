// Package helper provides utility functions for loading and managing environment variables
// required for Appwrite client configuration.
package helper

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Global variables that store Appwrite configuration loaded from environment variables.
// These are populated by calling Envvars() and should not be modified directly.
var (
    // AppwriteEndpointURL is the Appwrite server endpoint URL
    AppwriteEndpointURL string
    
    // AppwriteProjectID is the Appwrite project identifier
    AppwriteProjectID  string
    
    // AppwriteRESDEFAPIKey is the API key used for resource definition operations
    AppwriteRESDEFAPIKey    string
)

// Envvars loads environment variables from the .env.local file and populates the global
// configuration variables required for Appwrite client initialisation.
//
// This function reads the following environment variables:
//   - APPWRITE_ENDPOINT_URL: The Appwrite server endpoint URL
//   - APPWRITE_PROJECT_ID: The Appwrite project ID  
//   - APPWRITE_API_KEY_APPRES: The API key with appropriate permissions
//
// The function will terminate the program with log.Fatalf if the .env.local file
// cannot be loaded. Ensure the file exists in the current working directory
// and contains all required environment variables.
//
// This function is typically called automatically by appres.Utils() and should
// not need to be called directly by users of the appres package.
func Envvars() {
    // Load .env.local file
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatalf("Error loading .env.local file")
    }

    // Reference variables
    AppwriteEndpointURL = os.Getenv("APPWRITE_ENDPOINT_URL")
    AppwriteProjectID = os.Getenv("APPWRITE_PROJECT_ID")
    AppwriteRESDEFAPIKey = os.Getenv("APPWRITE_API_KEY_APPRES")
}
