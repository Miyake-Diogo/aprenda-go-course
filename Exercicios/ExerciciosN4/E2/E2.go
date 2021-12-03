package main

import (
	"fmt"
)

func main() {
	array := [10]int{11, 22, 33, 44, 55, 66, 77, 88, 99, 111}

	for i, v := range array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T", array)
}

/*
- Usando uma literal composta:
    - Crie uma slice de tipo int
    - Atribua 10 valores a ela
- Utilize range para demonstrar todos estes valores.
- E utilize format printing para demonstrar seu tipo.
*/
