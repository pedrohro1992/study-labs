package main

import "fmt"

func main() {
	room := "lake"

	switch { // Expressoes estao em cada caso
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("The ice seems solid enough.")
		fallthrough // Cai para o proximo caso
	case room == "underwater":
		fmt.Println("The water is freezing cold")
	}
	// Os valores serao:
	// The ice seems solid enough.
	// The water is freezing cold
}
