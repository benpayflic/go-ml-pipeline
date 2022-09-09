package amqp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/internal/ports"
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/config"
	"github.com/benpayflic/go-ml-pipeline/data-ingestion-service/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Adapter struct {
	api    ports.APIPort
	config *config.Config
}

func NewAdapter(api ports.APIPort, config *config.Config) *Adapter {
	return &Adapter{api: api, config: config}
}

var (
	uri           string
	exchange      string
	exchangeType      = "topic"
	consumerTag       = "data-ingestion-service-ctag"
	autoAck           = false
	ErrLog            = log.New(os.Stderr, "[AMQP][ERROR] ", log.LstdFlags|log.Lmsgprefix)
	Log               = log.New(os.Stdout, "[AMQP][INFO] ", log.LstdFlags|log.Lmsgprefix)
	deliveryCount int = 0
)

func (adapter Adapter) Consume() {

	uri = adapter.config.AMQP_URI
	exchange = adapter.config.AMQP_EXCHANGE

	c, err := newConsumer(adapter, uri, exchange, exchangeType, consumerTag)
	if err != nil {
		ErrLog.Fatalf("%s", err)
	}

	Log.Printf("Consuming ...")

	<-c.done
}

type Consumer struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	tag     string
	done    chan error
	adapter Adapter
}

func newConsumer(adapter Adapter, amqpURI, exchange, exchangeType, ctag string) (*Consumer, error) {
	c := &Consumer{
		conn:    nil,
		channel: nil,
		tag:     ctag,
		done:    make(chan error),
		adapter: adapter,
	}

	var err error

	config := amqp.Config{Properties: amqp.NewConnectionProperties()}
	config.Properties.SetClientConnectionName("sample-consumer")
	Log.Printf("dialing %q", amqpURI)
	c.conn, err = amqp.DialConfig(amqpURI, config)
	if err != nil {
		return nil, fmt.Errorf("dial: %s", err)
	}

	go func() {
		Log.Printf("closing: %s", <-c.conn.NotifyClose(make(chan *amqp.Error)))
	}()

	Log.Printf("got Connection, getting Channel")
	c.channel, err = c.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("channel: %s", err)
	}

	Log.Printf("got Channel, declaring Exchange (%q)", exchange)
	if err = c.channel.ExchangeDeclare(
		exchange,     // name of the exchange
		exchangeType, // type
		true,         // durable
		false,        // delete when complete
		false,        // internal
		false,        // noWait
		nil,          // arguments
	); err != nil {
		return nil, fmt.Errorf("exchange declare: %s", err)
	}

	consumers := make([]map[string]string, 0)
	err = json.Unmarshal([]byte(adapter.config.AMQP_CONSUMERS), &consumers)
	if err != nil {
		return nil, fmt.Errorf("unable to unmarshal consumer definitions: %s", err)
	}

	for _, consumer := range consumers {
		Log.Printf("declaring Queue %q on exhange %v", consumer["QUEUE"], exchange)

		queue, err := c.channel.QueueDeclare(
			consumer["QUEUE"], // name of the queue
			true,              // durable
			false,             // delete when unused
			false,             // exclusive
			false,             // noWait
			nil,               // arguments
		)
		if err != nil {
			return nil, fmt.Errorf("queue declare: %s", err)
		}

		Log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
			queue.Name, queue.Messages, queue.Consumers, consumer["ROUTING_KEY"])

		if err = c.channel.QueueBind(
			queue.Name,              // name of the queue
			consumer["ROUTING_KEY"], // bindingKey
			exchange,                // sourceExchange
			false,                   // noWait
			nil,                     // arguments
		); err != nil {
			return nil, fmt.Errorf("queue bind: %s", err)
		}

		Log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", c.tag)
		deliveries, err := c.channel.Consume(
			queue.Name, // name
			fmt.Sprintf("%v.%v", c.tag, utils.RandStringBytesMaskImprSrc(20)), // consumerTag,
			autoAck, // autoAck
			false,   // exclusive
			false,   // noLocal
			false,   // noWait
			nil,     // arguments
		)
		if err != nil {
			return nil, fmt.Errorf("queue consume: %s", err)
		}

		go c.handle(deliveries)

	}

	return c, nil
}

func (c *Consumer) handle(deliveries <-chan amqp.Delivery) {

	for d := range deliveries {
		deliveryCount++
		Log.Printf(
			"got %dB routing key: %v delivery: [%v] %q",
			len(d.Body),
			d.RoutingKey,
			d.DeliveryTag,
			d.Body,
		)

		Log.Println("Determining message handler...")

		handlers := NewMessageHandlers(c.adapter)
		handled, err := handlers.Handle(d)
		if err != nil {
			ErrLog.Println(err)
			d.Reject(false)
		}
		if !handled {
			d.Reject(false)
		}

		if !autoAck {
			d.Ack(false)
		}
	}
}
