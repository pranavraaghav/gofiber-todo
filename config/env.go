package config

import (
	"fmt"
	"os"
)

var (
	DB = GetEnv("DB", "gotodo.db")
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
