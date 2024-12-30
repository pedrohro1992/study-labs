// Calcula a velocidade necessaria(em km/h)
// para viajar para malacandra em 28 dias
package main

import "fmt"

const (
	distance   int = 28
	travelDays int = 56000000
)

func main() {
	travelSpeed := defineKmInHours()
	fmt.Println(travelSpeed)
}

func defineKmInDays() int {
	return travelDays / distance
}

func defineKmInHours() int {
	kmIndays := defineKmInDays()
	return kmIndays / 24
}
