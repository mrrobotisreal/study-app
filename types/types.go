package types

import "github.com/golang-jwt/jwt/v4"

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}
