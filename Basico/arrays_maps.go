package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	// arrays
	fmt.Println("Arrays")
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))

	// slices
	fmt.Println("Arrays vs Slices")

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	//slice[20] = 1 // nao vai pois extrapola os indices
	//fmt.Println(slice[20])

	// Slices Ranges
	fmt.Println("Slices Ranges")

	novaslice := []string{"banana", "maçã", "jaca", "pêssego"}

	//for índice, valor := range novaslice {fmt.Println("No índice", índice, "temos o valor:", valor)}

	//slice[4] = "melancia"
	novaslice = append(novaslice, "melancia")

	for _, valor := range slice {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

	novaslice2 := []int{20, 21, 22, 23, 1, 13}

	total := 0

	for _, valor := range novaslice2 {

		// mesma coisa que total = total + valor
		total += valor

	}

	fmt.Println("O valor total é:", total)

	// Fatiando Slices
	fmt.Println("Fatiando Slices")

	//			    1.           2.           3.         4.                5.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores[2:4]

	fmt.Println(fatia)

	sabores = append(sabores[:2], sabores[4:]...)

	fmt.Println(sabores)

	// Exercicio

	fmt.Println(sabores[0], sabores[1], sabores[2])

	for i := 0; i < len(sabores); i++ {

		fmt.Println(sabores[i])
	}

	//slice anexando outra slice
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12, 13}
	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...)
	fmt.Println(umaslice)

	// Make
	fmt.Println("Make")
	sliceMake := make([]int, 5, 10)

	sliceMake[0], sliceMake[1], sliceMake[2], sliceMake[3], sliceMake[4] = 1, 2, 3, 4, 5

	sliceMake = append(sliceMake, 6)
	sliceMake = append(sliceMake, 7)
	sliceMake = append(sliceMake, 8)
	sliceMake = append(sliceMake, 9)
	sliceMake = append(sliceMake, 10)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	sliceMake = append(sliceMake, 10)

	fmt.Println(sliceMake, len(sliceMake), cap(sliceMake))

	// Slices Multidimensionais
	fmt.Println("Slices Multidimensionais")
	ss := [][]int{
		// Índice: 0  1  2                   // Índice:
		[]int{1, 2, 3, 4, 5, 6},       // 0
		[]int{7, 8, 9, 10, 11, 12},    // 1
		[]int{13, 14, 15, 16, 17, 18}, // 2
	}
	fmt.Println(ss[2][4])
	fmt.Println(ss[0][:])
	
	// Arrays Subjacente
	fmt.Println("Arrays Subjacente")
	primeiroslice := []int{1, 2, 3, 4, 5}
	
	fmt.Println(primeiroslice)
	
	segundoslice := append(primeiroslice[:2], primeiroslice[4:]...)

	fmt.Println(segundoslice)

	fmt.Println(primeiroslice)	
	
	// Maps
	fmt.Println("Maps")
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 444444

	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])


	// comma ok idiom
	if será, ok := amigos["fantasma"]; !ok {
		fmt.Println("não tem!")
	} else {
		fmt.Println(será)
	}
	// maps range
		qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "esse é massa",
		19:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	soma := 0

	for key, _ := range qualquercoisa {
		soma += key
	}
	
	fmt.Println(soma)
		
// maps delete 

	fmt.Println(qualquercoisa)

	delete(qualquercoisa, 123)

	fmt.Println(qualquercoisa)	

}

/*
- Estruturas de dados, ou agrupamentos de dados, nos permitem agrupar valores diferentes. Estes valores podem ser ou não do mesmo tipo.
- As estruturas que veremos são: arrays, slices, structs e maps.
- Vamos começar com arrays. Arrays em Go são uma fundação, e não algo que utilizamos todo dia.
- Seu tamanho deve estar presente na declaração: var x [n]int
- Atribui-se valores a suas posições com: x[i] = y (0-based)
- Para ver o tamanho usa-se: len(x)
- ref/spec: "The length is part of the array's type" → [5]int != [6]int
- Effective Go: Arrays são úteis para [umas coisas que a gente não vai fazer nunca] e servem de fundação para slices. Use slices ao invés de arrays.

- O que são tipos de dados compostos?
    - Wikipedia: Composite_data_type
    - Effective Go: Composite literals
    - ref/spec: Composite literals
- Uma slice agrupa valores de um único tipo.
- Criando uma slice: literal composta → x := []type{values}

- Slices:
    - Tamanho: len(x)
    - Índice específico: x[i] (0-based)
- Para ver todos os itens de uma slice utilizamos o loop for com range.
- Range significa alcance, faixa, extensão.
- For range: for i, v := range x {}

- x[:], x[a:], x[:b], x[a:b]
- "a" é incluso;
- "b" não é.
- Exemplo: cabeça magnética de um disco rígido (relógio, fita).
    - Off-by-one error.

 É fatiando que se deleta um item de uma slice. Na prática:
    - x := append(x[:i], x[:i]...)

- Effective Go: append (package builtin)
- x = append(slice, ...values)
- x = append(slice, slice...)
- Todd: unfurl → desdobrar, desenrolar
- Nome oficial: enumeration


- Slices são feitas de arrays.
- Elas são dinâmicas, podem mudar de tamanho.
- Sempre que isso acontece, um novo array é criado e os dados são copiados.
- É conveniente, mas tem um custo computacional.
- Para otimizar as coisas, podemos utilizar make.
- make([]T, len, cap)
- "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
- len(x), cap(x)
- x[n] onde n é maior que len é out of range. Use append.
- Append maior que cap modifica o array subjacente.
- pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
- Effective Go.

- Slices multi-dimensionais são slices que contem slices.
- São como planilhas.
- [][]type

- Isso tudo aqui a gente já viu:
- Toda slice tem um array subjacente.
- Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
- Exemplo:
    - x := []int{...números}
    - y := append(x[:i], x[:i]...)
    - pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
    - Ou seja, y utiliza o mesmo array subjacente que x.
    - O que nos dá um resultado inesperado.
- Ou seja, bom saber de antemão pra não ter que aprender na marra.


- Utiliza o formato key:value.
- E.g. nome e telefone
- Performance excelente para lookups.
- map[key]value{ key: value }
- Acesso: m[key]
- Key sem value retorna zero. Isso pode trazer problemas.
- Para verificar: comma ok idiom.
    - v, ok := m[key]
    - ok é um boolean, true/false
- Na prática: if v, ok := m[key]; ok { }
- Para adicionar um item: m[v] = value
- Maps *não tem ordem.*


- Range: for k, v := range map { }
- Reiterando: maps *não tem ordem* e um range usará uma ordem aleatória.

- delete(map, key)
- Deletar uma key não-existente não retorna erros!


*/
