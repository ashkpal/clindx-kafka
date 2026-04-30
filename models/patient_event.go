package models

import (
	"time"

	"github.com/google/uuid"
)

type PatientEvent struct {
	EventVersion string `json:"event_version"`
	EventType    string `json:"event_type"`

	PatientID         uint       `json:"patient_id"`
	FirstName         string     `json:"first_name"`
	LastName          string     `json:"last_name"`
	MiddleInitial     string     `json:"middle_initial"`
	DOB               time.Time  `json:"dob"`
	Gender            string     `json:"gender"`
	TestOrderCadence  int        `json:"test_order_cadence"`
	CadenceEndDate    *time.Time `json:"cadence_end_date"`
	CadenceCollection string     `json:"cadence_collection"`
	SurgeryDate       *time.Time `json:"surgery_date"`
	PhoneCountryCode  string     `json:"phone_country_code"`
	Phone             string     `json:"phone"`
	Email             string     `json:"email"`
	Street            string     `json:"street"`
	City              string     `json:"city"`
	State             string     `json:"state"`
	Zip               string     `json:"zip"`
	Country           string     `json:"country"`
	LIMSPatientID     uuid.UUID  `json:"lims_patient_id"`

	PatientPractices []PatientPractice `json:"patient_practices"`
}

type PatientPractice struct {
	PatientID        uint   `json:"patient_id"`
	PracticeID       uint   `json:"practice_id"`
	MRN              string `json:"mrn"`
	RelationshipType string `json:"relationship_type"`
}
