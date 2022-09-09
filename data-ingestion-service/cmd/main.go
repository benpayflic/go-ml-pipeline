package main

import (
	"log"
	"sync"

	leftAMQP "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/adapters/framework/primary/amqp"
	leftGRPC "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/adapters/framework/primary/grpc"
	rightDB "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/adapters/framework/secondary/database"
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/application/api"
	c "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/config"
)

var wg sync.WaitGroup

func main() {
	// config, err := c.LoadConfig("./pkg/config/")
	config, err := c.LoadConfig("/Users/bengoodenough/Documents/Portfolio-Projects/go-ml-pipeline/data-ingestion-service/pkg/config/")
	if err != nil {
		log.Fatalln("Failed to load application configurations : ", err)
	}

	dbDrivenAdapter, err := rightDB.NewAdapter(&config)
	if err != nil {
		log.Fatalf("failed to initialize db driven adapter: %v", err)
	}
	dbDrivenAdapter.Connect()
	api := api.NewApplication(dbDrivenAdapter, config)

	wg.Add(2)
	go func() {
		defer wg.Done()
		grpcAdapter := leftGRPC.NewAdapter(api, &config)
		grpcAdapter.Start()
	}()

	go func() {
		defer wg.Done()
		amqpAdapter := leftAMQP.NewAdapter(api, &config)
		amqpAdapter.Consume()
	}()

	wg.Wait()
}
