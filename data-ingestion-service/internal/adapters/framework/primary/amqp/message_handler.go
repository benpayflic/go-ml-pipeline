package amqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageType interface {
	HandleMessage(amqp.Delivery) (bool, error)
}

type MessageHandlers struct {
	messageHandlers []MessageType
}

func NewMessageHandlers(adapter Adapter) *MessageHandlers {
	messageTypes := []MessageType{
		NewCreateJobPostingsMessageHandler(adapter), // Add message handlers here
	}
	return &MessageHandlers{messageTypes}
}

func (m *MessageHandlers) Handle(message amqp.Delivery) (bool, error) {
	for _, m := range m.messageHandlers {
		handled, err := m.HandleMessage(message)
		if err != nil {
			return true, err
		}
		if handled {
			return true, nil
		}
	}
	return false, nil
}
