package database

import (
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type Adapter struct {
	collection *mongo.Collection
	config     *config.Config
}

func NewAdapter(config *config.Config) (*Adapter, error) {
	collection := &mongo.Collection{}
	return &Adapter{collection: collection, config: config}, nil
}
