package appres

import (
	"github.com/appwrite/sdk-for-go/id"
	"github.com/appwrite/sdk-for-go/models"
)
func CreateBucket(buc BucketType) (*models.Bucket, error) {
	return AppwriteStorage.CreateBucket(
        id.Unique(),
        buc.Name,
        AppwriteStorage.WithCreateBucketPermissions(buc.Permissions),
        AppwriteStorage.WithCreateBucketFileSecurity(buc.FileSecurity),
        AppwriteStorage.WithCreateBucketEnabled(buc.Enabled),
        AppwriteStorage.WithCreateBucketMaximumFileSize(buc.MaxFileSize),
        AppwriteStorage.WithCreateBucketAllowedFileExtensions(buc.AllowedFileExtensions),
        AppwriteStorage.WithCreateBucketCompression(buc.Compression),
        AppwriteStorage.WithCreateBucketEncryption(buc.Encryption),
        AppwriteStorage.WithCreateBucketAntivirus(buc.Antivirus),
    )
}