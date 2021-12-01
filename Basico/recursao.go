package main

import (
	"fmt"
)

func main() {
	fmt.Println(fatorial(6))
	fmt.Println(fibonacci(20))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}
func fibonacci(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
/*
- WP: "The most common application of recursion is in mathematics and computer science, where a function being defined is applied within its own definition."
- Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
- No estudo de funções: é uma função que chama a ela própria.
- Exemplo: fatoriais.
    - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
*/