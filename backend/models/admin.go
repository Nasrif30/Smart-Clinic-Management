package models

// ReportData defines a structure for aggregated report data
type ReportData struct {
	TotalUsers         int `json:"totalUsers"`
	TotalAppointments  int `json:"totalAppointments"`
	PendingAppointments int `json:"pendingAppointments"`
}

// (You can add more admin-specific data models here)
