package producer

import (
	"context"
	"fmt"

	"github.com/ashkpal/clindx-kafka/models"
)

func (p *Producer) PublishPhysician(
	ctx context.Context,
	event *models.PhysicianEvent,
) error {
	key := fmt.Sprintf("Physician:%d", event.PhysicianID)
	return p.publish(ctx, key, event)
}
