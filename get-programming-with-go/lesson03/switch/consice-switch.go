package main

import "fmt"

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	command := "go inside"

	switch command { // Compara casos com "command"
	case "go east":
		fmt.Println("You head further up the montain")
	case "enter cave", "go inside": // Virgulas separam uma lista de valores possiveis
		fmt.Println("You find yourself dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that.")
	}
	// O Resultado sera:
	// You find yourself dimly lit cavern.
}
