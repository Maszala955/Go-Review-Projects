package main

import (
	"main/handlers"
	"main/repository"
)

func main() {
	conn, err := repository.Connect()
	repository.CreateDBTable(err, conn)
	handlers.StartServer(conn)
}
