package appres

import (
	"errors"

	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
	"github.com/appwrite/sdk-for-go/storage"
)
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