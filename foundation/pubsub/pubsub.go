package pubsub

type PubSub struct {
	*kafkaProducer
	*kafkaConsumer
}

type PubSubConfig struct {
	Brokers string
	groupID string
}

func NewPubSub(cfg PubSubConfig) (*PubSub, error) {
	p, err := newKafkaProducer(cfg.Brokers)
	if err != nil {
		return nil, err
	}
	c, err := newKafkaConsumer(cfg.Brokers, cfg.groupID)
	if err != nil {
		return nil, err
	}

	return &PubSub{
		kafkaProducer: p,
		kafkaConsumer: c,
	}, nil

}
