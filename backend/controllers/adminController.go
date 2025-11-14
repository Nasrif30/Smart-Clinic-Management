package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smartclinic/backend/models"
	"smartclinic/backend/utils"

	"[github.com/gorilla/mux](https://github.com/gorilla/mux)"
)

// AdminController holds database connection
type AdminController struct {
	DB *sql.DB
}

// NewAdminController creates a new admin controller
func NewAdminController(db *sql.DB) *AdminController {
	return &AdminController{DB: db}
}

// GetAllUsers (Admin)
func (c *AdminController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, name, email, role, created_at FROM users")
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
			utils.RespondWithError(w, http.StatusInternalServerError, "Failed to scan user")
			return
		}
		users = append(users, user)
	}

	utils.RespondWithJSON(w, http.StatusOK, users)
}

// GetAllAppointments (Admin)
func (c *AdminController) GetAllAppointments(w http.ResponseWriter, r *http.Request) {
	rows, err := c.DB.Query("SELECT id, user_id, user_name, doctor, date, time, reason, status, created_at FROM appointments")
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

// UpdateAppointmentStatus (Admin)
func (c *AdminController) UpdateAppointmentStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	apptID := vars["id"]

	var payload struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if payload.Status != "pending" && payload.Status != "confirmed" && payload.Status != "cancelled" {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid status value")
		return
	}

	result, err := c.DB.Exec("UPDATE appointments SET status = $1 WHERE id = $2", payload.Status, apptID)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to update appointment")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to check update")
		return
	}
	if rowsAffected == 0 {
		utils.RespondWithError(w, http.StatusNotFound, "Appointment not found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "Appointment status updated successfully"})
}
