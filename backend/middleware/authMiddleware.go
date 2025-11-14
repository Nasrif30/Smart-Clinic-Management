package middleware

import (
	"context"
	"net/http"
	"smartclinic/backend/utils"
	"strings"
)

// UserContextKey is a custom type for context key to avoid collisions
type UserContextKey string

// AuthMiddleware verifies the JWT token
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authorization header required")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token format")
			return
		}

		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			utils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
			return
		}

		// Add user info to context for downstream handlers
		ctx := context.WithValue(r.Context(), UserContextKey("id"), claims.UserID)
		ctx = context.WithValue(ctx, UserContextKey("role"), claims.Role)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminMiddleware checks if the user has an 'admin' role
// This middleware MUST run AFTER AuthMiddleware
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(UserContextKey("role")).(string)
		
		if !ok || role != "admin" {
			utils.RespondWithError(w, http.StatusForbidden, "Admin access required")
			return
		}

		next.ServeHTTP(w, r)
	})
}
