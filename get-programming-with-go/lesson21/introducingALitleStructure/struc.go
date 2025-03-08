package main

import "fmt"

func main() {
	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895   // Atribui valores para os campos da estrutura
	curiosity.long = 137.4417 // Atribui valores para os campos da estrutura

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)
}
