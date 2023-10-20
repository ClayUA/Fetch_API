package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(uri string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	clientInstance = client
	log.Println("Connected to MongoDB!")
	return client, nil
}

func GetDBInstance() *mongo.Client {
	return clientInstance
}

func DisconnectDB() {
	if clientInstance != nil {
		err := clientInstance.Disconnect(context.TODO())
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Connection to MongoDB closed.")
	}
}
