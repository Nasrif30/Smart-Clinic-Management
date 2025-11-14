package routes

import (
	"database/sql"
	"smartclinic/backend/controllers"
	"smartclinic/backend/middleware"

	"[github.com/gorilla/mux](https://github.com/gorilla/mux)"
)

// RegisterAdminRoutes sets up protected routes for admin actions
func RegisterAdminRoutes(r *mux.Router, db *sql.DB) {
	c := controllers.NewAdminController(db)

	// All admin routes are protected by Auth AND Admin middleware
	adminRouter := r.PathPrefix("/admin").Subrouter()
	adminRouter.Use(middleware.AuthMiddleware, middleware.AdminMiddleware)

	adminRouter.HandleFunc("/users", c.GetAllUsers).Methods("GET")
	adminRouter.HandleFunc("/appointments", c.GetAllAppointments).Methods("GET")
	adminRouter.HandleFunc("/appointments/{id}/status", c.UpdateAppointmentStatus).Methods("PATCH")
}