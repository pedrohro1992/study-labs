package main

import "fmt"

func main() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ { // Faz iteracao por cada caractere ASCII
		c := message[i]
		// verifica se c esta entre a e z
		if c >= 'a' && c <= 'z' { // Deixa os espacos e pontuacoes como estao
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
