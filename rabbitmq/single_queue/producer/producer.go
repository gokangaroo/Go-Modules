package main

import (
	"log"
	"rabbitmq"
	"time"
)

func main() {
	send_mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "hello")
	for {
		time.Sleep(1 * time.Second)
		msg := "Hello World!"
		log.Printf("Sent a message: %s", msg)
		send_mq.Send(msg)
	}
}
