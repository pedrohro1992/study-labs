package main

import (
	"fmt"
	"math"
)

type world struct {
	radius float64
}

type location struct {
	lat, long float64
}

/*
Marte tem um raio medio volumetrico de 3,389.5 kilometros
Ao inves de declarar 3389.5 como uma constante, usamos
o tipo world para declara Marte(Mars) como um dos muitos
mundos possiveis
*/

var mars = world{radius: 3389.5}

/*
Entao o metodo distance e anexado ao tipo world
Dando acesso ao campo radius. Ele aceita dois
parametros, os dois do tipo location, e retorna
a distancia em kilometros
*/

/*
Agora para o calculo de distancia. Ele usa uma seria de funcoes
de trigonometria, incluindo seno, cosseno e arco cosseno
*/

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

/*
O tipo location usa graus para altitude e longitude, mas as funcoes
de matematica na biblioteca padrao usa radianos. Dado que um circulo
tem 360 graus ou 2π radianos. A funcao abaixo faz as conversoes necessarias
*/

// rad converts degrees to radians
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	/*
	   Para ver os resultados de distance em acao
	   declaramos algumas localidades e usamos
	   as variaves de mars declarado anteriormente
	*/
	spirit := location{-14.5684, 175.472636}
	oppotunity := location{-1.9462, 354.4734}

	dist := mars.distance(spirit, oppotunity)
	fmt.Printf("%.2f km\n", dist)
}
