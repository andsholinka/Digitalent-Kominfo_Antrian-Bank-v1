package model

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var client *db.Client
var ctx context.Context

func init() {
	ctx = context.Background()
	conf := &firebase.Config{
		DatabaseURL: "https://digitalent-be-23-baa6a.firebaseio.com/",
	}

	// Fetch the service account key JSON file contents
	opt := option.WithCredentialsFile("firebase-admin-sdk.json")

	// Initialize the app with service account, granting admin priviliges
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		log.Fatalln("Error initializing app: ", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client: ", err)
	}
}
