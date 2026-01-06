package producer

import (
	"context"
	"fmt"

	"github.com/ashkpal/clindx-kafka/models"
)

func (p *Producer) PublishPractice(
	ctx context.Context,
	event *models.PracticeEvent,
) error {
	key := fmt.Sprintf("Practice:%d", event.PracticeID)
	return p.publish(ctx, key, event)
}
