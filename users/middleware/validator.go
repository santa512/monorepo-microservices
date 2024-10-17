package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// ValidateToken is a middleware that validate token from header
func ValidateToken(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Get the token from the Authorization header
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization header"))
			return
		}

		// Strip the Bearer prefix from the token string, if present
		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid Authorization header format"))
			return
		}

		tokenString = tokenParts[1]

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the algorithm, ensure it's what you expect (e.g., HMAC SHA256)
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key used for signing
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}

		// Token is valid, proceed to the next handler
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
