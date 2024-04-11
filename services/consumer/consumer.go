package eventslogger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumer struct {
	consumer *kafka.Consumer
	topics   []string
}
type Config struct {
	Kafka *kafka.ConfigMap
}

func NewConsumer(topics []string, cfg *kafka.ConfigMap) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(cfg)
	return &Consumer{
		consumer: consumer,
		topics:   topics,
	}, err
}

func (c *Consumer) Consume() error {
	err := c.consumer.SubscribeTopics(c.topics, nil)

	if err != nil {
		return err
	}
	run := true
	// const MIN_COMMIT_COUNT = 1
	// msg_count := 0
	for run {
		ev := c.consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			// msg_count += 1
			// if msg_count%MIN_COMMIT_COUNT == 0 {
			c.consumer.Commit()
			// }
			slog.Info("Message", "Topic", "users", "Body", string(e.Value))

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			// fmt.Printf("Ignored %v\n", e)
		}
	}

	defer c.consumer.Close()

	return nil
}
