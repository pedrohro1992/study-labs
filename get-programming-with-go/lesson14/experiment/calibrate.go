package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // TODO implemet a real sensor
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // Declara e retorna uma funcao anonima
		return s() + offset
	}
}

func main() {
	var offset kelvin = 5
	// Toda chamada para sensor() e tambem uma chamada pra calibrate()
	// que, nesse exemplo, a chamada e feita usando o resultado da fakeSensor
	// o que resulta em, o valor de sensor() sendo modificado mesmo ele estando fora
	// do body do for

	sensor := calibrate(fakeSensor, offset)
	for count := 0; count < 10; count++ {
		// Faz a chamada para a sensor() que faz a chamada para calibrate()

		fmt.Println(sensor())
	}
}
