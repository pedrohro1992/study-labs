package main

import "fmt"

func main() {
	var board [8][8]string // Um array de 8 arrays de 8 strings

	board[0][0] = "r"
	board[0][7] = "r" // Coloca uma torre na cordenada [row][column]

	for column := range board[1] {
		board[1][column] = "p"
	}
	fmt.Println(board)
}
