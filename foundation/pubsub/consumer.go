package pubsub

import (
	"fmt"

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

	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, err
	}
	return &kafkaConsumer{groupID: groupID, consumer: c}, nil
}

func (kc *kafkaConsumer) Consume(topic string, errChan chan error) <-chan string {
	outputChan := make(chan string)

	topicList := []string{topic}

	err := kc.consumer.SubscribeTopics(topicList, nil)
	if err != nil {
		errChan <- fmt.Errorf("failed to subscribe topics: %v", err)
		close(outputChan)
		return outputChan
	}
	go kc.consumeMessages(errChan, outputChan)

	return outputChan
}

func (kc *kafkaConsumer) consumeMessages(errChan chan error, outputChan chan<- string) {
	defer close(outputChan)

	for {
		msg, err := kc.consumer.ReadMessage(-1)
		if err != nil {
			errChan <- fmt.Errorf("failed to read message: %v", err)
			continue
		}
		outputChan <- fmt.Sprintf("Message on %s: %s", msg.TopicPartition, string(msg.Value))

	}
}

func (kc *kafkaConsumer) Close() error {
	return kc.consumer.Close()
}
