package main

// import "fmt"
/*
C'est quoi un pointeur ? 
Quand tu crées une variable, Go réserve un espace en mémoire pour stocker sa valeur. Cet espace a une adresse.

Un pointeur c'est simplement une variable qui stocke une adresse mémoire au lieu d'une valeur.

Avant d'aller plus loin — dans ta tête, c'est quoi la différence entre ces deux lignes ?
name := "Joe"
name2 := &name

Maintenant la question clé — pourquoi voudrait-on stocker une adresse plutôt qu'une valeur ?
Imagine que tu passes un User à une fonction. Go fait une copie de ce User. Si la fonction modifie le User, elle modifie la copie — l'original ne change pas.
Avec un pointeur, tu passes l'adresse de l'original. La fonction peut alors modifier directement l'original.

Exercice 4 : 
1. Crée une fonction doubler(n int) qui multiplie n par 2
2. Appelle-la depuis main avec x := 5, affiche x avant et après
3. Observe que x n'a pas changé
4. Crée une fonction doublerPtr(n *int) qui fait la même chose mais avec un pointeur
5. Appelle-la avec &x et observe la différence
*/

// func main() {
// 	x := 5
// 	fmt.Println("avant:", x)
// 	Doubler(x)
// 	fmt.Println("après:", x)
// 	fmt.Println("avant:", x)
// 	DoublerPtr(&x)
// 	fmt.Println("avant:", x)

// }

// func Doubler(n int) {
// 	result := n * 2
// 	fmt.Println(result)
// }

// func DoublerPtr(n *int) {
// 	result := *n * 2
// 	*n = *n * 2
// 	fmt.Println(result)
// }


