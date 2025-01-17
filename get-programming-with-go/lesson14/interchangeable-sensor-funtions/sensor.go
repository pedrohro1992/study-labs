package main

import (
	"fmt"
	"math/rand"
)

type kelvin float32

func fakeSenson() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 // TODO: implement a real sensor
}

func main() {
	sensor := fakeSenson // Atribui uma funcao a uma variavel
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
