package main

import "fmt"

func main() {
	var bh float64 = 32767
	h := int16(bh)

	fmt.Println(h)
}
