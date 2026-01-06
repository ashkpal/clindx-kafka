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

// ReportHandler is implemented by the application
type ReportHandler func(ctx context.Context, event *models.TestReportEvent) error

func ConsumeTestReports(
	ctx context.Context,
	cfg Config,
	handler ReportHandler,
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

	log.Printf("[Kafka] Consuming test reports from topic=%s", cfg.Topic)

	backoff := time.Second

	for {
		select {
		case <-ctx.Done():
			log.Println("[Kafka] TestReport consumer shutting down")
			return nil

		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return nil
				}
				log.Printf("[Kafka] read error: %v", err)
				time.Sleep(backoff)
				backoff = min(backoff*2, time.Minute)
				continue
			}
			backoff = time.Second

			var event models.TestReportEvent
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

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
