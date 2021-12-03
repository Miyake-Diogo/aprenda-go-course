package main

import (
	"fmt" 
	"time")

func main() {
	anoNascimento := 1990
	t := time.Now()
	anoAtual := t.Year()
	for  {
	if anoNascimento <= anoAtual {
		fmt.Println(anoNascimento)
		anoNascimento++
	} else { break}
	}
}
/*
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/