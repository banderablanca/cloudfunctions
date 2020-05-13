package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
)

var ctx = context.Background()
var projectID = os.Getenv("GCLOUD_PROJECT")

func initConfig() (*firebase.App, error) {
	conf := &firebase.Config{ProjectID: projectID}
	app, err := firebase.NewApp(ctx, conf)

	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
		return nil, err
	}

	return app, nil
}

// GetDbInstance get a instance of firestore
func GetDbInstance() *firestore.Client {
	app, err := initConfig()
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	db, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("app.Firestore: %v", err)
	}

	return db
}

// GetMessagingInstance get a instance of Messaging
func GetMessagingInstance() *messaging.Client {
	app, err := initConfig()
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}

	messaging, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("app.Messaging: %v", err)
	}

	return messaging
}

// GetStorageInstance get a instance of Messaging
func GetStorageInstance() *storage.Client {
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Storage: %v", err)
	}

	return client
}

// GetBucketName get name of bucket
func GetBucketName() string {
	return fmt.Sprintf("%s.appspot.com", projectID)
}
