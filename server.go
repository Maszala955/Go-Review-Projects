package main

import (
	"net/http"
	"io"
	"log"
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
	
	http.HandleFunc("/hello", h1)
	http.HandleFunc("/", h2)
	
	log.Println("Server listen port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}