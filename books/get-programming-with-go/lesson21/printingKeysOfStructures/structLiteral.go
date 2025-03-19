package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	curiosity := location{-4.5895, 137.4417}

	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)
}
