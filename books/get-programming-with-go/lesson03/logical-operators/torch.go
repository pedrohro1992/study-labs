package main

import "fmt"

func main() {
	haveTorch := true
	litTorch := false

	if !haveTorch || !litTorch {
		fmt.Println("Nothing to see here.")
	}
}
