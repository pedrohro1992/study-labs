package main

import (
	"fmt"
)

type person struct {
	name, superpower string
	age              int
}

func main() {
	timmy := &person{
		name: "Timmoty",
		age:  10,
	}
	timmy.superpower = "flying"
	fmt.Printf("%+v\n", timmy)
}
