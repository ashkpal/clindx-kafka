package producer

import (
	"context"
	"fmt"

	"clindx.com/kafka/models"
)

func (p *Producer) PublishPhysician(
	ctx context.Context,
	event *models.PhysicianEvent,
) error {
	key := fmt.Sprintf("Physician:%d", event.PhysicianID)
	return p.publish(ctx, key, event)
}
