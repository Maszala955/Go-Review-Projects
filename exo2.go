package main

import (
	"errors"
	"fmt"
)


type Product struct {
	Name 	string
	Price float64
	Stock int
}

func (p Product) Display() {
	fmt.Printf("%s - %.2f€ - (%d en stock) \n", p.Name, p.Price, p.Stock)
}

type Displayable interface {
	Display()
}

func AfficherProduit(d Displayable) {
	d.Display()
}

func NouveauProduit(name string, price float64, stock int) (Product, error) {
	if (price < 0 || stock < 0) {
		return Product{}, errors.New("Error price or stock equals to 0")
	} else {
		product := Product{Name: name, Price: price, Stock: stock}
		return product, nil
	}
}

/*
Crée un struct Product avec Name (string), Price (float64) et Stock (int)
Ajoute une méthode Display() qui affiche "<Name> - <Price>€ (<Stock> en stock)"
Crée une interface Displayable avec une méthode Display()
Crée une fonction AfficherProduit(d Displayable) qui appelle d.Display()
Crée une fonction NouveauProduit(name string, price float64, stock int) (Product, error) qui retourne une erreur si price est négatif ou si stock est négatif
Dans main, crée 3 produits — un valide, un avec prix négatif, un avec stock négatif — et gère les erreurs
*/