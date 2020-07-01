package lib

import (
	"log"
	"fmt"
	
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel *amqp.Channel
	Queue amqp.Queue
	Err error
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func (rabbit *RabbitMQ) RabbitConnect(hostname string, port string, username string, password string) {
	host := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, hostname, port)
	rabbit.Connection, rabbit.Err = amqp.Dial(host)
	FailOnError(rabbit.Err, "Failed to connect to RabbitMQ")
	return
}

func (rabbit *RabbitMQ) RabbitCreateChannel() {
	rabbit.Channel, rabbit.Err = rabbit.Connection.Channel()
	FailOnError(rabbit.Err, "Failed to open a channel")
	return
}

func (rabbit *RabbitMQ) RabbitQueueDeclare(queueName string) {
	rabbit.Queue, rabbit.Err = rabbit.Channel.QueueDeclare(
		queueName, // name
		true,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil)       // arguments
	FailOnError(rabbit.Err, "Failed to declare a queue")
	return
}

func (rabbit *RabbitMQ) RabbitSendMessage(message string) {
	rabbit.Err = rabbit.Channel.Publish(
		"",                // exchange
		rabbit.Queue.Name, // routing key
		false,             // mandatory
		false,             // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	FailOnError(rabbit.Err, "Failed to publish a message")
	return
}
