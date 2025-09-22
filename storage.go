package appres

import (
	"errors"

	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/storage"
)

// CreateBucket creates a new storage bucket with the specified configuration.
// It creates a bucket with customizable security, file size limits, and permissions.
//
// Parameters:
//   - buc: BucketType struct containing the bucket configuration
//
// Global Variables Used:
//   - AppwriteStorage: The initialized Appwrite storage client
//
// Returns:
//   - *models.Bucket: Pointer to the created bucket
//   - error: Any error that occurred during the operation
//
// Example:
//
//	bucket := appres.BucketType{
//		Name:         "my-bucket",
//		Enabled:      true,
//		FileSecurity: true,
//		MaxFileSize:  10000000, // 10MB
//		Permissions:  []string{"read(\"any\")"},
//	}
//	buc, err := appres.CreateBucket(bucket)
//	if err != nil {
//		log.Fatal("Failed to create bucket:", err)
//	}
func CreateBucket(buc BucketType) (*models.Bucket, error) {

	var opts []storage.CreateBucketOption

	opts = append(opts, AppwriteStorage.WithCreateBucketAntivirus(buc.FileSecurity))
	opts = append(opts, AppwriteStorage.WithCreateBucketAntivirus(buc.Enabled))
	opts = append(opts, AppwriteStorage.WithCreateBucketAntivirus(buc.Antivirus))
	opts = append(opts, AppwriteStorage.WithCreateBucketAntivirus(buc.Encryption))
	if buc.MaxFileSize < 0 || buc.MaxFileSize > 30000000 {
		return nil,  errors.New("MaxFileSize must be between 0 and 30MB")
	} else {
		opts = append(opts, AppwriteStorage.WithCreateBucketMaximumFileSize(buc.MaxFileSize))
	}
	if buc.Permissions != nil {
		opts = append(opts, AppwriteStorage.WithCreateBucketPermissions(buc.Permissions))
	}
	if buc.AllowedFileExtensions != nil {
		opts = append(opts, AppwriteStorage.WithCreateBucketAllowedFileExtensions(buc.AllowedFileExtensions))
	}
	if buc.Compression != "" {
		opts = append(opts, AppwriteStorage.WithCreateBucketCompression(buc.Compression))
	}


	return AppwriteStorage.CreateBucket(
        id.Unique(),
        buc.Name,
		opts...,
    )
}