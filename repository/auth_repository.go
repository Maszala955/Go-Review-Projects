package repository

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
)

func Login(conn *pgx.Conn, email, password string) (string, error) {
	query := `SELECT id, password FROM users WHERE email = $1`
	rows := conn.QueryRow(context.Background(), query, email)

	var id int
	var dbPassword string
	err := rows.Scan(&id, &dbPassword)
	if err != nil {
		return "", err
	}

	if dbPassword != password {
		return "", errors.New("Erreur lors de l'auth")
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  id,
			"exp": time.Now().Add(24 * time.Hour).Unix(),
		})
		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			return "", err
		}
		return tokenString, nil
	}
}
