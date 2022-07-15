package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	Roles []string `json:"roles"`
	jwt.StandardClaims
}

func CreateClaims(target string, issuer string) Claims {
	claims := Claims{
		Roles: []string{"STANDARD"},
	}

	timeNow := time.Now()

	claims.Id = target
	claims.Issuer = issuer
	claims.Subject = target
	claims.Audience = "auth,did,dwn"
	claims.NotBefore = timeNow.Unix()
	claims.IssuedAt = timeNow.Unix()
	claims.ExpiresAt = timeNow.Add(24 * time.Hour).Unix()

	return claims
}
