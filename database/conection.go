package database

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/db"
	"google.golang.org/api/option"
)

var FirebaseClient *db.Client

func Connect() {
	// Initialize Firebase
	opt := option.WithCredentialsFile("../xunami_mobile/xunami-userbase-firebase-adminsdk-tlmx5-15eeb96930.json") // Replace with your Firebase service account key path
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Get a database client from the Firebase app
	client, err := app.DatabaseWithURL(context.Background(), "https://your-database-name.firebaseio.com/") // Replace with your Firebase database URL
	if err != nil {
		log.Fatalf("error connecting to Firebase Database: %v\n", err)
	}

	FirebaseClient = client

	fmt.Println("Successfully connected to Firebase Database")
}
