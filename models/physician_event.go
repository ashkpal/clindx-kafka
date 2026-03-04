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

	PhysicianPractices []PhysicianPractice `json:"physician_practices"`
}

type PhysicianPractice struct {
	PhysicianID uint `json:"physician_id"`
	PracticeID  uint `json:"practice_id"`
	EmailFlag   bool `json:"email_flag"`
	FaxFlag     bool `json:"fax_flag"`
	PortalFlag  bool `json:"portal_flag"`
}
