package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}
	groups := make(map[float64][]float64) // Um map com chaves em float64 e valores []float64

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // Arrendonda a temperatura para -20, -30, e assim por diante
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}
}
