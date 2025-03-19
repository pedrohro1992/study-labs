package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	fmt.Println(len(worlds))
	fmt.Println(len(terrestrial))
	fmt.Println(planets)

	// Se o terceiro index nao for especificado, terrestrial vai ter a capacidade de 8
	// Fazer o append de Ceres nao aloca um array novo, ao invezs disso,
	// sobrescreve Jupiter

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(len(terrestrial))
	fmt.Println(len(planets))
	fmt.Println(planets)
}
