package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %vº C.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println(temperature)

	// Se for acessado uma chave que nao existe no map
	// o resultado e valo zerado para o tipo (int)
	moon := temperature["Moon"]
	fmt.Println(moon)

	//Go prove a syntax "virgula, ok", que pode ser usada para distinguir entr
	//"Moon" nao existir no map versus estar presente no map com a temperatura de 0
	if moon, ok := temperature["Moon"]; ok { // NOTA: No lugar de ok, pode ser usado qualquer nome de variavel
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
