package main

import "fmt"

type location struct {
	lat  float64
	long float64
}

func main() {
	lats := []float64{-4.5895, -14.5684, -1.9462}
	longs := []float64{137.4417, 175.472636, 354.4734}
	fmt.Println(lats, longs)
}
