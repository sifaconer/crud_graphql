package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ApiClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

var secret string = "secret-api"

func GenerateJWT(userID int) (string, error) {
	claims := ApiClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			Audience:  "Api",
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Id:        "",
			IssuedAt:  0,
			Issuer:    "Graphql Api CRUD",
			NotBefore: 0,
			Subject:   "User",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func ValidJWT(tokenJWT string) error {
	if len(tokenJWT) == 0 {
		return errors.New("sin acceso por falta de token")
	}

	_, err := jwt.Parse(tokenJWT, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token no valido")
		}

		return []byte(secret), nil
	})

	return err
}
