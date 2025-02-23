package main

import "fmt"

func main() {
	// Declara uma slice vazia
	s := []string{}
	// Declara lastCap e faz o init com o cap de s
	lastCap := cap(s)

	// Loop para fazer o append em s
	for i := 0; i < 1000; i++ {
		s = append(s, "An element")
		// Verifica se a capacidade de verifica se a capacidade de s e maior que lastCap
		if cap(s) != lastCap {
			fmt.Println(cap(s))
			// Guarda cap de s em lastCap para ser usada na proxima interacao do loop
			lastCap = cap(s)
		}
	}
}
