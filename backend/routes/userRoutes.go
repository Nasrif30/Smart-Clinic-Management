package routes

import (
	"database/sql"
	"smartclinic/backend/controllers"
	"smartclinic/backend/middleware"

	"[github.com/gorilla/mux](https://github.com/gorilla/mux)"
)

// RegisterUserRoutes sets up routes for user authentication and profile
func RegisterUserRoutes(r *mux.Router, db *sql.DB) {
	c := controllers.NewUserController(db)

	// Public routes
	r.HandleFunc("/auth/register", c.RegisterUser).Methods("POST")
	r.HandleFunc("/auth/login", c.LoginUser).Methods("POST")

	// Protected route
	// We apply middleware directly here
	profileRouter := r.PathPrefix("/user/profile").Subrouter()
	profileRouter.Use(middleware.AuthMiddleware)
	profileRouter.HandleFunc("", c.GetUserProfile).Methods("GET")
}
