package main

import (
	"log"
	"rabbitmq"
	"time"
)

func main() {
	mq := rabbitmq.New("amqp://guest:guest@127.0.0.1:5672/", "hello")
	rabbitmq.NewExchange("amqp://guest:guest@127.0.0.1:5672/", "exchange", "direct")
	// 绑定队列=> direct就是routing,全值匹配模式
	// topic的话支持通配符: *只匹配一个单词, #匹配0或多个字符(与一般的通配符匹配还有一定区别)
	// 当topic模式下, queue的bind key为#的话, 就等同fanout, 如果没有*和#, 就等同direct
	mq.Bind("exchange", "key")
	for {
		time.Sleep(1 * time.Second)
		msg := "Hello World!"
		log.Printf("Sent a message: %s", msg)
		mq.Publish("exchange", msg, "key")
	}
}
