package main

import (
	"fmt"
	"math/rand"
)

const (
	nickels  = 5
	dimes    = 10
	quarters = 25
)

func main() {
	for piggyAccount := 0; piggyAccount < 2000; {
		switch rand.Intn(3) {
		case 0:
			piggyAccount += nickels
		case 1:
			piggyAccount += dimes
		case 2:
			piggyAccount += quarters
		}

		dollars := piggyAccount / 100
		cents := piggyAccount % 100
		fmt.Printf("$%d.%02d\n", dollars, cents)

	}
}
