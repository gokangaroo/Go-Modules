package main

import (
	"context"
	"log"
	"rabbitmq"
)

func main() {
	receive_mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "hello")
	msgs := receive_mq.Consume()
	c, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			d.Ack(false)
		}
		cancel()
	}(c)
	<-c.Done()
}
