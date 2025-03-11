package main

import "fmt"

type report struct {
	sol         int
	temperature // O tipo temperature incorparado em report
	location    location
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	lat, long float64
}

type celsius float64

func main() {
	/*
		Todos os metodos no tipo temperature sao feitos
		acessiveis atraves do tipo report
	*/
	report := report{
		sol:         15,
		location:    location{-4.5895, 137.4417},
		temperature: temperature{high: -1.0, low: -78.0},
	}
	fmt.Printf("average %v C\n", report.average())
	/*
		Mesmo que nenhum nome de campo seja especificado, um campo existe
		com o mesmo nome que no tipo encorporado. Podemos acessar temperatura
		da seguinte forma
	*/
	fmt.Printf("average %v C\n", report.temperature.average())
	/*
		Incorporar nao apenas encaminha metodos. Campos de ums estrutura interna
		sao acessiveis a partir da outra estrutura. Em adicao a report.temperature.high
		podemos acessar a alta temperatura(high temperature) com report.high:
	*/
	fmt.Printf("%v C\n", report.high)
	report.high = 32
	fmt.Printf("%v C\n", report.temperature.high)
}
