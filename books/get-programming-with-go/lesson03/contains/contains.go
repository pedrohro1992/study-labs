package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("You find yourself in a dimly lit cavern.")

	command := "walk outside"
	exit := strings.Contains(command, "outside")

	fmt.Println("You leave the cave:", exit)
}
