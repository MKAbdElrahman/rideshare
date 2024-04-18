package pubsub

import (
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type kafkaConsumer struct {
	groupID  string
	consumer *kafka.Consumer
}

func newKafkaConsumer(brokers string, groupID string) (*kafkaConsumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": brokers,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	}

	// Create a new Kafka consumer with the configuration
	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return &kafkaConsumer{groupID: groupID, consumer: c}, nil
}

func (kc *kafkaConsumer) Consume(topic string) error {
	topicList := []string{topic}

	err := kc.consumer.SubscribeTopics(topicList, nil)
	if err != nil {
		return errors.New("failed to subscribe topics: " + err.Error())
	}
	for {
		msg, err := kc.consumer.ReadMessage(time.Second)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		}
		if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}

// Close closes the Kafka consumer
func (kc *kafkaConsumer) Close() error {
	return kc.consumer.Close()
}
