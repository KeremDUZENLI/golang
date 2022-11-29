package main

import (
	"fmt"
	"rabbitmq/customer"

	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("RabbitMQ")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("CONNECTED!")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"Test Queu",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("HELLO"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("SUCCESSFUL")

	customer.Customer()
}
