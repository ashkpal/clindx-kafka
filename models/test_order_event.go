package models

import "time"

type TestOrderEvent struct {
	EventVersion string    `json:"event_version"`
	EventType    string    `json:"event_type"`
	OccurredAt   time.Time `json:"occurred_at"`

	OrderID    uint   `json:"order_id"`
	TRFNum     string `json:"trf_num,omitempty"`
	PatientID  uint   `json:"patient_id"`
	PracticeID uint   `json:"practice_id"`
	Status     string `json:"status"`
}
