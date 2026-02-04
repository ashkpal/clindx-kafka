package models

import "github.com/google/uuid"

type PhysicianEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PhysicianID     uint      `json:"physician_id"`
	FirstName       string    `json:"first_name"`
	LastName        string    `json:"last_name"`
	NPI             string    `json:"npi"`
	Email           string    `json:"email"`
	Fax             string    `json:"fax"`
	LIMSPhysicianID uuid.UUID `gorm:"type:lims_physician_id"`
}
