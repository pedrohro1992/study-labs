package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func (p *person) birthday() {
	p.age++
}

func main() {
	terry := &person{
		name: "Terry",
		age:  15,
	}
	terry.birthday()
	fmt.Printf("%+v", terry)
}
