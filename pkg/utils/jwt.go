package util

import (
	"errors"
	"os"
	"strings"
	"time"
	"todo-list/pkg/e"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct {
	Id        uint   `json:"id"`
	Username  string `json:"username"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

func GenerateToken(id uint, username string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id:        id,
		Username:  username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "to-do-list",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(tokenHeader string) (*Claims, error) {
	// Bearer xxx
	parts := strings.Split(tokenHeader, " ")
	if len(parts) != 2 {
		return nil, errors.New(e.GetMsg(e.ErrorAuthCheckTokenFail))
	}
	parts[0] = strings.TrimSpace(parts[0])
	if parts[0] != "Bearer" {
		return nil, errors.New(e.GetMsg(e.ErrorAuthCheckTokenFail))
	}

	tokenClaims, err := jwt.ParseWithClaims(parts[1], &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
