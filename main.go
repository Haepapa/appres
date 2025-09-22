// Package appres provides utilities for creating and managing Appwrite resources programmatically.
// It simplifies the process of creating various resources in your Appwrite backend.
//
// This package offers a simplified interface to the Appwrite Go SDK, providing functions to:
//   - Initialize the Appwrite client
//   - Create databases with duplicate checking
//   - Create collections within databases
//   - Create various types of attributes (string, email, integer, datetime, boolean, relationship, url)
//   - Create storage buckets with security and file constraints
//
// All functions include built-in duplicate checking to prevent errors when resources already exist.
//
// Basic Usage:
//
//	package main
//
//	import (
//		"log"
//		"github.com/Haepapa/appres"
//	)
//
//	func main() {
//		// Initialize the Appwrite client (required first step)
//		appres.Utils()
//
//		// Create a database
//		db, err := appres.CreateDatabase("my-database")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Create a collection
//		col, err := appres.CreateCollection(db.Id, "my-collection")
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// Create attributes
//		attr := appres.AttributeType{
//			Type:     "string",
//			Name:     "title",
//			Size:     255,
//			Required: true,
//		}
//		err = appres.CreateAttribute(db.Id, col.Id, attr)
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
//
// Environment Setup:
//
// Before using this package, create a .env.local file in your project root with:
//
//	APPWRITE_ENDPOINT_URL=https://your-appwrite-endpoint.com/v1
//	APPWRITE_PROJECT_ID=your-project-id
//	APPWRITE_API_KEY_APPRES=your-api-key
//
// The API key should have all permissions on database and storage objects in Appwrite.
package appres
