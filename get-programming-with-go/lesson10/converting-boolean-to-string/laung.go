package main

import "fmt"

func main() {
	lauch := false

	launchText := fmt.Sprintf("%v", lauch)
	fmt.Println("Ready for launch: ", launchText)

	var yesNO string
	if lauch {
		yesNO = "yes"
	} else {
		yesNO = "no"
	}
	fmt.Println("Ready for launch: ", yesNO)
}
