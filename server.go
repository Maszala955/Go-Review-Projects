package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strconv"
)

// Deux concept clés :
// 1. handler - une fonction reçoit une requête et écrit une réponse
// func(w http.ResponseWriter, r *http.Request)
// - `w` : là où tu écris ta réponse
// - `r` : la requête entrante (méthode, URL, headers...)

// 2. http.ListenAndServe - démarre le serveur sur un port
// Exercice :
/*
Dans un nouveau fichier server.go :
Crée un handler pour la route `/hello` qui répond `"Hello, World!"`
Crée un handler pour la route `/` qui répond `"Bienvenue sur mon serveur !"`
Démarre le serveur sur le port `8080`
*/

func StartServer() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello World")
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Bienvenue sur mon serveur")
	}
	h3 := func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		email := r.URL.Query().Get("email")
		ageStr := r.URL.Query().Get("age")
		age, _ := strconv.Atoi(ageStr)
		user, _ := NewUser(name, email, age)
		w.Header().Set("Content-Type", "application/json")

		jsonStr, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonStr)

	}
	h4 := func(w http.ResponseWriter, r *http.Request) {
		var user User
		// décode le body JSON dans user

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// retourne user en JSON comme tu sais déjà faire
		userStr, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(userStr)
	}

	http.HandleFunc("/hello", h1)
	http.HandleFunc("/", h2)
	http.HandleFunc("/newUser", h3)
	http.HandleFunc("/user", h4)

	log.Println("Server listen port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

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
