package repository

import (
	"fmt"
	"context"

	"github.com/jackc/pgx/v5"
)

func Connect() (*pgx.Conn, error) {
	conn, err := ConnectDB()
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
	fmt.Println("Connecté à PostgreSQL")
	return conn, err
}
func CreateDBTable(err error, conn *pgx.Conn) {
	err = CreateTable(conn)
	if err != nil {
		fmt.Println("Erreur de connexion", err)
	}
}
func ConnectDB() (*pgx.Conn, error) {

	urlDB := "postgres://gouser:password@localhost:5432/go_review"

	conn, err := pgx.Connect(context.Background(), urlDB)
	if err != nil {
    return nil, err
	}

	return conn, nil
}