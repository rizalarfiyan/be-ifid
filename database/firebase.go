package database

import (
	"be-ifid/config"
	"context"
	"encoding/json"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

var (
	client *db.Client
)

func Init() {
	ctx := context.Background()
	conf := config.Get()

	firebaseConfig := &firebase.Config{
		DatabaseURL: conf.FirebaseConfig.DatabaseUrl,
	}

	credentials, err := json.Marshal(conf.FirebaseCredential)
	if err != nil {
		log.Fatalln("Something wrong in firebase credentials")
	}

	opt := option.WithCredentialsJSON(credentials)
	app, err := firebase.NewApp(ctx, firebaseConfig, opt)
	if err != nil {
		log.Fatalln("error in initializing firebase app: ", err)
	}

	client, err = app.Database(ctx)
	if err != nil {
		log.Fatalln("error in creating firebase DB client: ", err)
	}
}

func GetClient() *db.Client {
	return client
}
