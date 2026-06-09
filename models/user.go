package models

import "errors"

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func NewUser(name, email string, age int) (User, error) {
	if age < 10 {
		return User{}, errors.New("student too young")
	} else {
		return User{Name: name, Email: email, Age: age}, nil
	}
}
