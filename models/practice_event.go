package models

import "github.com/google/uuid"

type PracticeEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PracticeID     uint      `json:"practice_id"`
	Name           string    `json:"name"`
	LIMSPracticeID uuid.UUID `json:"lims_practice_id"`

	Locations []PracticeLocation `json:"locations"`
}

type PracticeLocation struct {
	PracticeID uint   `json:"practice_id"`
	Name       string `json:"name"`
	Address1   string `json:"Address1"`
	Address2   string `json:"Address2"`
	City       string `json:"city"`
	State      string `json:"state"`
	Zip        string `json:"zip"`
	Phone      string `json:"phone"`
	Fax        string `json:"fax"`
	SiteNum    string `json:"site_num"`
	IsPrimary  bool   `json:"is_primary"`
}
