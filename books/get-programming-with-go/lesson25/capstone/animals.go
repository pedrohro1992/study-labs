package main

import (
	"fmt"
	"math/rand"
	"time"
)

// honeyBee
type honeyBee struct {
	name string
}

// String returs the honeyBee name
func (hb honeyBee) String() string {
	return hb.name
}

// move makes a honeyBee move
func (hb honeyBee) move() string {
	switch rand.Intn(2) {
	case 0:
		return "buzzes about"
	default:
		return "flies to infinity and beyond"
	}
}

// eat makes a honeyBee eat
func (hb honeyBee) eat() string {
	switch rand.Intn(2) {
	case 0:
		return "pollen"
	default:
		return "nectar"
	}
}

// Cats
type cat struct {
	name string
}

// string return the cat name
func (c cat) String() string {
	return c.name
}

// moves makes a cat move
func (c cat) move() string {
	switch rand.Intn(3) {
	case 0:
		return "jumps arround"
	case 1:
		return "ask for food"
	default:
		return "is going to lie down"
	}
}

// eat makes a cat eat
func (c cat) eat() string {
	switch rand.Intn(5) {
	case 0:
		return "fish"
	case 1:
		return "chicken"
	case 2:
		return "turkey"
	case 3:
		return "cat food"
	default:
		return "churu"
	}
}

// animal interface calls the move and eat methods
type animal interface {
	move() string
	eat() string
}

func step(a animal) {
	switch rand.Intn(2) {
	case 0:
		fmt.Printf("%v %v.\n", a, a.move())
	default:
		fmt.Printf("%v eats the %v.\n", a, a.eat())
	}
}

const sunrise, sunset = 8, 18

func main() {
	rand.Seed(time.Now().UnixNano())

	animals := []animal{
		honeyBee{name: "Bzz Lightyear"},
		cat{name: "Galadriel"},
	}

	var sol, hour int

	for {
		fmt.Printf("%2d:00 ", hour)
		if hour < sunrise || hour >= sunset {
			fmt.Println("The animals are sleeping")
		} else {
			i := rand.Intn(len(animals))
			step(animals[i])
		}

		time.Sleep(500 * time.Millisecond)
		hour++

		if hour >= 24 {
			hour = 0
			sol++
			if sol >= 3 {
				break
			}
		}
	}
}
