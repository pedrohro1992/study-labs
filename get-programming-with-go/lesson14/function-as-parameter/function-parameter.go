package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func measureTemperature(samples int, sensor func() kelvin) { // measureTemperature aceota uma funcao como segundo parametro
	for i := 0; 1 < samples; i++ {
		k := sensor()
		fmt.Printf("%vº K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor) // Passa o nome de uma funcao para uma funcao
}
