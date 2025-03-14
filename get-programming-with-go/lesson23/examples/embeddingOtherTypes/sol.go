package main

import "fmt"

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

type celsius float64

type sol int

/*
Qualquer metodo declarado no tipo sol, pode ser acessado atraves
do campo sol ou atraves do tipo report
*/
func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}
func main() {
	report := report{sol: 15}
	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))
}
