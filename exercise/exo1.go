package main

import (
	"errors"
	"fmt"
)

type Book struct {
	Title  string
	Author string
	Pages  int
}

type Summarize interface {
	Summarize()
}

func (b Book) Summarize() {
	fmt.Printf("%s par %s - %d pages\n", b.Title, b.Author, b.Pages)
}

func Afficher(s Summarize) {
	s.Summarize()
}

func NouveauLivre(title, author string, pages int) (Book, error) {
	if pages == 0 {
		return Book{}, errors.New("Errors")
	}
	book := Book{Title: title, Author: author, Pages: pages}
	return book, nil
}

/**
Crée un struct Book avec les champs Title (string), Author (string) et Pages (int)
Ajoute une méthode Summary() qui affiche "<Title> par <Author> - <Pages> pages"
Crée une interface Summarizable avec une méthode Summary()
Crée une fonction Afficher(s Summarizable) qui appelle s.Summary()
Crée une fonction NouveauLivre(title, author string, pages int) (*Book, error) qui retourne une erreur si pages vaut 0 ou moins
Dans main, crée deux livres — un valide, un avec pages = 0 — et gère les deux cas
**/