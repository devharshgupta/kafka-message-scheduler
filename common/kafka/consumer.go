package kafka

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func consumeMessages(consumer *kafka.Consumer, topics []string, pollTimeoutMs int) {

	if pollTimeoutMs == 0 {
		pollTimeoutMs = 100
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := consumer.Poll(pollTimeoutMs)
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Received message from topic %s [%d] at offset %v: %s\n",
					*e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset, string(e.Value))
			case kafka.Error:
				fmt.Printf("Error🔥 : %v\n", e)
				run = false
			}
		}
	}
}

func InitConsumer(topics []string, timeoutMs int) {

	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
		"group.id":          os.Getenv("KAKFA_GROUP_ID"),
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		fmt.Printf("Error creating consumer: %v\n", err)
	}

	defer consumer.Close()

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		fmt.Printf("Error subscribing to topic: %v\n", err)
	}

	consumeMessages(consumer, topics, timeoutMs)
}
