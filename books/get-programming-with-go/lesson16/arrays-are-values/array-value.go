package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	planetsMarkII := planets // Copia o array planets
	planets[2] = "woops"     // Da um jeito de fazer um bypass interstellar

	fmt.Println(planets)
	fmt.Println(planetsMarkII)
}
