package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

var UserKey = contextKey("user")



func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var tokenString string
		cookie, err := r.Cookie("auth_token")
		
		if err == nil {
			tokenString = cookie.Value
		} else {
			authHeader := r.Header.Get("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				tokenString = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}
		if tokenString == "" {
			http.Error(w, "Unauthorized request", http.StatusUnauthorized)
			return
		}

		// Parse and verify token
		secret := []byte(os.Getenv("JWT_SECRET"))
		parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return secret, nil
		})

		if err != nil || !parsedToken.Valid {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		// Extract claims
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid claims", http.StatusUnauthorized)
			return
		}

		// fmt.Println("claims is: ", claims)

		// Example: yahan DB se user fetch karna hoga (jaise tum Candidate.findById karte the)
		userID := claims["user_id"]
		email := claims["email"]

		// Dummy check: in real code DB me query karo
		if userID == nil || email == nil {
			http.Error(w, "Invalid user in token", http.StatusUnauthorized)
			return
		}

		// Candidate ko request context me daal do
		ctx := context.WithValue(r.Context(), UserKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}