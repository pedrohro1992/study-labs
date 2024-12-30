package main

import "fmt"

func main() {
	command := "go east"

	if command == "go east" { // se command e igual a "go east"
		fmt.Println("You head further up the montain")
	} else if command == "Go inside" { // Ou se command e igual "go inside"
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.") // Ou, se qualquer outra coisa
	}
}
