package models

type PracticeEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PracticeID uint   `json:"practice_id"`
	Name       string `json:"name"`
}
