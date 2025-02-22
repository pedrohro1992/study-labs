package main

import "fmt"

func main() {
	dwarfs := make([]string, 0, 10)
	fmt.Println(len(dwarfs))
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(len(dwarfs))
}
