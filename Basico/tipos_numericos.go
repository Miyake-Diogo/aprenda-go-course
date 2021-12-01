package main

import (
	"fmt"
	"runtime"
)

var x = 10

func main() {
	a := "e"
	b := "é"
	c := "💩"

	fmt.Printf("%v, %v, %v\n", a, b, c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)

	fmt.Printf("%v, %v, %v\n", d, e, f)

	y := 10.0
	fmt.Printf("x: %v -> %T\n", x, x)
	fmt.Printf("y: %v -> %T\n", y, y)

	fmt.Println("OS onde go esta sendo executado: ", runtime.GOOS)
	fmt.Println("Arquitetura onde GO esta sendo executado: ", runtime.GOARCH)

	var i uint16
	i = 65535
	fmt.Println(i)
	i++
	fmt.Println(i) // cuidado com overflow, pois como se fosse um odometro ele zera e reinicia.
	i++
	fmt.Println(i)
}

/*
- int vs. float: Números inteiros vs. números com frações.
- golang.org/ref/spec → numeric types
- Integers:
    - Números inteiros
    - int & uint → “implementation-specific sizes”
    - Todos os tipos numéricos são distintos, exceto:
        - byte = uint8
        - rune = int32 (UTF8)
        (O código fonte da linguagem Go é sempre em UTF-8).
    - Tipos são únicos
        - Go é uma linguagem estática
        - int e int32 não são a mesma coisa
        - Para "misturá-los" é necessário conversão
    - Regra geral: use somente int
- Floating point:
    - Números racionais ou reais
    - Regra geral: use somente float64
- Na prática:
    - Defaults com :=
    - Tipagem com var
    - Dá pra colocar número com vírgula em tipo int?
    - Overflow
	- Um uint16, por exemplo, vai de 0 a 65535.
	- Que acontece se a gente tentar usar 65536?
	- Ou se a gente estiver em 65535 e tentar adicionar mais 1?
*/
