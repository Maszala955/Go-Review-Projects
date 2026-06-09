package repository

import (
	"context"
	"main/models"

	"github.com/jackc/pgx/v5"
)

func CreateTable(conn *pgx.Conn) error {
	query := `CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    age INT NOT NULL)
	`
	_, err := conn.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func InsertUser(conn *pgx.Conn, user models.User) error {
	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3)`
	_, err := conn.Exec(context.Background(), query, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func GetUsers(conn *pgx.Conn) ([]models.User, error) {
	query := `SELECT name, email, age FROM users`
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.Name, &u.Email, &u.Age)
		users = append(users, u)
	}
	return users, nil
}

func DeleteUser(conn *pgx.Conn, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}