package producer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(cfg Config) (*Producer, error) {
	mechanism, err := scram.Mechanism(
		scram.SHA512,
		cfg.Username,
		cfg.Password,
	)
	if err != nil {
		return nil, err
	}

	dialer := &kafka.Dialer{
		Timeout:       30 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      cfg.Brokers,
		Topic:        cfg.Topic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: -1,
		Dialer:       dialer,
	})

	return &Producer{writer: writer}, nil
}

func (p *Producer) publish(ctx context.Context, key string, value any) error {
	payload, err := json.Marshal(value)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(key),
		Value: payload,
		Time:  time.Now(),
	}

	return p.writer.WriteMessages(ctx, msg)
}
