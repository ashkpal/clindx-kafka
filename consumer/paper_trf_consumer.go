package consumer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"log"
	"time"

	"github.com/ashkpal/clindx-kafka/models"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

// PaperTRFHandler is implemented by the application
type PaperTRFHandler func(ctx context.Context, event *models.TestOrderEvent) error

func ConsumePaperTRF(
	ctx context.Context,
	cfg Config,
	handler PaperTRFHandler,
) error {

	mechanism, err := scram.Mechanism(
		scram.SHA512,
		cfg.Username,
		cfg.Password,
	)
	if err != nil {
		return err
	}

	dialer := &kafka.Dialer{
		Timeout:       30 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           &tls.Config{},
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Brokers,
		Topic:   cfg.Topic,
		GroupID: cfg.GroupID,
		Dialer:  dialer,
	})
	defer reader.Close()

	log.Printf("[Kafka] Consuming paper TRF from topic=%s", cfg.Topic)

	backoff := time.Second

	for {
		select {
		case <-ctx.Done():
			log.Println("[Kafka] Paper TRF consumer shutting down")
			return nil

		default:
			msg, err := reader.ReadMessage(ctx)

			if err != nil {
				log.Printf("[Kafka] Paper TRF read error: %+v", err)
				log.Printf("username=%s", cfg.Username)
				if ctx.Err() != nil {
					return nil
				}
				log.Printf("[Kafka] Paper TRF read error: %v", err)
				time.Sleep(backoff)
				backoff = min(backoff*2, time.Minute)
				continue
			}
			backoff = time.Second

			log.Printf("Raw message: %s", string(msg.Value))
			var event models.TestOrderEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("[Kafka] invalid JSON: %v", err)
				continue
			}

			if err := handler(ctx, &event); err != nil {
				log.Printf("[Kafka] handler error (key=%s): %v", string(msg.Key), err)
				// at-least-once semantics: do NOT crash
				continue
			}
		}
	}
}
