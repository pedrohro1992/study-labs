package main

import "fmt"

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // TODO implemet a real sensor
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // Declara e retorna uma funcao anonima
		return s() + offset
	}
}

func main() {
	var sensorReading kelvin = 98
	sensor := calibrate(realSensor, sensorReading)
	fmt.Println(sensor())
}
