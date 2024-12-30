package main

import "fmt"

func main() {
	celsius := 21.0
	// Com erro de arrendotamento

	fmt.Print((celsius/5.0*9.0)+32, "º F\n")
	fmt.Print((9.0/5.0*celsius)+32, "º F\n")

	// Realizando a multiplicacao primeiro
	fahrenheit := (celsius * 9.0 / 5.0) + 32
	fmt.Print(fahrenheit, "º F\n")
}
