package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	// A funcao append e variadic, como Println, entao podemos passar elementos
	// multiplos para fazer o append de uma so vez
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}
