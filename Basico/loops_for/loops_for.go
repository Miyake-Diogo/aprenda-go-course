package main

import (
	"fmt"
)

func main() {
	for c := 0; c < 4; c++ {
	fmt.Println(c)
		for d := 5; d < 10; d++ {
			fmt.Println(d)
			}
	}
    x := 0
	for x <= 10 {
		fmt.Println(x)
		x++
	}
	y := 0
	for {
		if y < 10 {
			y++
			fmt.Println("y é menor que dez")
		} else { 
			fmt.Println("y é maior que dez")
			break
		}
	}
	for i := 0; i < 20; i++ {
		if i == 3 {
			// faz isso quando o número não é par
			continue
		}
		fmt.Println(i)
	}
	for x := 33; x <= 122; x++ {
		fmt.Printf("%d - %v\n", x, string(x))
	}
}

/*
- For
    - Inicialização, condição, pós
    - Ponto e vírgula?
    - gobyexample.com
    - Não existe while! 

	- For
    - Repetição hierárquica
    - Exemplos: relógio, calendário 
	- For: inicialização, condição, pós
	- For: condição ("while")
	- For: ...ever? (http servers)
	- For: break
	- golang.org/ref/spec#For_statements, Effective Go
	- (Range vem mais pra frente.) 

	- Operação módulo: %
	- For: break
	- For: continue

	- Desafio surpresa!
	- Format printing:
		- Decimal       %d
		- Hexadecimal   %#x
		- Unicode       %#U
		- Tab           \t
		- Linha nova    \n
	- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
*/
