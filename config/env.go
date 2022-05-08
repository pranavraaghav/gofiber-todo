package config

import (
	"fmt"
	"os"
)

var (
	DB       = GetEnv("DB", "gotodo.db")
	PORT     = GetEnv("PORT", "5000")
	TOKENKEY = GetEnv("TOKEN_KEY", "secret")
	// TOKENEXP returns the jwt token expiration duration.
	// Should be time.ParseDuration string. Source: https://golang.org/pkg/time/#ParseDuration
	// default: 10h
	TOKENEXP = GetEnv("TOKEN_EXP", "10h")
)

func GetEnv(name string, fallback string) string {
	value, exists := os.LookupEnv(name)
	if exists {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}
