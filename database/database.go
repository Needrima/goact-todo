package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func initDB() *mongo.Database {
	uri := "mongodb://localhost:27017"
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error getting client oprions:", err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal("Error pinging database:", err)
	}

	return client.Database("goact-todo")
}

func TodosCollection() *mongo.Collection {
	return initDB().Collection("todos")
}
