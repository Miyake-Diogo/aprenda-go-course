package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "dormir"

	switch esporteFavorito {

	case "futebol":
		fmt.Println("futebol é tosco")
	case "basquete":
		fmt.Println("basquete é legal")
	case "dormir":
		fmt.Println("dormir é top")
	}

}
/*
- Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

*/