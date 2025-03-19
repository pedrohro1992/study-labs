package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Mercury" // Atribui um planeja no index 0
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] // Recupera o planeta no index 2
	fmt.Println(earth)
}
