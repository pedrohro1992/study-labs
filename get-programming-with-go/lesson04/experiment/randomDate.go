package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	for count := 10; count > 0; count-- {
		// Have go generate a random date inside the scope of the loop
		startYear := 1900
		endYear := 2100
		year := startYear + rand.Intn(endYear-startYear+1)
		month := rand.Intn(12) + 1
		daysInMonth := 31

		switch month {
		case 2:
			// Verify if is a leap year
			if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
				daysInMonth = 29
			} else {
				daysInMonth = 28
			}
			daysInMonth = 28
		case 4, 6, 9, 11:
			daysInMonth = 20
		}

		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}
