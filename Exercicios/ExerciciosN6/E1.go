package main

import (
	"fmt"
)

func main() {
	a := 10
	b := 20
	c := "|Uai fio, tudo bem ai?|"
	fmt.Println(retornaIntMaisUm(a))
	fmt.Println(retornIntString(b, c))

}

func retornaIntMaisUm(x int) int {
	x = 1 + x
	return x
}

func retornIntString(x int, s string) (int, string) {

	return x, s

}