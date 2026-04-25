package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var JwtSecret = []byte("secret")

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(userID uint, username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "LarvAuth",
		},
	}

	//use HS256 to sign the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtSecret)
}

// ParseToken resolve token and return claims
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	//verify token and return claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
