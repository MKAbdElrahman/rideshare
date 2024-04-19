package pubsub

type PubSub struct {
	*kafkaProducer
	*kafkaConsumer
}

type PubSubConfig struct {
	Brokers string
	GroupID string
}

func NewPubSub(cfg PubSubConfig) (*PubSub, error) {
	p, err := newKafkaProducer(cfg.Brokers)
	if err != nil {
		return nil, err
	}
	c, err := newKafkaConsumer(cfg.Brokers, cfg.GroupID)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		kafkaProducer: p,
		kafkaConsumer: c,
	}, nil

}
