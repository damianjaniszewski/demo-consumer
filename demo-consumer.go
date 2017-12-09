package main

import (
	"os"
	"log"
	"time"
	"strconv"

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

	q, err := channel.QueueDeclare(os.Getenv("queueName"), true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Consumer started")
	log.Printf("> queue:%s\n> uri:%s\n", q.Name, os.Getenv("RABBITMQ_URL"))

	for {
    // q, err := channel.QueueInspect("demo-queue")
    // if err != nil {
		// 	log.Println("inspect error", err)
		// }
    // log.Println("queue length:", q.Messages)

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
				num, _ := strconv.ParseInt(string(msg.Body), 10, 64)

				start := time.Now()
				res := Compute(num)
				log.Printf("compute completed: fib(%d)=%d (%s)", num, res, time.Since(start))
			}

		} else {
			// log.Println("queue empty")
			time.Sleep(2 * time.Second)
		}
	}
}
