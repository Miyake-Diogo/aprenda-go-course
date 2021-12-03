package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		[]string{
			"Claube",
			"Trash",
			"Puxar o saco eentregar pastel",
		},
		[]string{
			"Junes",
			"Puxa",
			"Resolver problemas que não existem",
		},
		[]string{
			"Ro",
			"wedge",
			"Usar do abuso de poder",
		},

	}
	
	for _, v := range ss {
		fmt.Println(v)
		}
		
		fmt.Println("\n\n")

	for _, v := range ss {
		fmt.Println(v[0])
		for _, item := range v {
			fmt.Println("\t", item)
		}

	}

}

/*

- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
    - "Nome", "Sobrenome", "Hobby favorito"
- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

*/