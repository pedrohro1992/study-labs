package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 0
	for count < 10 { // Um novo escopo comeca
		num := rand.Intn(10) + 1
		fmt.Println(num)
		count++
	} // O escopo termina
}
