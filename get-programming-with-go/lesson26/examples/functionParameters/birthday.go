package main

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}
