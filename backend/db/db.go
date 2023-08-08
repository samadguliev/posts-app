package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Client *mongo.Client

func Connect() {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	db, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Connection to Database failed: ", err)
	}

	err = db.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection to Database failed: ", err)
	}

	log.Println("connected")
	Client = db
}
