package main

import (
	"fmt"
	
)


func main() {
	conn, err := ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
	fmt.Println("Connecté à PostgreSQL")

	err = CreateTable(conn)
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
	
	user := User{Name: "Test", Email: "Test@email.com", Age: 40}
	err = InsertUser(conn, user)
	if err != nil {
		fmt.Println("Erreur d'insert", err)
	}
	// StartServer()
}