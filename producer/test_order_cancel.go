package producer

import (
	"context"
	"fmt"

	"github.com/ashkpal/clindx-kafka/models"
)

func (p *Producer) PublishTestOrderCancelation(
	ctx context.Context,
	event *models.TestOrderCancelEvent,
) error {
	key := fmt.Sprintf("Order:%d", event.OrderID)

	return p.publish(ctx, key, event)
}
