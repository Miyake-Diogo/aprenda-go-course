package main

import (
	"fmt"
)

var y = "OI!!" // package level scope
func main() {
	_ , erros := fmt.Println("Hello, playground", "E ai beleuza?", 100)
	fmt.Println(erros)
	
	// variaveis := atribuição e eclaração
	x := 10 // code block scope

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	
	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
	
	// = é usado para atribuir valores a variáveis

	
}
