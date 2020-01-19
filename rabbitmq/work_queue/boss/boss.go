package main

import (
	"fmt"
	"rabbitmq"
	"time"
)

func main() {
	send_mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "hello")
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		job := fmt.Sprintf("job id: %d", i)
		send_mq.Send(job)
	}
}
