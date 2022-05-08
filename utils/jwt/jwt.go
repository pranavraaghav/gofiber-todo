package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pranavraagz/gofiber-todo/config"
)

type TokenPayload struct {
	ID uint
}

func Generate(payload *TokenPayload) string {
	v, err := time.ParseDuration(config.TOKENEXP)

	if err != nil {
		panic("Invalid time duration. Should be a time.ParseDuration string")
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(v).Unix(),
		"ID":  payload.ID,
	})

	token, err := t.SignedString([]byte(config.TOKENKEY))

	if err != nil {
		panic(err)
	}

	return token

}
