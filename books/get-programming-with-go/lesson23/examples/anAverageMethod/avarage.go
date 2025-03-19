package main

import "fmt"

type report struct {
	sol         int
	temperature temperature // O campo temperatura e uma estrutura do tipo temperature
	location    location
}

type temperature struct {
	high, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

func main() {

	/*
	   O tipo temperatura e o metodo average podem ser usados
	   independetes do report do tempo, como abaixo
	*/
	t := temperature{high: -1.0, low: -78.0}
	fmt.Printf("average %v\n", t.average())
}
