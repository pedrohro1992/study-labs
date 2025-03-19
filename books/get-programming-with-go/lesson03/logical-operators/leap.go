package main

import "fmt"

func main() {
	fmt.Println("The year 2100, should you leap?")

	year := 2100
	leap := year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap")
	} else {
		fmt.Println("Keep you feet on the ground")
	}
}
