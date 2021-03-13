package storage

import (
	"context"
	"errors"

	"cloud.google.com/go/storage"
)

// ActivateImage - Migrate items from temporary bucket to image bucket.
func ActivateImage(imageHash string) error {
	ctx := context.Background()

	temp, err := Client.Bucket(Temp)
	if err != nil {
		return errors.New("Failed to setup temp bucket handler")
	}

	image, err := Client.Bucket(Image)
	if err != nil {
		return errors.New("Failed to setup image bucket handler")
	}

	src := temp.Object("image/" + imageHash)
	dst := image.Object(imageHash)

	if _, err := dst.CopierFrom(src).Run(ctx); err != nil {
		return errors.New("Failed to copy")
	}

	newAttrs := storage.ObjectAttrsToUpdate{
		CacheControl: "public, max-age=2592000, immutable",
		Metadata:     map[string]string{},
	}
	if _, err := dst.Update(ctx, newAttrs); err != nil {
		return errors.New("Failed to update")
	}

	if err := src.Delete(ctx); err != nil {
		return errors.New("Failed to delete")
	}

	return nil
}
