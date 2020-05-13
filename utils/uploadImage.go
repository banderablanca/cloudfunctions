package utils

import (
	"cloudfunctions/config"
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

var clientStorage = config.GetStorageInstance()
var ctx = context.Background()

// UploadImage upload image to storage & return URL
func UploadImage(reader io.Reader, path string) (string, error) {
	bk := clientStorage.Bucket(config.GetBucketName())
	if _, err := bk.Attrs(ctx); err != nil {
		return "Bucket isn't exist", err
	}

	obj := bk.Object(path)
	w := obj.NewWriter(ctx)
	if _, err := io.Copy(w, reader); err != nil {
		return "Error copy data to writer", err
	}

	if err := w.Close(); err != nil {
		return "Error close writer", err
	}

	if err := obj.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "Error config ACL", err
	}

	attrs, _ := obj.Attrs(ctx)
	return fmt.Sprintf("https://storage.googleapis.com/%s/%s", attrs.Bucket, attrs.Name), nil
}
