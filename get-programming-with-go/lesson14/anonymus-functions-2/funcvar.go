package main

import "fmt"

func main() {
	f := func(message string) { // Atribui uma funcao anonima a uma variavel
		fmt.Println(message)
	}
	f("Go to the party.") // Faz a chamada da variavel com a funcao anonima e printa message
}
