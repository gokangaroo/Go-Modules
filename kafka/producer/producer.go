package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	// kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 2 --topic sarama
	broker := "localhost:9092"
	topic := "sarama"
	kafkaVersion := "2.4.0"

	version, err := sarama.ParseKafkaVersion(kafkaVersion)
	if err != nil {
		log.Panicf("Error parsing Kafka version: %v", err)
	}

	config := sarama.NewConfig()
	config.Version = version
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 异步生产者
	producer, err := sarama.NewAsyncProducer([]string{broker}, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	chars := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	var enqueued, errors int
	doneCh := make(chan struct{})
	go func() {
		for {

			time.Sleep(1 * time.Second)

			// 随机串
			buf := make([]byte, 4)
			for i := 0; i < 4; i++ {
				buf[i] = chars[rand.Intn(len(chars))]
			}

			// 时间戳作为key
			strTime := strconv.Itoa(int(time.Now().Unix()))
			msg := &sarama.ProducerMessage{
				Topic: topic,
				Key:   sarama.StringEncoder(strTime),
				Value: sarama.StringEncoder(buf),
			}
			select {
			case producer.Input() <- msg:
				enqueued++
				fmt.Printf("Produce message: %s\n", buf)
			case err := <-producer.Errors():
				errors++
				fmt.Println("Failed to produce message:", err)
			case <-signals:
				doneCh <- struct{}{}
			}
		}
	}()

	<-doneCh
	log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}
