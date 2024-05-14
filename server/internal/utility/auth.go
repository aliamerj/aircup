package utility

import (
	"github.com/golang-jwt/jwt/v5"
)

type JwtAuthClaims struct {
	AccountID uint   `json:"account_id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	jwt.RegisteredClaims
}
