# AppRes - Appwrite Resource Creator

A Go package for creating and managing Appwrite resources programmatically. This package simplifies the process of creating databases, collections, and attributes in your Appwrite backend.

## Features

- Create databases with duplicate checking
- Create collections with duplicate checking  
- Create string and email attributes with full configuration support
- Environment-based configuration
- Comprehensive error handling and logging

## Installation

Install the package using `go get`:

```bash
go get github.com/Haepapa/appres
```

## Setup

### 1. Environment Configuration

Create a `.env.local` file in your project root with your Appwrite configuration:

```env
NEXT_PUBLIC_APPWRITE_ENDPOINT=https://your-appwrite-endpoint.com/v1
NEXT_PUBLIC_APPWRITE_PROJECT=your-project-id
APPWRITE_API_KEY_RESDEF=your-api-key
```

### 2. Import the Package

```go
import (
    "crypto/tls"
    "log"
    "net/http"
    
    app "github.com/Haepapa/appres"
)
```

## Usage

### Basic Example

Here's a complete example showing how to create a database, collection, and attributes:

```go
package main

import (
    "crypto/tls"
    "log"
    "net/http"
    
    app "github.com/Haepapa/appres"
)

func main() {
    // Suppress insecure warning (if using self-signed certificates)
    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    // Initialize Appwrite client
    app.Utils()

    // Create a database
    db, err := app.CreateDatabase("synergysquares")
    if err != nil {
        log.Println("Error creating database:", err)
        return
    }

    // Create collection(s)
    colContactUs, err := app.CreateCollection(db.Id, "contact_us")
    if err != nil {
        log.Println("Error creating collection:", err)
        return
    }

    // Create attributes in collection(s)
    attVals := []app.AttributeType{
        {
            Type:        "string",
            Name:        "name",
            Size:        100,
            Required:    false,
            Default:     "",
            Array:       false,
            Encrypt:     false,
        },
        {
            Type:        "email",
            Name:        "email",
            Size:        200,
            Required:    false,
            Default:     "email@email.com",
            Array:       false,
            Encrypt:     false,
        },
        {
            Type:        "string",
            Name:        "subject",
            Size:        200,
            Required:    false,
            Default:     "",
            Array:       false,
            Encrypt:     false,
        },
        {
            Type:        "string",
            Name:        "message",
            Size:        5000,
            Required:    false,
            Default:     "",
            Array:       false,
            Encrypt:     false,
        },
    }

    for _, att := range attVals {
        err = app.CreateAttribute(db.Id, colContactUs.Id, att)
        if err != nil {
            log.Println("Error creating attribute:", err)
            return
        }
    }

    log.Println("Successfully created database, collection, and attributes!")
}
```

## API Reference

### Functions

#### `Utils()`
Initializes the Appwrite client with environment variables. Must be called before using other functions.

```go
app.Utils()
```

#### `CreateDatabase(name string) (*models.Database, error)`
Creates a new database or returns existing one if it already exists.

```go
db, err := app.CreateDatabase("my-database")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Database ID:", db.Id)
```

#### `CreateCollection(dbId string, name string) (*models.Collection, error)`
Creates a new collection in the specified database or returns existing one if it already exists.

```go
col, err := app.CreateCollection(db.Id, "my-collection")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Collection ID:", col.Id)
```

#### `CreateAttribute(dbID string, colID string, att AttributeType) error`
Creates a new attribute in the specified collection or skips if it already exists.

```go
attr := app.AttributeType{
    Type:     "string",
    Name:     "title",
    Size:     255,
    Required: true,
    Default:  "",
    Array:    false,
    Encrypt:  false,
}

err := app.CreateAttribute(db.Id, col.Id, attr)
if err != nil {
    log.Fatal(err)
}
```

### Types

#### `AttributeType`
Defines the structure for creating attributes:

```go
type AttributeType struct {
    Type     string // "string" or "email" (more types coming soon)
    Name     string // Attribute key/name
    Size     int    // Maximum size for the attribute
    Required bool   // Whether the attribute is required
    Default  string // Default value
    Array    bool   // Whether the attribute is an array
    Encrypt  bool   // Whether to encrypt the attribute
}
```

### Supported Attribute Types

- **string**: Text attributes with configurable size
- **email**: Email validation attributes
- More types coming soon (integer, boolean, datetime, etc.)

## Environment Variables

| Variable | Description | Required |
|----------|-------------|----------|
| `NEXT_PUBLIC_APPWRITE_ENDPOINT` | Your Appwrite server endpoint URL | Yes |
| `NEXT_PUBLIC_APPWRITE_PROJECT` | Your Appwrite project ID | Yes |
| `APPWRITE_API_KEY_RESDEF` | API key with appropriate permissions | Yes |

## Error Handling

All functions return appropriate errors that should be handled:

```go
db, err := app.CreateDatabase("test-db")
if err != nil {
    log.Printf("Failed to create database: %v", err)
    return
}
```

The package also provides detailed logging for debugging purposes.

## Requirements

- Go 1.22.5 or later
- Active Appwrite server instance
- Valid API key with database creation permissions

## License

This project is licensed under the terms included in the LICENSE file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
