package amqp

import (
	jp "github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/application/domain/job_posting"
	amqp "github.com/rabbitmq/amqp091-go"
)

type CreateJobPostingsMessageHandler struct {
	adapter    Adapter
	routingKey string
}

func NewCreateJobPostingsMessageHandler(adapter Adapter) *CreateJobPostingsMessageHandler {
	return &CreateJobPostingsMessageHandler{adapter: adapter, routingKey: "jobposting.cmd.create"}
}

func (jh *CreateJobPostingsMessageHandler) HandleMessage(message amqp.Delivery) (bool, error) {
	if message.RoutingKey != jh.routingKey {
		return false, nil
	}
	Log.Println("CreateJobPostingsMessageHandler handling message with routing key: ", jh.routingKey)

	postings, err := jp.UnmarshalJobPostings(message.Body)
	if err != nil {
		return true, err
	}
	jh.adapter.api.CreateJobPostings(&postings)
	if err != nil {
		return true, err
	}

	Log.Println("Created Job Postings!")
	return true, nil
}
