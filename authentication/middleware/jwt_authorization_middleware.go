package middleware

import (
	"fmt"
	jwt2 "github.com/fairxio/go/authentication/jwt"
	"github.com/golang-jwt/jwt"
	"net/http"
)

var secretKey string = ""

//check whether user is authorized or not
func IsJWTAuthorized(handler http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Should be in the Authorization: header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Parse out the token
		jwtString := authHeader[8:] // Bearer:
		token, _ := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {

			// Verify Algorithm
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Return secret key
			return []byte(secretKey), nil
		})

		if claims, ok := token.Claims.(jwt2.Claims); ok && token.Valid {
			r.Header.Set("subject", claims.Subject)
			handler.ServeHTTP(w, r)
			return
		}

		// Not authorized
		w.WriteHeader(http.StatusUnauthorized)
		return

	}

}
