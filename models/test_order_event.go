package models

import "time"

type TestOrderEvent struct {
	EventVersion string    `json:"event_version"`
	EventType    string    `json:"event_type"`
	OccurredAt   time.Time `json:"occurred_at"`

	OrderID              uint       `json:"order_id"`
	TRFNum               string     `json:"trf_num,omitempty"`
	PatientID            uint       `json:"patient_id"`
	PracticeID           uint       `json:"practice_id"`
	PhysicianID          uint       `json:"physician_id"`
	TestName             string     `json:"test_name"`
	TestType             string     `json:"test_type"`
	EmailPhysicianFlag   bool       `json:"email_physician_flag"`
	FaxPhysicianFlag     bool       `json:"fax_physician_flag"`
	PortalPhysicianFlag  bool       `json:"portal_physician_flag"`
	CopyRecipientName    string     `json:"copy_recipient_name"`
	CopyPracticeName     string     `json:"copy_practice_name"`
	CopyPracticeEmail    string     `json:"copy_practice_email"`
	CopyPracticeFax      string     `json:"copy_practice_fax"`
	OrderDate            time.Time  `json:"order_date"`
	Diagnosis            string     `json:"diagnosis"`
	DiagnosisOther       string     `json:"diagnosis_other"`
	Stage                string     `json:"stage"`
	StageSubtype         string     `json:"stage_subtype"`
	ICD10Codes           string     `json:"icd10_codes"`
	ActiveDiseaseFlag    bool       `json:"active_disease_flag"`
	ImmunoTxName         string     `json:"immuno_tx_name"`
	Treatments           string     `json:"treatments"`
	CancerHistory        bool       `json:"cancer_history"`
	SurgeryDate          *time.Time `json:"surgery_date"`
	TumorRequest         string     `json:"tumor_request"`
	TumorSpecimenID      string     `json:"tumor_specimen_id"`
	TumorCollectionDate  *time.Time `json:"tumor_collection_date"`
	TumorExhaustFlag     bool       `json:"tumor_exhaust_flag"`
	TumorHandling        string     `json:"tumor_handling"`
	TumorBiopsySite      string     `json:"tumor_biopsy_site"`
	TumorBiopsySiteOther string     `json:"tumor_biopsy_site_other"`
	PathName             string     `json:"path_name"`
	PathInstName         string     `json:"path_inst_name"`
	PathInstCity         string     `json:"path_inst_city"`
	PathInstEmail        string     `json:"path_inst_email"`
	PathInstFax          string     `json:"path_inst_fax"`
	PathInstPhone        string     `json:"path_inst_phone"`
	BloodSpecimenID      string     `json:"blood_specimen_id"`
	BloodPreBiopsyFlag   bool       `json:"blood_prebiopsy_flag"`
	BloodHandling        string     `json:"blood_collection"`
	BloodCollectionDate  *time.Time `json:"blood_collection_date"`
	DiseaseStatuses      string     `json:"disease_statuses"`
	InsuranceType        string     `json:"insurance_type"`
	PatientDischargeDate *time.Time `json:"patient_discharge_date"`
	InsuranceProvider    string     `json:"insurance_provider"`
	InsurancePolicyNum   string     `json:"insurance_policynum"`
	InsuredName          string     `json:"insured_name"`
	InsuredDOB           *time.Time `json:"insured_dob"`
	InsuredRelationship  string     `json:"insured_relationship"`
	PriorAuthNum         string     `json:"prior_authnum"`
	Consent              bool       `json:"consent"`
	ConsentSignedBy      string     `json:"consent_signed_by"`
	ConsentSignedDate    *time.Time `json:"consent_signed_date"`
	TRFCreatedBy         string     `json:"trf_created_by"`
	FileName             string     `json:"pdf_filename"`
	MedFiles             []MedFile  `json:"medfiles"`
}

type MedFile struct {
	Name       string    `json:"file_name"`
	FileType   string    `json:"file_type"`
	FilePath   string    `json:"file_path"`
	MimeType   string    `json:"mime_type"`
	UploadedBy string    `json:"uploaded_by"`
	UploadDate time.Time `json:"upload_date"`
}
