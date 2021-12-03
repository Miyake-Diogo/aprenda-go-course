package main

import (
	"fmt"
)

func main() {
	example1()


}

func example1(){
	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)
	
	for i := 0; i < x; i++ {
		select {
			case v := <-a:
				fmt.Println("Canal A:", v)			
			case v := <-b:
				fmt.Println("Canal B:", v)
		}
	}

}

func example2(){	canal := make(chan int)
	quit := make(chan int)
	go recebeQuit(canal, quit)
	enviaPraCanal(canal, quit)
}

func recebeQuit (canal chan int, quit chan int) {
	for i := 0; i < 50; i++ {
		fmt.Println("Recebido:", <-canal)
	}
	quit <- 0
}

func enviaPraCanal (canal chan int, quit chan int) {
	qualquercoisa := 1
	for {
		select {
			case canal <- qualquercoisa:
				qualquercoisa++
			case <-quit:
				return
		}
	}
}
//        - Chans par, ímpar, quit
//        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//        - Func receive é um select entre os três canais, encerra no quit
func example3(){	par := make(chan int)
	ímpar := make(chan int)
	quit := make(chan bool)

	go mandaNúmeros(par, ímpar, quit)

	receive(par, ímpar, quit)
}

func mandaNúmeros(par, ímpar chan int, quit chan bool) {
	total := 100
	for i := 0; i < total; i++ {
		if i % 2 == 0 {
			par <- i
		} else {
			ímpar <- i
		}
	}
	close(par)
	close(ímpar)
	quit <- true
}

func receive(par, ímpar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("O número", v, "é par.")
		case v := <-ímpar:
			fmt.Println("O número", v, "é ímpar.")
		case <-quit:
			return
		}
	}
}

/*
- Select é como switch, só que pra canais, e não é sequencial.
- "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
- Na prática:
    - Exemplo 1:
        - Duas go funcs enviando X/2 numeros cada uma pra um canal
        - For loop X valores, select case ←x
    - Exemplo 2:
        - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
        - Func 2 for infinito, select: case envia pra canal, case recebe de quit
    - Exemplo 3:
        - Chans par, ímpar, quit
        - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
        - Func receive é um select entre os três canais, encerra no quit
        - Problema!
*/