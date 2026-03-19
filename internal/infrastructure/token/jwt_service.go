package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService struct {
	secretKey []byte
	expMin    int
}

func NewJWTService(secret string, expMin int) *JWTService {
	return &JWTService{
		secretKey: []byte(secret),
		expMin:    expMin,
	}
}

func (s *JWTService) Generate(userID int64, login string, userType string) (string, error) {
	claims := jwt.MapClaims{
		"sub":       userID,
		"login":     login,
		"user_type": userType,
		"exp":       time.Now().Add(time.Duration(s.expMin) * time.Minute).Unix(),
		"iat":       time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}
