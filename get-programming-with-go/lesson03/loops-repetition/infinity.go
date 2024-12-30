package main

import (
	"fmt"
	"math/rand"
)

func main() {
	degress := 0

	for {
		fmt.Println(degress)
		degress++
		if degress >= 360 {
			degress = 0
			if rand.Intn(2) == 0 {
				break
			}
		}
	}
}
