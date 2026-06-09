package models

import "errors"

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
	Password string `json:"password"`
}

func NewUser(name, email string, age int) (User, error) {
	if age < 10 {
		return User{}, errors.New("student too young")
	} else {
		return User{Name: name, Email: email, Age: age}, nil
	}
}
