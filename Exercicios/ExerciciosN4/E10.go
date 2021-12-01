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
	
	myMap["loureiro_kiko"] = []string{"usar os trequinho no punho", "tacar fogo na guitarra"}

	delete(myMap, "loureiro_kiko")
	for key, value := range myMap {
		fmt.Println(key)
		for i, h := range value {
			fmt.Println("\t", i, " - ", h)
		}
	}
}

/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
*/