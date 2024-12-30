package main

import "fmt"

const (
	lightSpeed    = 299792
	secondsPerDay = 86400
)

func main() {
	var distance int64 = 41.3e12
	fmt.Println("Alpha Centaury is ", distance, "km away.")

	days := distance / lightSpeed / secondsPerDay
	fmt.Println("That is", days, "days of travel at light speed")
}
