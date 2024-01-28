package kafka

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
