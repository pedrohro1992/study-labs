package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // era fica disponivel por todo o pacote

func main() {
	year := 2018 // era e year estao em escopo

	switch month := rand.Intn(28) + 1; month { // era, year, month estao em escopo
	case 2:
		day := rand.Intn(28) + 1 // era, year, month e day estao em escopo
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1 // muda o valor de day
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1 // muda o valor de day
		fmt.Println(era, year, month, day)
	} // month e day estao fora de escopo
} // year esta fora de escopo
