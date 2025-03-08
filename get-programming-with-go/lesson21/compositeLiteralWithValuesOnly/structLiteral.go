package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	spirit := location{-14.5684, 175.472636}
	fmt.Println(spirit)
}
