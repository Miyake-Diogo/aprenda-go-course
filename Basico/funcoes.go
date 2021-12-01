package main
import (
	"fmt"
)


func soma(x int, y int) int {
	return x + y
}

func basica() {
	fmt.Println("Oi, bom dia!")
}

				// (argumentos) (Retorno da função)
func multiplica(x ...int) (int, int, string) {
	multiplica := 1
	for _, v := range x {
		multiplica *= v
	}
	return multiplica, len(x), "Multiplicado!"
}

func main() {
	// funçao que retorna algo
	x, y := 2, 3
	fmt.Printf("Soma de x: %v + y: %v = %v \n", x, y ,soma(x, y))

	// Função que não tem argumento -> len (retorna o comprimento de uma string)
	const z = "Eai brother?"
	fmt.Println(len(z))

	// Função que não recebe argumento -> basica()
	basica()

	// Função que recebe argumentos variados -> multiplica()
	total, quantos, oi := multiplica(10, 10, 1, 2, 3, 5)
	fmt.Println(total, quantos, oi)

}
