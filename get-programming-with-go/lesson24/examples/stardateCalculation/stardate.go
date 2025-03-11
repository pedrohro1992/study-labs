package main

import "time"

// stardate returns a fictional measure of the for a givem date

func stardate(t time.Time) float64 {
	doy := float64(t.YearDay())
}
