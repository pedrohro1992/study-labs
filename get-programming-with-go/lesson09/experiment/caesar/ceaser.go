package main

import "fmt"

func main() {
	message := "L fdph, L vdz, L frqtxhuhg."

	for i := 0; i < len(message); i++ { // Faz iteracao por cada caractere ASCII
		c := message[i]
		// verifica se c esta entre a e z para deixar os espacos e pontuacoes como estao
		if c >= 'a' && c <= 'z' { // Deixa os espacos e pontuacoes como estao
			c -= 3
			// Verifica se precisa fazer a volta
			if c < 'a' {
				c = +26
			}
			// verifica se c esta entre A e Z para deixar os espacos e pontuacoes como estao
		} else if c >= 'A' && c <= 'Z' {
			c -= 3
			// Verifica se precisa fazer a volta
			if c < 'A' {
				c += 26
			}

		}
		fmt.Printf("%c", c)
	}
}
