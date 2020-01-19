package main

import (
	"context"
	"log"
	"rabbitmq"
	"time"
)

func main() {
	c, _ := context.WithCancel(context.Background())
	createConsumer(1)
	createConsumer(2)
	createConsumer(3)
	createConsumer(4)
	createConsumer(5)
	<-c.Done()
}

func createConsumer(id int) {
	receive_mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "hello")
	// Qos配合, 当前面的消息没有ack的时候, 是不会再发消息给这个jober了, 需要配合手动ack使用
	//receive_mq.Qos()
	msgs := receive_mq.Consume()
	log.Printf("create a new consumer, consumer id: %d", id)
	go func() {
		for d := range msgs {
			log.Printf("recevie %d Received a message: %s", id, d.Body)
			time.Sleep(time.Duration(id*2) * time.Second)
			d.Ack(false)
		}
	}()
}
