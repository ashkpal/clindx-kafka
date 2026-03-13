package models

import (
	"time"

	"github.com/google/uuid"
)

type TestOrderUpdateEvent struct {
	EventVersion string    `json:"event_version"`
	EventType    string    `json:"event_type"`
	OccurredAt   time.Time `json:"occurred_at"`

	TestOrderID              uint       `json:"next_portal_order_id"`
	PatientID                uint       `json:"next_portal_patient_id"`
	LIMSAccessionID          string     `json:"accession_id"`
	LIMSAccessionTestOrderID uuid.UUID  `json:"portal_test_order_id"`
	LIMSTestRequisitionID    uuid.UUID  `json:"test_order_form_id"`
	Status                   string     `json:"current_status"`
	Diagnosis                string     `json:"diagnosis"`
	DiagnosisOther           string     `json:"diagnosis_other"`
	BloodCollectionDate      *time.Time `json:"plasma_collection_date"`
	TumorCollectionDate      *time.Time `json:"tumor_collection_date"`
}
