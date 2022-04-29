package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/streadway/amqp"
)

func main() {
	// define rabbitmq server url
	amqpServerUrl := os.Getenv("AMQP_SERVER_URL")

	// create a new RabbitMQ connection

	connectionRabbitMQ, err := amqp.Dial(amqpServerUrl)
	if err != nil {
		panic(err)
	}

	defer connectionRabbitMQ.Close()

	// opening a channel to RabbitMQ over the connection established
	channelRabbitMQ, err := connectionRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}

	defer channelRabbitMQ.Close()

	// with the channel we declare the queue that we can publish and subscribe to

	_, err = channelRabbitMQ.QueueDeclare(
		"onos-queue1", // the queue name
		true,          // durable
		false,         // autodelete
		false,         // exclusive
		false,         // nowait
		nil,           // other args
	)

	if err != nil {
		panic(err)
	}

	// creating new fiber instance

	app := fiber.New()

	// add a middleware
	app.Use(
		logger.New(), // a basic loggger
	)

	//add a route to fiber
	app.Get("/", func(c *fiber.Ctx) error {
		// create a message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(c.Query("msg")),
		}

		// Attempt to publish a message to the queue

		if err := channelRabbitMQ.Publish("", "onos-queue1", false, false, message); err != nil {
			return err
		}

		return nil
	})

	//starting the Fiber API server
	log.Fatal(app.Listen(":3000"))

}
