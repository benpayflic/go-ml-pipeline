package api

import (
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/ports"
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/config"
)

// Application implements the APIPort interface
type Application struct {
	db     ports.DBPort
	config config.Config
}

// Return new instance of Application
func NewApplication(db ports.DBPort, config config.Config) *Application {
	return &Application{
		db:     db,
		config: config,
	}
}
