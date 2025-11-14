package routes

import (
	"database/sql"
	"smartclinic/backend/controllers"
	"smartclinic/backend/middleware"

	"[github.com/gorilla/mux](https://github.com/gorilla/mux)"
)

// RegisterAppointmentRoutes sets up protected routes for appointments
func RegisterAppointmentRoutes(r *mux.Router, db *sql.DB) {
	c := controllers.NewAppointmentController(db)

	// All appointment routes are protected
	apptRouter := r.PathPrefix("/appointments").Subrouter()
	apptRouter.Use(middleware.AuthMiddleware)

	apptRouter.HandleFunc("", c.CreateAppointment).Methods("POST")
	apptRouter.HandleFunc("/my", c.GetMyAppointments).Methods("GET")
}