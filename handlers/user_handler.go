package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"main/models"
	"main/repository"

	"github.com/jackc/pgx/v5"
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

func StartServer(conn *pgx.Conn) {
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
		user, _ := models.NewUser(name, email, age)
		w.Header().Set("Content-Type", "application/json")

		jsonStr, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(jsonStr)

	}
	h4 := func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = repository.InsertUser(conn, user)
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
	h5 := func(w http.ResponseWriter, r *http.Request) {
		users, err := repository.GetUsers(conn)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		userStr, err := json.Marshal(users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(userStr)

	}
	http.HandleFunc("/", h2)
	http.HandleFunc("/hello", h1)
	http.HandleFunc("/newUser", h3)
	http.HandleFunc("/user", h4)
	http.HandleFunc("/users", h5)

	log.Println("Server listen port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
