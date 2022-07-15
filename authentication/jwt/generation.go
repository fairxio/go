package jwt

import "github.com/golang-jwt/jwt"

func SignJWT(token *jwt.Token, signingKey string) string {
	signedToken, _ := token.SignedString([]byte(signingKey))
	return signedToken
}

func GenerateJWT(target string, issuer string) *jwt.Token {

	claims := CreateClaims(target, issuer)
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

}
