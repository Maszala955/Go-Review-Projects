package main

import "fmt"

func main() {
	// Greet("Gladys")
	// person := Person{
	// 	Name: "Joe Lap",
	// 	Age:  40,
	// }
	// fmt.Println(person.Name)
	// fmt.Println(person.Age)
	// person.Describe()

	Présenter(Animal{"Chien"})
	Présenter(Person{Name: "Joe", Age: 33})
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

func Présenter (d Describer) {
	d.Describe()
}
