package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

const quantidadeGoroutines = 100

func criaGoroutines (quantidade int) {
	wg.Add(quantidade)
	for j := 0; j < quantidade; j++ {
		go func() {
			valor := contador
			runtime.Gosched()
			valor++
			contador = valor
			wg.Done()
		}()
	}
}

func main() {
	criaGoroutines(quantidadeGoroutines)
	wg.Wait()
	fmt.Println("Quantidade do contador: ", contador)
	fmt.Println("Quantidade Goroutines: ", quantidadeGoroutines)
}