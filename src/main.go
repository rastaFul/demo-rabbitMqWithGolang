package main

import (
	"lib"
)

func main () {
	myRabbit := &lib.RabbitMQ{}
	myRabbit.RabbitConnect("localhost", "5672", "guest", "guest")
	myRabbit.RabbitCreateChannel()
	myRabbit.RabbitQueueDeclare("Teste")
	myRabbit.RabbitSendMessage("Olá Mundo!")
}