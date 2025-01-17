package main

import "fmt"

func main() {
	func() { // Declara uma funcao anonima
		fmt.Println("Functions anonymous")
	}() // Invoca a funcao
}
