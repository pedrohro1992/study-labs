package main

import (
	"fmt"
	"math/rand"
)

const (
	secondsPerDay  = 86400
	distanceToMars = 62100000
)

func main() {
	trip := ""
	company := ""
	fmt.Println("Spaceline       ", "Days", "Trip", "type", "Price")
	fmt.Println("=================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		// Define company
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX          "
		case 2:
			company = "Virgin Galactic "
		}
		// define days
		speed := rand.Intn(15) + 16
		days := distanceToMars / speed / secondsPerDay

		// define price
		price := 20.0 + speed

		// define trip
		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", company, days, trip, price)
	}
}
