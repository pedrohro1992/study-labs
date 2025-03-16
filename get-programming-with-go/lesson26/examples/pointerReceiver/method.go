package main

type person struct {
	name, superpower string
	age              int
}

func (p *person) birthday() {
	p.age++
}

func main() {
}
