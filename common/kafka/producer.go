// kafka_producer.go
package kafka

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type ProducerOptions struct {
	Topic     string
	Key       string
	Message   string
	Partition int32
}

func ProduceMessage(options []ProducerOptions) error {

	config := &kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
	}

	producer, err := kafka.NewProducer(config)
	if err != nil {
		fmt.Printf("Error creating producer: %v\n", err)
		return err
	}

	for _, option := range options {
		deliveryChan := make(chan kafka.Event)
		err := producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &option.Topic, Partition: kafka.PartitionAny},
			Key:            []byte(option.Key),
			Value:          []byte(option.Message),
		}, deliveryChan)

		if err != nil {
			fmt.Printf("Error producing message: %v\n", err)
		}

		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Delivered message to topic %s [%d] at offset %v\n", *m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}
		close(deliveryChan)
	}

	defer producer.Close()

	return nil
}
