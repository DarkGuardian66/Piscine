package main

import "fmt"

func main() {
	var nom string
	var age int
	fmt.Println("Entrez votre nom :")
	fmt.Scanf("%s", &nom)
	fmt.Println("Entrez votre âge:")
	fmt.Scanf("%d", &age)

	fmt.Printf("Bonjour, %s ! Vous avez %d", nom, age)
}
