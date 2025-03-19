package main

import "fmt"

var f = func() { // Atribui uma funcao anonima a uma variavel
	fmt.Println("Dress up for the masquerade.")
}

func main() {
	f() // faz a chamada para a variavel com a funcao anonima
}
