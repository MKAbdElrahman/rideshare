package publisher

import "github.com/confluentinc/confluent-kafka-go/v2/kafka"

type Config struct {
	Bootstrap string
	Topic     string
}
type Publisher struct {
	producer *kafka.Producer
	cfg      Config
}

func NewPublisher(cfg Config) (*Publisher, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Bootstrap,
	})
	if err != nil {
		return nil, err
	}
	return &Publisher{
		producer: p,
		cfg:      cfg,
	}, nil
}

func (p *Publisher) Publish(msg string) error {
	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.cfg.Topic, Partition: kafka.PartitionAny},
		Value:          []byte(msg)},
		nil, // delivery channel
	)
	return err
}
