package main

import (
	"fmt"
	"sort"
)

var temperature = []float64{
	-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
}

func main() {
	set := make(map[float64]bool)

	for _, t := range temperature {
		set[t] = true
	}

	if set[-28] {
		fmt.Println("set member")
	}

	fmt.Println(set)

	// Podemos ver que o map contem apenas uma chave
	// para cada temperatura, com cada duplicacao removida
	// mas chaves de map tem uma ordem arbritaria no Go, entao
	// antes de elas poderem ser usadas, as temperaturas tem que ser
	// convertidas para um slice novamente
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
