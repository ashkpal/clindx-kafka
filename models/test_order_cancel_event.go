package models

import (
	"time"
)

type TestOrderCancelEvent struct {
	EventVersion string    `json:"event_version"`
	EventType    string    `json:"event_type"`
	OccurredAt   time.Time `json:"occurred_at"`

	OrderID   uint `json:"order_id"`
	PatientID uint `json:"patient_id"`
}
