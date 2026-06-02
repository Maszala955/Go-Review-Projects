package main

import "fmt"

func main() {
	Greet("Gladys")
	person := Person{
		Name: "Joe Lap",
		Age:  40,
	}
	fmt.Println(person.Name)
	fmt.Println(person.Age)
	person.Describe()
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

/* Exercice 3 — Les méthodes

En Go, on peut attacher une fonction directement à un struct. On appelle ça une méthode.
Crée une méthode Describe sur Person qui affiche "Je m'appelle <Name> et j'ai <Age> ans". Appelle-la depuis main.
> Indice : la syntaxe ressemble à une fonction normale, mais avec quelque chose de plus entre func et le nom...
*/
