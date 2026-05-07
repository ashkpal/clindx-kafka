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
	log.Println("dialer init done")

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
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				log.Printf("[Kafka] Test Report read error: %+v", err)
				log.Printf("username=%s", cfg.Username)
				if ctx.Err() != nil {
					return nil
				}
				time.Sleep(backoff)
				backoff = min(backoff*2, time.Minute)
				continue
			}
			backoff = time.Second

			var event models.TestReportEvent
			if err := json.Unmarshal(msg.Value, &event); err != nil {
				log.Printf("[Kafka] invalid JSON topic=%s partition=%d offset=%d: %v",
					msg.Topic, msg.Partition, msg.Offset, err,
				)

				// Commit JSON so it does not block the partition forever.
				if err := reader.CommitMessages(ctx, msg); err != nil {
					log.Printf("[Kafka] failed to commit invalid JSON offset: %v", err)
				}
				continue
			}

			if err := handler(ctx, &event); err != nil {
				log.Printf("[Kafka] handler error topic=%s partition=%d offset=%d key=%s: %v",
					msg.Topic, msg.Partition, msg.Offset, string(msg.Key), err,
				)

				time.Sleep(backoff)
				backoff = min(backoff*2, time.Minute)

				// Do NOT commit. Message will be retried.
				continue
			}
			backoff = time.Second

			if err := reader.CommitMessages(ctx, msg); err != nil {
				log.Printf("[Kafka] failed to commit offset topic=%s partition=%d offset=%d: %v",
					msg.Topic, msg.Partition, msg.Offset, err,
				)
				continue
			}

			log.Printf("[Kafka] committed topic=%s partition=%d offset=%d",
				msg.Topic, msg.Partition, msg.Offset,
			)
		}
	}
}

func min(a, b time.Duration) time.Duration {
	if a < b {
		return a
	}
	return b
}
