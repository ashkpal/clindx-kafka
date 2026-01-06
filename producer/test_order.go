package producer

import (
	"context"
	"fmt"

	"github.com/ashkpal/clindx-kafka/models"
)

func (p *Producer) PublishTestOrder(
	ctx context.Context,
	event *models.TestOrderEvent,
) error {
	key := fmt.Sprintf("Order:%d", event.OrderID)
	if event.TRFNum != "" {
		key = fmt.Sprintf("Order:%s", event.TRFNum)
	}
	return p.publish(ctx, key, event)
}
