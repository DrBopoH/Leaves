// source/auth/jwt.go
package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)



type Claims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}



func GenerateToken(userID int, username string, rememberMe bool) (string, time.Time, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	expirationTime := time.Now().Add(24 * time.Hour)
	if rememberMe {
		expirationTime = time.Now().Add(30 * 24 * time.Hour)
	}

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	
	return tokenString, expirationTime, err
}

func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}
	return claims, nil
}