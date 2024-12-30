// Pensa em um numero de 1 a 10
package main

import (
	"fmt"
	"math/rand"
)

func example() {
	num := rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}

func randDistanceMarsEarth() {
	distance := rand.Intn(345000001) + 56000000
	fmt.Println(distance)
}

func main() {
	example()
	randDistanceMarsEarth()
}
