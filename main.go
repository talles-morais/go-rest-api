package main

import (
	"fmt"

	"github.com/talles-morais/go-rest-api/database"
	"github.com/talles-morais/go-rest-api/routes"
)

func main() {
	database.ConnectDB()
	fmt.Println("Iniciando o servidor")
	routes.HandleRequest()
}
