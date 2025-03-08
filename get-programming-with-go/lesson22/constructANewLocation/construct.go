package main

type location struct {
	lat, long float64
}

// decimal convertes a d/m/s coordinate to decimal degrees.
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}
