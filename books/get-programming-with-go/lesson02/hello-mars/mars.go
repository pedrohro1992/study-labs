// Meu programa de perca de peso >> //comentario para leitores humanos
package main

import "fmt"

// main e a funcao onde tudo comeca
func main() { // O compilador em go, coloca um ; em todo final de linha, pra evitar erro, sempre coloca o { na mesma linha
	fmt.Print("My weight on the surface of Mars is ")
	fmt.Print(149.0 * 0.3793) // Prints 56.3667
	fmt.Print(" lbs, and i would be ")
	fmt.Print(41 * 365 / 687) // Prints 56.3667
	fmt.Print(" years old.")
	// Faz a mesma coisa, mas usando uma unica chamada pro fmt.Println e passando os calculos
	// como uma lista de argumentos para a funcao
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and I would be", 41*365.2425/687, "years old.")
}
