package main

import (
	"fmt"
	"math/rand"
)

const (
	nickels  = 0.05
	dimes    = 0.10
	quarters = 0.25
)

func main() {
	for piggyAccount := 0.0; piggyAccount < 20.00; {
		switch rand.Intn(3) {
		case 0:
			piggyAccount += nickels
		case 1:
			piggyAccount += dimes
		case 2:
			piggyAccount += quarters
		}

		fmt.Printf("$%5.2f\n", piggyAccount)
	}
}
