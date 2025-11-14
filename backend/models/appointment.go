package models

import (
	"time"
)

// Appointment represents an appointment in the system
type Appointment struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	UserName  string    `json:"userName"`
	Doctor    string    `json:"doctor"`
	Date      string    `json:"date"` // Using string for YYYY-MM-DD
	Time      string    `json:"time"` // Using string for HH:MM
	Reason    string    `json:"reason"`
	Status    string    `json:"status"` // e.g., 'pending', 'confirmed', 'cancelled'
	CreatedAt time.Time `json:"createdAt"`
}

// CreateAppointmentRequest is used for binding new appointment data
type CreateAppointmentRequest struct {
	Doctor string `json:"doctor"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	Reason string `json:"reason"`
}