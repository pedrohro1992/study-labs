package main

import "fmt"

func main() {
	third := 1.0 / 3.0
	fmt.Println(third + third + third) // printa um mesmo o resultado exato sendo 0.99999...

	piggyBank := 0.1
	piggyBank += 0.2

	fmt.Println(piggyBank) // Nao faz a soma exata de 0.1 + 0.2, que seria 0.3, ao invez disso, se perde na representacao e mostra um numero bizarro
}
