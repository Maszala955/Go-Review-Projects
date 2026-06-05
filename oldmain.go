package main

import (
	"errors"
	"fmt"
)

func oldmain() {
	// Greet("Gladys")
	// person := Person{
	// 	Name: "Joe Lap",
	// 	Age:  40,
	// }
	// fmt.Println(person.Name)
	// fmt.Println(person.Age)
	// person.Describe()

	// Présenter(Animal{"Chien"})
	// Présenter(Person{Name: "Joe", Age: 33})
	// result, err := Diviser(2, 3)
	// if err != nil {
	// 	fmt.Printf("Une erreur s'est produite %v", err)
	// } else {
	// 	fmt.Printf("Le résultat est %f\n", result)
	// }

	// result2, err := Diviser(2, 0)
	// if err != nil {
	// 	fmt.Printf("%v\n", err)
	// } else {
	// 	fmt.Printf("Le résultat est %f\n", result2)
	// }
	//
	// Exo1
	// book1, err := NouveauLivre("1984", "George Orwell", 400)
	// if err != nil {
	// 	fmt.Print("Il n'ya pas de pages", err)
	// } else {
	// 	Afficher(book1)
	// }
	// book2, err := NouveauLivre("test", "Neil Armstrong", 0)
	// if err != nil {
	// 	fmt.Printf("Il n'ya pas de pages\n")
	// } else {
	// 	Afficher(book2)
	// }
	//
	// Exo 2
	// product1, err := NouveauProduit("Iphone", 980, 150)
	// if err != nil {
	// 	fmt.Printf("Prix inférieur à 0 ou pas de stock")
	// } else {
	// 	AfficherProduit(product1)
	// }
	// product2, err := NouveauProduit("IPad", -1 , 1)
	// if err != nil {
	// 	fmt.Printf("Prix inférieur à 0 \n")
	// } else {
	// 	AfficherProduit(product2)
	// }
	// product3, err := NouveauProduit("Télé", 1230, -1)
	// if err != nil {
	// 	fmt.Printf("Pas de stock\n")
	// } else {
	// 	AfficherProduit(product3)
	// }
	// Exo3
	// student1, err := NouvelEtudiant("Alexis", 80, 90.44)
	// if err != nil {
	// 	fmt.Printf("Errors")
	// } else {
	// 	AfficherEtudiant(student1)
	// }
	// student2, err := NouvelEtudiant("Marie", 50, -90.44)
	// if err != nil {
	// 	fmt.Printf("Grade inférieur à 0 \n")
	// } else {
	// 	AfficherEtudiant(student2)
	// }
	// student3, err := NouvelEtudiant("Alexis", 80, 190.44)
	// if err != nil {
	// 	fmt.Printf("Score inférieur à 0\n")
	// } else {
	// 	AfficherEtudiant(student3)
	// }
	// Go routine
	// go Compter("Alex", 100)
	// go Compter("Marie", 100)
	// time.Sleep(time.Second) // Permet de faire de la concurrence
	//
	// Channels
	ch := make(chan int)
	go Calculer(5, 2, ch)
	ch1 := <-ch
	fmt.Print(ch1, "\n")
}

func Greet(name string) {
	fmt.Printf("Welcome %s \n", name)

}

func (p Person) Describe() {
	fmt.Printf("Je m'appelle %s et j'ai %d ans \n", p.Name, p.Age)
}

type Person struct {
	Name string
	Age  int
}

type Describer interface {
	Describe()
}

type Animal struct {
	Species string
}

func (a Animal) Describe() {
	fmt.Printf("Je suis un %s\n", a.Species)
}

func Présenter(d Describer) {
	d.Describe()
}

func Diviser(a, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Inerdiction de diviser par 0")
	}
	result := a / b
	return result, nil
}

// go routine
func Compter(nom string, n int) {
	for i := range n {
		fmt.Printf("%s : %d \n", nom, i)
	}
}

// Channels
func Calculer(a, b int, ch chan int) {
	result := a + b
	ch <- result
}
