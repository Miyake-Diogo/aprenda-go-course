package main

import (
	"fmt"
)

var z interface{}

func main() {

	if x := 10; !(x > 100) {
		fmt.Println("Hello, playground")
	}

	if x := 500; x > 100 {
		fmt.Println("chis é maior que cem")
	} else if x < 10 {
		fmt.Println("chis é menor que déis")
	} else {
		fmt.Println("chis não é menor que déis nem maior que cem")
	}

	y := 10
	switch {
	case (y > 10), (y < 20):
		fmt.Println("ipísilon é maior que déis")
	case y < 10:
		fmt.Println("ipísilon é menor que déis")
	case y == 10:
		fmt.Println("ipísilon é igual a déis")
	default:	
		fmt.Println("ipísilon não é menor, nem maior que déis")

	}
	z = true
	switch z.(type) {
		case int:
			fmt.Println("ipísilon é um inteiro")
		case string:
			fmt.Println("ipísilon é uma string")
		case bool:
			fmt.Println("ipísilon é um booleano")
		case float64:
			fmt.Println("ipísilon é um float64")
		default:
			fmt.Println("ipísilon é um tipo desconhecido")
	}

	a := 9

	if !(a%2 == 0) && a%3 == 0 {
		fmt.Println("é múltiplo de dois e tambem de treis")
	}
}
/*

- If: bool
- If: o operador não → "!"
- If: declaração de inicialização

- If, else.
- If, else if, else.
- If, else if, else if, ..., else.

- Switch:
    - pode avaliar uma expressão 
        - switch statement == case (value)
        - default switch statement == true (bool)
    - não há fall-through por padrão
    - criando fall-through
    - default
    - cases compostos

- Pode avaliar tipos
- Pode haver uma expressão de inicialização
- &&
- ||
- !
- Qual o resultado de fmt.Println...
    - true && true
    - true && false
    - true || true
    - true || false
    - !true
*/