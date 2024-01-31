package kafka

import (
	"context"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// EnsureKafkaTopicsExitInit ensures that Kafka topics exist, creating them if needed.
func EnsureKafkaTopicsExitInit(topics []string) error {
	admin, err := kafka.NewAdminClient(&kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_HOST"),
	})

	if err != nil {
		return fmt.Errorf("error creating admin client: %v", err)
	}

	defer admin.Close()

	// Check if topics already exist
	meta, err := admin.GetMetadata(nil, false, 5000)
	if err != nil {
		return fmt.Errorf("error getting metadata: %v", err)
	}

	var topicsToCreate []kafka.TopicSpecification

	for _, topic := range topics {
		if _, exists := meta.Topics[topic]; !exists {
			// Create a new topic specification
			newTopic := kafka.TopicSpecification{
				Topic:             topic,
				NumPartitions:     1, // Adjust as needed
				ReplicationFactor: 1, // Adjust as needed
			}
			topicsToCreate = append(topicsToCreate, newTopic)
		}
	}

	if len(topicsToCreate) > 0 {
		// Create new topics
		ctx := context.Background()
		_, err := admin.CreateTopics(ctx, topicsToCreate)
		if err != nil {
			return fmt.Errorf("error creating topics: %v", err)
		}

		fmt.Printf("Topics created successfully.\n")
	} else {
		fmt.Printf("All specified topics already exist.\n")
	}

	return nil
}
