package flag

import (
	"banderablancafunctions/config"
	"banderablancafunctions/models"
	"banderablancafunctions/utils"
	"context"
	"fmt"
	"net/http"

	"cloud.google.com/go/firestore"
)

var ctx = context.Background()
var db = config.GetDbInstance()
var clientStorage = config.GetStorageInstance()

// GetByID get flag by ID
func GetByID(flagID string) <-chan models.Flag {
	flag := make(chan models.Flag)

	go func() {
		defer close(flag)

		dsnap, _ := db.Collection("flags").Doc(flagID).Get(ctx)

		var f models.Flag
		dsnap.DataTo(&f)

		flag <- f
	}()

	return flag
}

// Update flag
func Update(flagID string, data map[string]interface{}) <-chan error {
	err := make(chan error)

	go func() {
		defer close(err)

		_, e := db.Collection("flags").Doc(flagID).Set(ctx, data, firestore.MergeAll)
		err <- e
	}()

	return err
}

// UploadImageMarked upload the flag's image marked with logo & get storage URL
func UploadImageMarked(flagID string, flag models.Flag) (string, error) {
	res, err := http.Get(flag.MediaContent.DownloadURL)

	if err != nil {
		return "Error don't get image from URL", err
	}

	resourceOriginal := res.Body
	defer resourceOriginal.Close()

	imageMarked := utils.SetLogoToImage(resourceOriginal, flag.MediaContent.MimeType)

	imageMarkedURL, err := utils.UploadImage(&imageMarked, fmt.Sprintf("flags/%s/imgbanderablanca.jpg", flagID))
	if err != nil {
		return "Error upload image marked", err
	}

	return imageMarkedURL, nil
}
