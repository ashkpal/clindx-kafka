package models

type PhysicianEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PhysicianID uint `json:"physician_id"`
}
