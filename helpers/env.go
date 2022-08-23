package helpers

import (
	"os"
	"strings"
)

func GetEnv() string {
	env := os.Getenv("env")
	if env != "" {
		return strings.ToLower(env)
	}
	// production
	return env
}
