package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"smartclinic/backend/config"
	"smartclinic/backend/db"
	"smartclinic/backend/routes"

	"[github.com/gorilla/mux](https://github.com/gorilla/mux)"
	_ "[github.com/lib/pq](https://github.com/lib/pq)"
)

func main() {
	// 1. Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load configuration: %v", err)
	}

	// 2. Initialize Database Connection
	connStr := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	database, err := db.InitDB(connStr)
	if err != nil {
		log.Fatalf("Could not initialize database: %v", err)
	}
	defer database.Close()

	// 3. Create tables
	if err := db.CreateTables(database); err != nil {
		log.Fatalf("Could not create tables: %v", err)
	}

	// 4. Initialize Router
	r := mux.NewRouter()

	// 5. Register Routes
	registerRoutes(r, database)

	// 6. Start Server
	log.Printf("Server starting on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}

// registerRoutes sets up all API routes
func registerRoutes(r *mux.Router, database *sql.DB) {
	// Create a subrouter for /api
	api := r.PathPrefix("/api").Subrouter()

	// Pass the db connection to the route initializers
	routes.RegisterUserRoutes(api, database)
	routes.RegisterAppointmentRoutes(api, database)
	routes.RegisterAdminRoutes(api, database)

	// Health check route
	api.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Server is healthy"))
	}).Methods("GET")
}
