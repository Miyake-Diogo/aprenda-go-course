package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var contador int32

const quantidadeGoroutines = 100

func criaGoroutines (quantidade int) {
	wg.Add(quantidade)
	for j := 0; j < quantidade; j++ {
		go func() {
			atomic.AddInt32(&contador, 1)
			v := atomic.LoadInt32(&contador)
			runtime.Gosched()
			fmt.Println(v)
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