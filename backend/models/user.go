package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"` // Do not expose password hash
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
}

// UserRegistration is used for binding registration data
type UserRegistration struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLogin is used for binding login data
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
