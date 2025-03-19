package main

import (
	"fmt"
	"math/rand"
)

func main() {
	numberToGuess := 31

	for {
		numberToTry := rand.Intn(100)
		fmt.Println(numberToTry)
		if numberToTry == numberToGuess {
			fmt.Println("Acertooo!!")
			break
		} else if numberToTry%numberToGuess != 0 {
			fmt.Println("Passou Longe")
		} else if numberToTry%numberToGuess == 0 {
			fmt.Println("QUASE!!!")
		}
	}
}
