package main

import "fmt"

func main() {
	year := 2018
	fmt.Printf("Type %T for %v\n", year, year)

	// Ao inves de digitar a variavel duas vezes, podemos dizer para o Printf usar o
	// primeiro argumento no segundo verbo de formatacao
	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)
}
