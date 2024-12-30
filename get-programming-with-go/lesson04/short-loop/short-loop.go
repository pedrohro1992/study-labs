package main

import (
	"fmt"
)

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	} // count nao esta mais em escopo
}
