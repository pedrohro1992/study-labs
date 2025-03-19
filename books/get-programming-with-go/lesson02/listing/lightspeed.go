// Quando tempo demora pra chegar em marte

package main

import "fmt"

func main() {
	const lightSpeed = 299792 // km/s
	distance := 56000000      // km

	fmt.Println(distance/lightSpeed, "seconds") // Prints 186 seconds
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds") // Prints 1337 seconds
}
