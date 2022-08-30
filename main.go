package main

import (
	"fmt"

	"github.com/Ron2511/sdk-go/client"
)

func main() {
	//token := "" //se va a borrar
	credentials := `{
		"user_name": "new_user_1659988043",
		"client_id": "01g28sVBNwGJDKz8QX2hRyH3mJQDyymk",
		"client_secret_id": "y--SbJipBCbbcptm8u-vD3chIApD7thl0gC87TD5GH-8YCzXUsDFHCiQhMKNvw7h",
		"grant_type": "client_credentials"
	}`
	cli := client.New(nil)
	cli.Auth(credentials)
	order := cli.GetOrder("052f288c-5c37-4760-971c-3613b140fde9")
	fmt.Println(order)
}
