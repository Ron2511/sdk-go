package main

import (
	"fmt"

	"github.com/Ron2511/sdk-go/client"
)

func main() {
	token := ""
	cli := client.New(token)
	order := cli.GetOrder("052f288c-5c37-4760-971c-3613b140fde9")
	fmt.Println(order)
}
