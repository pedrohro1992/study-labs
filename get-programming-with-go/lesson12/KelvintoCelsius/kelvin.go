package main

import "fmt"

// kelvinToCelsius converts Kº to Cº
func kelvinToCelsius(k float64) float64 { // Declara uma funcaoo que aceita um parametro e retorna um resultado
	k -= 273.15
	return k
}

func main() {
	kelvin := 294.0
	celsius := kelvinToCelsius(kelvin) // Chama a funcao passando kelvin como primeiro argumento
	fmt.Println(kelvin, "Kº is ", celsius, "Cº")
}
