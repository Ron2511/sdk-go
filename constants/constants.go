package constants

import (
	"fmt"

	"github.com/Ron2511/sdk-go/helpers"
)

var env = helpers.GetEnv()

var (
	AUTH_URL, BASE_URL = getURLS(env)
)

func getURLS(env string) (string, string) {
	if env != "" {
		return fmt.Sprintf("https://auth.%s.ua.la", env), fmt.Sprintf("https://checkout.%s.ua.la", env)
	}
	return "https://auth.ua.la", "https://checkout.ua.la"
}
