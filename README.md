# AppRes - Appwrite Resource Creator

A Go package for creating and managing Appwrite resources programmatically. Simplifies creating databases, collections, attributes, and storage buckets with built-in duplicate checking.

## Features

- **Databases**: Create with duplicate checking
- **Collections**: Create within databases with duplicate checking
- **Attributes**: Support for string, email, integer, datetime, boolean, relationship, and url types
- **Storage**: Create buckets with security and file constraints
- **Environment-based configuration**
- **Built-in error handling and logging**

## Installation

```bash
go get github.com/Haepapa/appres
```

## Setup

Create a `.env.local` file in your project root with:

```bash
APPWRITE_ENDPOINT_URL=https://your-appwrite-endpoint.com/v1
APPWRITE_PROJECT_ID=your-project-id
APPWRITE_API_KEY_APPRES=your-api-key  # API key with Database and Storage scopes
```

## Import

```go
import (
    "log"
    app "github.com/Haepapa/appres"
)
```

## Usage

```go
package main

import (
    "log"
    app "github.com/Haepapa/appres"
)

func main() {
    // Initialize Appwrite client
    app.Utils()

    // Create database
    db, err := app.CreateDatabase("my-database")
    if err != nil {
        log.Fatal(err)
    }

    // Create collection
    col, err := app.CreateCollection(db.Id, "users")
    if err != nil {
        log.Fatal(err)
    }

    // Create attribute
    attr := app.AttributeType{
        Type:     "string",
        Name:     "username",
        Size:     50,
        Required: true,
    }
    err = app.CreateAttribute(db.Id, col.Id, attr)
    if err != nil {
        log.Fatal(err)
    }

    // Create storage bucket
    bucket := app.BucketType{
        Name:         "user-uploads",
        Enabled:      true,
        FileSecurity: true,
        MaxFileSize:  10000000, // 10MB
    }
    buc, err := app.CreateBucket(bucket)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Resources created successfully!")
}
```

## API Reference

### Functions

| Function | Description |
|----------|-------------|
| `Utils()` | Initialize Appwrite client (required first) |
| `CreateDatabase(name)` | Create database with duplicate checking |
| `CreateCollection(dbId, name)` | Create collection with duplicate checking |
| `CreateAttribute(dbId, colId, attr)` | Create attribute with duplicate checking |
| `CreateBucket(bucket)` | Create storage bucket |

### Attribute Types

| Type | Fields | Example |
|------|--------|---------|
| `string` | `Size`, `Encrypt` | Text with optional encryption |
| `email` | `Size` | Email validation |
| `integer` | `Min`, `Max` | Numbers with constraints |
| `datetime` | | Date and time values |
| `boolean` | | True/false values |
| `relationship` | `RelatedCollectionID`, `RelationshipType` | Link collections |
| `url` | | URL validation |

**Common Fields**: `Type`, `Name`, `Required`, `Default`, `Array`

## Environment Variables

| Variable | Description |
|----------|-------------|
| `APPWRITE_ENDPOINT_URL` | Your Appwrite server endpoint URL |
| `APPWRITE_PROJECT_ID` | Your Appwrite project ID |
| `APPWRITE_API_KEY_APPRES` | API key with Database and Storage permissions |

## Requirements

- Go 1.22.5 or later
- Active Appwrite server instance  
- Valid API key with appropriate permissions

## License

Licensed under the terms in the LICENSE file.

## Contributing

Contributions welcome! Please submit a Pull Request.
