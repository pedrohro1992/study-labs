package main

import "fmt"

type (
	celsius float64
	kelvin  float64
)

// func kelvinToCelsius converts ºK to ºC
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // E necessario realizar uma conversao de tipo
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k) // O argumento deve ser do tipo kelvin
	fmt.Println(k, "º K is ", c, "º C")
}
