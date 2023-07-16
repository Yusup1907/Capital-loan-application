package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const KEY = "your-secret-key"

type JwtClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func GenerateToken(email string) (string, error) {
	now := time.Now().UTC()
	end := now.Add(24 * time.Hour)
	claims := &JwtClaims{
		Username: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: end.Unix(),
			IssuedAt:  now.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(KEY))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return signedToken, nil
}
func VerifyAccessToken(tokenString string) (string, error) {
	claim := &JwtClaims{}
	t, err := jwt.ParseWithClaims(tokenString, claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(KEY), nil
	})
	if err != nil {
		return "", fmt.Errorf("VerifyAccessToken: %w", err)
	}
	if !t.Valid {
		return "", fmt.Errorf("VerifyAccessToken: Invalid token")
	}
	return claim.Username, nil
}
