package main

import (
	"errors"
	"fmt"
)

type Student struct {
	Name 	string
	Grade int
	Score float64
}

type Reportable interface {
	Report()
}

func (s Student) Report() {
	fmt.Printf("%s - Grade %d - Score %.2f \n", s.Name, s.Grade, s.Score)
}

func AfficherEtudiant(r Reportable) {
	r.Report()
}

func NouvelEtudiant(name string, grade int, score float64) (Student, error) {
	if(score < 0 || score > 100) {
		return Student{}, errors.New("Pas de score \n")
	} else {
		return Student{Name: name, Grade: grade, Score: score}, nil
	}
}
/*
Crée un struct Student avec Name (string), Grade (int) et Score (float64)
Ajoute une méthode Report() qui affiche "<Name> - Grade <Grade> - Score: <Score>"
Crée une interface Reportable avec une méthode Report()
Crée une fonction AfficherEtudiant(r Reportable) qui appelle r.Report()
Crée une fonction NouvelEtudiant(name string, grade int, score float64) (Student, error) qui retourne une erreur si score est inférieur à 0 ou supérieur à 100
Dans main, crée 3 étudiants — un valide, un avec score négatif, un avec score > 100
*/