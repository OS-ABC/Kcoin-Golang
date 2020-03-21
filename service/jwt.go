package service

import (
	"Kcoin-Golang/conf"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(conf.Config.Jwt.JwtSecret)

type Claims struct {
	UserID string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(userID string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(30 * time.Minute)

	claims := Claims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "Kcoin",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
