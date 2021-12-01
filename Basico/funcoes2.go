package main

import (
	"fmt"
)

func main() {

	// função anônima

	lambda := 387

	func(x int) {
		fmt.Println(x, "vezes 873648 é:")
		fmt.Println(x * 873648)
	}(lambda)

	
	

	// Função como expressão
	y := func(x int) int {
		//fmt.Println(x, "vezes 57 é:")
		return x * 57
	}

	fmt.Println(lambda, "vezes 57 é:", y(lambda))
	
	// Funçao que retorna outra função
	x := retornaumafuncao()
	z := x(3)

	fmt.Println(z)

	fmt.Println(retornaumafuncao()(3))

}

// Funçao que retorna outra função
func retornaumafuncao() func(int) int {
	return func(i int) int {
		return i * 10
	}
}
