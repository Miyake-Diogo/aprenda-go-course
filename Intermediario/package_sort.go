package main

import (
	"fmt"
	"sort"
)


func simpleStringSort(valor []string) []string {
	sort.Strings(valor)
	return valor
}

func simpleNumberSort(number []int) []int {
	sort.Ints(number)
	return number
}

type carro struct {
	nome     string
	potencia int
	consumo  int
}

// Sort Customized 
type ordenarPorPotencia []carro
// Type interface
func (x ordenarPorPotencia) Len() int           { return len(x) }
func (x ordenarPorPotencia) Less(i, j int) bool { return x[i].potencia < x[j].potencia }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorConsumo []carro

func (x ordenarPorConsumo) Len() int           { return len(x) }
func (x ordenarPorConsumo) Less(i, j int) bool { return x[i].consumo > x[j].consumo }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ordenarPorLucroProDonoDoPosto []carro

func (x ordenarPorLucroProDonoDoPosto) Len() int           { return len(x) }
func (x ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return x[i].consumo < x[j].consumo }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main(){
	si := []int{123, 987, 324, 876, 234, 987, 234, 76}
	fmt.Println("Antes do Sort", si)
	fmt.Println("Após do Sort",simpleNumberSort(si))
	

	ss := []string{"z", "x", "p", "e", "b", "d", "a", "h"}
	fmt.Println("Antes do Sort", ss)
	fmt.Println("Após do Sort",simpleStringSort(ss))
	
	// Custom Sort
	carros := []carro{
		carro{"Chevete", 50, 12},
		carro{"Porsche", 300, 5},
		carro{"Fusca", 20, 16},
	}

	fmt.Println("Inicial: Nome do Carro | Potência | Consumo \n", carros)

	sort.Sort(ordenarPorPotencia(carros))

	fmt.Println("Potência:\n", carros)

	sort.Sort(ordenarPorConsumo(carros))

	fmt.Println("Consumo:\n", carros)

	sort.Sort(ordenarPorLucroProDonoDoPosto(carros))

	fmt.Println("Lucro:\n", carros)


}


/*

- O sort que eu quero não existe. Quero fazer o meu.
- Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
    - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
- Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
- E aí posso fazer do jeito que eu quiser.
- Exemplo:
    - struct carros: nome, consumo, potencia
    - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
    - tipo ordenarPorPotencia
    - tipo ordenarPorConsumo
    - tipo ordenarPorLucroProDonoDoPosto
*/