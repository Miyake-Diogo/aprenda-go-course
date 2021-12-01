package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

func (q quadrado) area() float64 {
	return q.lado * q.lado
}

type circulo struct {
	raio float64	
}
func (c circulo) area() float64 {	
	return math.Pi * c.raio * c.raio
}

type info interface {
	area() float64
}
func medida(i info) {
	fmt.Println("A area é:", i.area())
}

func main() {
	
	x := quadrado{
		lado: 25.0,
	}

	y := circulo{
		raio: 3.43,
	}
	medida(x)
	medida(y)
}
/*
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"


*/

