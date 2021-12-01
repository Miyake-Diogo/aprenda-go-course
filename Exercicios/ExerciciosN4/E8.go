package main

import (
	"fmt"
)

func main() {
	myMap := map[string][]string{
		"Wedge_Ro": []string{
			"Ser zé Ruela","Ferrar com a vida do próximo"},
		"Trash_Claube":   []string{
			"Não fazer nada", "o caminho mais fácil é o melhor","Usar ferramentas toscas"},
		"So_Ever":   []string{
			"Ficar de futrico","Tomar cerva em canequinha termica"},
		"Lama_Andy":   []string{
			"Jogar playlacration com os friends","Ficar a toa"},
	}
	
	for key, value := range myMap {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}

}
/*
- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/