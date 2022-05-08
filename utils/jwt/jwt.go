package jwt

import (
	"errors"
	"fmt"
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

func parse(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.TOKENKEY), nil
	})
}

func Verify(token string) (*TokenPayload, error) {
	parsed, err := parse(token)

	if err != nil {
		return nil, err
	}

	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	id, ok := claims["ID"].(float64)
	if !ok {
		return nil, errors.New("Something went wrong")
	}
	return &TokenPayload{
		ID: uint(id),
	}, nil
}
