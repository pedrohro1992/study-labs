package main

import "fmt"

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)

	// Para mostrar os carcteres ao inves dos numeros, o verbo de formatacao %c pode ser usado com Printf
	fmt.Printf("%c%c%c%c\n", pi, alpha, omega, bang)
	// PS Quanlquer tipo inteiro pode fazer uso do %c, mas o alias rune indica que o numero 650 representa um caractere
}
