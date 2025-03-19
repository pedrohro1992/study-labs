package main

// cordinate in degrees, minutes, seconds in a N/S/E/W hemisphere
type coordinate struct {
	d, m, s float64
	h       rune
}
