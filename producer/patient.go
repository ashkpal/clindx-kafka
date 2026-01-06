package producer

import (
	"context"
	"fmt"

	"github.com/ashkpal/clindx-kafka/models"
)

func (p *Producer) PublishPatient(
	ctx context.Context,
	event *models.PatientEvent,
) error {
	key := fmt.Sprintf("Patient:%d", event.PatientID)
	return p.publish(ctx, key, event)
}
