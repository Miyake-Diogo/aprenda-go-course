package main

import (
	"fmt"
)

func main() {
	goehlegal := 3
	
	switch {
	case goehlegal == 0:
		fmt.Println("go é legal mas prefiro Rust")
	
	case goehlegal == 1:
		fmt.Println("golang da pro gasto")
	
	default:
		fmt.Println("Elixir é top também")
		}
	
}
/*
- Crie um programa que utilize a declaração switch, sem switch statement especificado.

*/