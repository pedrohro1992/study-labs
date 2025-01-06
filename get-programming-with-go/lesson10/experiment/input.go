package main

import "fmt"

func main() {
	stringToConvert := "yes"
	var convertedBool bool

	switch stringToConvert {
	case "true", "yes", "1":
		convertedBool = true
	case "false", "no", "0":
		convertedBool = false
	default:
		fmt.Println(stringToConvert, "Is not valid")
	}
	fmt.Println("Ready for launch:", convertedBool)
}
