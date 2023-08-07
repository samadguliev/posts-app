package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var Сlient *mongo.Client

func DbConnect() {
	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017/db")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Connection to Database failed")
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Connection to Database failed")
	}

	Сlient = client
}
