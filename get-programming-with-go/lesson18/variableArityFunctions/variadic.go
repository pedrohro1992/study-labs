package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

// O parametro worlds e um slice de strings que contem
// zero ou mais arumentos passados para terraform

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println(twoWorlds)
	// Para passar uma slice de multiplos argumentos, expandimos
	// a slice com tres pontos
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)
}
