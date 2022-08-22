package client

import (
	"flag"
	"fmt"
)

func main() {

	fmt.Println("Esto es nuestro SDK en GO")

	token := flag.String("token", "", "Token para usar la api")
	flag.Parse()
	// Imprimirlo
	fmt.Println("El token es ", *token)

}
