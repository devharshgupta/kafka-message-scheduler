package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/devharshgupta/kafka-message-scheduler/models"
	"github.com/devharshgupta/kafka-message-scheduler/repository"
)

type PushSeheduledMessageOptions struct {
	Key      string
	Value    string
	Priority int
}

func PushSeheduledMessage(data PushSeheduledMessageOptions) error {

	switch data.Priority {
	case -1:
		return ProduceMessage([]ProducerOptions{{
			Topic:   "Kafka_message_scheduler.P-1_Message",
			Key:     data.Key,
			Message: data.Value,
		}})
	case 1:
		ProduceMessage([]ProducerOptions{{
			Topic:   "Kafka_message_scheduler.P1_Message",
			Key:     data.Key,
			Message: data.Value,
		}})
	default:
		ProduceMessage([]ProducerOptions{{
			Topic:   "Kafka_message_scheduler.P0_Message",
			Key:     data.Key,
			Message: data.Value,
		}})

	}
	return nil

}

func MessageConsumer(data *kafka.Message) {
	switch *data.TopicPartition.Topic {
	case "Kafka_message_scheduler.P-1_Message":
		pushPmius1MessagetoDb(data)
	case "Kafka_message_scheduler.P1_Message":
		pushP1MessagetoDb(data)
	case "Kafka_message_scheduler.P0_Message":
		pushP0MessagetoDb(data)
	default:
		fmt.Printf("Received message from topic %s [%d] at offset %v: %s\n",
			*data.TopicPartition.Topic, data.TopicPartition.Partition, data.TopicPartition.Offset, string(data.Value))
	}
}

var pushPmius1MessageCount = 0
var pushPmius1Message []models.Message

func pushPmius1MessagetoDb(data *kafka.Message) {
	var message models.Message
	if err := json.Unmarshal(data.Value, &message); err != nil {
		log.Fatal("error while consuming p-1 message", err)
		return
	}
	pushPmius1Message = append(pushPmius1Message, message)
	pushPmius1MessageCount++

	if pushPmius1MessageCount > 1 { // set this count according to your message frquency and db wirte speed for eg this can be 100 if you are receving 1000 messages per second
		err := repository.CreateMessages(pushPmius1Message)
		if err != nil {
			log.Fatal("error while pusing p-1 message to db", err)
			return
		}
		pushPmius1Message = []models.Message{}
		pushPmius1MessageCount = 0
	}

}

var pushP1MessageCount = 0
var pushP1Message []models.Message

func pushP1MessagetoDb(data *kafka.Message) {
	var message models.Message
	if err := json.Unmarshal(data.Value, &message); err != nil {
		log.Fatal("error while consuming p1 message", err)
		return
	}
	pushP1Message = append(pushP1Message, message)
	pushP1MessageCount++

	if pushP1MessageCount > 2000 {
		err := repository.CreateMessages(pushPmius1Message)
		if err != nil {
			log.Fatal("error while pusing p1 message to db", err)
			return
		}
		pushP1Message = []models.Message{}
		pushP1MessageCount = 0
	}
}

var pushP0MessageCount = 0
var pushP0Message []models.Message

func pushP0MessagetoDb(data *kafka.Message) {

	var message models.Message
	if err := json.Unmarshal(data.Value, &message); err != nil {
		log.Fatal("error while consuming p0 message", err)
		return
	}
	pushP0Message = append(pushP0Message, message)
	pushP0MessageCount++

	if pushP0MessageCount > 500 {
		err := repository.CreateMessages(pushPmius1Message)
		if err != nil {
			log.Fatal("error while pusing p0 message to db", err)
			return
		}
		pushP0Message = []models.Message{}
		pushP0MessageCount = 0

	}
}
