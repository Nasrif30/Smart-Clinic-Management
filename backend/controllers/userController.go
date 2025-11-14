package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"smartclinic/backend/middleware"
	"smartclinic/backend/models"
	"smartclinic/backend/utils"
)

// UserController holds database connection
type UserController struct {
	DB *sql.DB
}

// NewUserController creates a new user controller
func NewUserController(db *sql.DB) *UserController {
	return &UserController{DB: db}
}

// RegisterUser handles user registration
func (c *UserController) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var regData models.UserRegistration
	if err := json.NewDecoder(r.Body).Decode(&regData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if regData.Name == "" || regData.Email == "" || regData.Password == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "All fields (name, email, password) are required")
		return
	}

	// Check if user already exists
	var exists bool
	err := c.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)", regData.Email).Scan(&exists)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}
	if exists {
		utils.RespondWithError(w, http.StatusConflict, "Email already in use")
		return
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(regData.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password")
		return
	}

	// Insert new user
	var newUser models.User
	role := "user" // Default role
	err = c.DB.QueryRow(
		"INSERT INTO users (name, email, password_hash, role) VALUES ($1, $2, $3, $4) RETURNING id, name, email, role, created_at",
		regData.Name, regData.Email, hashedPassword, role,
	).Scan(&newUser.ID, &newUser.Name, &newUser.Email, &newUser.Role, &newUser.CreatedAt)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, newUser)
}

// LoginUser handles user login
func (c *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginData models.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	var user struct {
		ID           string
		Email        string
		PasswordHash string
		Role         string
	}

	err := c.DB.QueryRow("SELECT id, email, password_hash, role FROM users WHERE email = $1", loginData.Email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}

	// Check password
	if !utils.CheckPasswordHash(loginData.Password, user.PasswordHash) {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// GetUserProfile handles fetching user profile
func (c *UserController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	// Get user ID from context (set by AuthMiddleware)
	userID, ok := r.Context().Value(middleware.UserContextKey("id")).(string)
	if !ok {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	var user models.User
	err := c.DB.QueryRow(
		"SELECT id, name, email, role, created_at FROM users WHERE id = $1",
		userID,
	).Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			return
		}
		utils.RespondWithError(w, http.StatusInternalServerError, "Database error")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, user)
}
