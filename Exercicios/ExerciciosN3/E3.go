package main

import (
	"fmt" 
	"time")

func main() {
	anoNascimento := 1990
	t := time.Now()
	anoAtual := t.Year()
	for anoNascimento<= anoAtual {
		fmt.Println(anoNascimento)
		anoNascimento++
	}
}
/*
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/