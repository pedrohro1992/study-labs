package main

import "fmt"

func main() {
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783) // Prints My weight on the surface of Mars is 56.3667 lbs
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)               //  Prints and I would  be 21 years old.

	// Printf recebe dois argumentos e faz a substuicao dos verbos de formatacao na ordem
	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)

	// Adicionando pads na string
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}
