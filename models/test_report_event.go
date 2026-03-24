package models

import "time"

type TestReportEvent struct {
	TestOrderID          uint       `json:"TestOrderID"`
	AccessionID          string     `json:"AccessionID"`
	ReportDate           *time.Time `json:"ReportDate"`
	PlasmaCollectionDate *time.Time `json:"PlasmaCollectionDate"`
	Notes                string     `json:"Notes"`
	DxResult             string     `json:"DxResult"`
	Result               string     `json:"Result"`
	CtDNAFraction        float32    `json:"CtDNAFraction"`
	Status               string     `json:"Status"`
	ReportType           string     `json:"ReportType"`
	AmendReason          string     `json:"AmendReason"`
	FilePath             string     `json:"FilePath"`
	MimeType             string     `json:"MimeType"`
	CreatedBy            string     `json:"CreatedBy"`
}
