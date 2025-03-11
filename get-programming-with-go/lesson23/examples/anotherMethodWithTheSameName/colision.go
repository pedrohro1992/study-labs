package main

type report struct {
	sol
	temperature // O tipo temperature incorparado em report
	location    location
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

type location struct {
	lat, long float64
}

func (l location) days(l2 location) int {
	// TODO add some math
	return 5
}

func (r report) days(s2 sol) int {
	return r.sol.days(s2)
}

type celsius float64

type sol int

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

func main() {
}
