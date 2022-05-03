package main

import (
	"log"
	"os"

	"github.com/streadway/amqp"
)

func main() {
	// define RabbitMQ server URL
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	// create new rabbitmq connection
	connectionRabbit, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}

	defer connectionRabbit.Close()

	// opening a channel to rabbit server
	channelRabbit, err := connectionRabbit.Channel()
	if err != nil {
		panic(err)
	}

	defer channelRabbit.Close()

	// subscribe to onos channel for getting messages
	messages, err := channelRabbit.Consume("onos-queue1", "", true, false, false, false, nil)
	if err != nil {
		log.Println(err)
	}

	// welcome message
	log.Println("Successfully connected to RabbitMQ")
	log.Println("Waiting for messages")

	// main loop

	forever := make(chan bool)

	go func() {
		for message := range messages {
			log.Printf("Message received from the broker: %s ", message.Body)
		}
	}()

	<-forever

}
