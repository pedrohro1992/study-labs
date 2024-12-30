package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10 // Declara e inicializa

	for count > 0 { // Uma condicao
		fmt.Println(count)
		time.Sleep(time.Second)
		count-- // Faz o decremento de count, caso contrario, o loop roda para sempre
	}
	fmt.Println("Liftoff!")
}
