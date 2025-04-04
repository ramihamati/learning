package main

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

type MqConnection struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	context    context.Context
	cancel     context.CancelFunc
}

type MyQueue struct {
	queue      amqp.Queue
	connection *MqConnection
}

func Connect(connectionString string) (*MqConnection, error) {
	conn, err := amqp.Dial("amqp://admin:dev1234!@100.103.106.35:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		return nil, err
	}
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	return &MqConnection{
		connection: conn,
		channel:    ch,
		context:    ctx,
		cancel:     cancel,
	}, nil
}

func (connection *MqConnection) Close() {
	connection.cancel()
	connection.channel.Close()
	connection.connection.Close()
}

func (connection *MqConnection) CreateQueue(name string) (*MyQueue, error) {
	// declare a queue
	q, err := connection.channel.QueueDeclare(
		"go-test",
		true,  // survives a server restart, it's persisted on disk and messages survive
		false, // automatically delete a queue when last consumer disconnects
		false, // restricting the queue to the connection creating it
		false, // waiting for a reply to the queue declaration.
		nil)

	failOnError(err, "Failed to declare a queue")

	if err != nil {
		return nil, err
	}

	return &MyQueue{
		queue:      q,
		connection: connection,
	}, nil
}

func (queue *MyQueue) Publish(contentType string, data []byte) error {
	// publishing to a queue directly
	err := queue.connection.channel.PublishWithContext(
		queue.connection.context,
		"",
		queue.queue.Name, // key
		false,            // IF TRUE AND MESSAGE CANNOT be routed to a queue, MQ will return it to the sender using a basic.return AMQP message
		false,            // deprecated and ignored
		amqp.Publishing{
			ContentType: contentType,
			Body:        data,
		})

	if err != nil {
		return err
	}

	log.Printf(" [x] Sent %s", data)

	return nil
}

func main() {
	conn, err := Connect("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	queue, err := conn.CreateQueue("go-test")
	failOnError(err, "Failed to create a queue")

	err = queue.Publish("text/plain", []byte("Hello World 2"))
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
