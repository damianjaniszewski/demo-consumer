package main

import (
	"os"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	connection, err := amqp.Dial(os.Getenv("RABBITMQ_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer connection.Close()

	channel, err := connection.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer channel.Close()

	q, err := channel.QueueDeclare("demo-queue", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Consumer started, queue:%s, uri:%s\n", q.Name, os.Getenv("RABBITMQ_URL"))

	for {
		msg, ok, err := channel.Get(q.Name, false)
		if err != nil {
			log.Println("get error", err)
		}

		if ok {
			log.Printf("message received: %s\n", msg.Body)
			err := msg.Ack(false)
			if err != nil {
				log.Println("ack error", err)
			} else {
				start := time.Now()
				log.Printf("compute completed: %s", time.Since(start))
			}

		} else {
			log.Println("queue empty")
			time.Sleep(2 * time.Second)
		}
	}
}
