package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ErrLog = log.New(os.Stderr, "[MONGO][ERROR] ", log.LstdFlags|log.Lmsgprefix)
	Log    = log.New(os.Stdout, "[MONGO][INFO] ", log.LstdFlags|log.Lmsgprefix)
)

func (dba *Adapter) Connect() {
	client, err := mongo.NewClient(options.Client().ApplyURI(dba.config.MongoURI))
	if err != nil {
		ErrLog.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		ErrLog.Fatal(err)
	}

	Log.Println("Connected to database !")
	dba.collection = client.Database(dba.config.MongoDBame).Collection(dba.config.MongoCollectionName)
}
