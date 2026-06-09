package main

import (
	"fmt"

	"main/handlers"
	"main/repository"
)

func main() {
	conn, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
	fmt.Println("Connecté à PostgreSQL")

	err = repository.CreateTable(conn)
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
	

	handlers.StartServer(conn)
}
