package main

import (
	"fmt"

	"github.com/talles-morais/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
