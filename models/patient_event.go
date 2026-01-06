package models

type PatientEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PatientID uint   `json:"patient_id"`
	MRN       string `json:"mrn,omitempty"`
}
