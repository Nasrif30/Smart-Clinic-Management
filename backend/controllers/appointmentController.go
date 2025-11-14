package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smartclinic/backend/middleware"
	"smartclinic/backend/models"
	"smartclinic/backend/utils"
)

// AppointmentController holds database connection
type AppointmentController struct {
	DB *sql.DB
}

// NewAppointmentController creates a new appointment controller
func NewAppointmentController(db *sql.DB) *AppointmentController {
	return &AppointmentController{DB: db}
}

// CreateAppointment handles booking a new appointment
func (c *AppointmentController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	// Get user ID and role from context
	userID, ok := r.Context().Value(middleware.UserContextKey("id")).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	// In this system, only 'user' role can book.
	role, _ := r.Context().Value(middleware.UserContextKey("role")).(string)
	if role != "user" {
		utils.RespondWithError(w, http.StatusForbidden, "Only users can book appointments")
		return
	}
	
	var req models.CreateAppointmentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	// Get user's name
	var userName string
	err := c.DB.QueryRow("SELECT name FROM users WHERE id = $1", userID).Scan(&userName)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to find user")
		return
	}

	var newAppt models.Appointment
	err = c.DB.QueryRow(
		"INSERT INTO appointments (user_id, user_name, doctor, date, time, reason) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, user_id, user_name, doctor, date, time, reason, status, created_at",
		userID, userName, req.Doctor, req.Date, req.Time, req.Reason,
	).Scan(&newAppt.ID, &newAppt.UserID, &newAppt.UserName, &newAppt.Doctor, &newAppt.Date, &newAppt.Time, &newAppt.Reason, &newAppt.Status, &newAppt.CreatedAt)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create appointment")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, newAppt)
}

// GetMyAppointments fetches appointments for the logged-in user
func (c *AppointmentController) GetMyAppointments(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserContextKey("id")).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}
	
	role, _ := r.Context().Value(middleware.UserContextKey("role")).(string)
	if role != "user" {
		utils.RespondWithError(w, http.StatusForbidden, "Only users can view their appointments")
		return
	}
	
	rows, err := c.DB.Query(
		"SELECT id, user_id, user_name, doctor, date, time, reason, status, created_at FROM appointments WHERE user_id = $1",
		userID,
	)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	defer rows.Close()

	var appointments []models.Appointment
	for rows.Next() {
		var appt models.Appointment
		if err := rows.Scan(&appt.ID, &appt.UserID, &appt.UserName, &appt.Doctor, &appt.Date, &appt.Time, &appt.Reason, &appt.Status, &appt.CreatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan appointment")
			return
		}
		appointments = append(appointments, appt)
	}

	utils.RespondWithJSON(w, http.StatusOK, appointments)
}