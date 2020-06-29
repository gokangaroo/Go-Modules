package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ready chan bool
}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(consumer.ready)
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("GoRoutineId: %d, Message claimed: key = %s, value = %v, topic = %s, partition = %v, offset = %v", goId(), string(message.Key), string(message.Value), message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}
	return nil
}

func (consumer Consumer) Run(ctx context.Context, wg *sync.WaitGroup, consumerGroup sarama.ConsumerGroup, topics string) {
	wg.Add(1)
	defer wg.Done()
	for {
		if err := consumerGroup.Consume(ctx, strings.Split(topics, ","), &consumer); err != nil {
			log.Panicf("Error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
		consumer.ready = make(chan bool, 0)
	}
}

func main() {
	// 中断消费后继续消费的消息有序性, 只在同一个partition内才能保证
	broker := "localhost:9092"
	groupId := "demo"
	topics := "sarama"
	kafkaVersion := "2.4.0"

	version, err := sarama.ParseKafkaVersion(kafkaVersion)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	ctx, cancel := context.WithCancel(context.Background())
	consumerGroup, err := sarama.NewConsumerGroup([]string{broker}, groupId, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	wg := &sync.WaitGroup{}
	// 自定义Consumer结构体,需要实现ConsumerGroupHandler接口, 并保证并发安全
	consumer := Consumer{
		ready: make(chan bool, 0),
	}
	go consumer.Run(ctx, wg, consumerGroup, topics)

	<-consumer.ready

	// block here
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}

	cancel()
	wg.Wait()
	if err = consumerGroup.Close(); err != nil {
		log.Panicf("Error closing client: %v", err)
	}
}

// https://www.cnblogs.com/binHome/p/13052397.html
func goId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	idField := strings.Fields(stk)[0]
	id, _ := strconv.Atoi(idField)
	return id
}
