package main

import "fmt"

func main() {
	fmt.Printf("%v is %[1]T\n", "literal string")
	fmt.Printf("%v is %[1]T\n", `raw string literal`)
}
