package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (dba *Adapter) Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(dba.config.MongoURI))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database !")
	dba.collection = client.Database(dba.config.MongoDBame).Collection(dba.config.MongoCollectionName)
}
