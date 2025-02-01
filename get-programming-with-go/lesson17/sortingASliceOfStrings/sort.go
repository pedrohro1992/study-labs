package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	sort.StringSlice(planets).Sort() // Deixa planets em ordem alfabetica
	fmt.Println(planets)
}
