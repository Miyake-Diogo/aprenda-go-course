package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())

	wg.Add(2)

	go func1()
	go func2()

	fmt.Println(runtime.NumGoroutine())

	wg.Wait()

}

func func1() {
	for i := 0; i < 100; i++ {
		fmt.Println("func1:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 100; i++ {
		fmt.Println("func2:", i)
		time.Sleep(20)
	}
	wg.Done()
}

/*
- Concorrência é quando abre uma padaria do lado da outra e as duas quebram :)
- Fun facts: 
    - O primeiro CPU dual core "popular" veio em 2006
    - Em 2007 o Google começou a criar a linguagem Go para utilizar essa vantagem
    - Go foi a primeira linguagem criada com multi-cores em mente
    - C, C++, C#, Java, JavaScript, Python, etc., foram todas criadas antes de 2006
    - Ou seja, Go tem uma abordagem única (fácil!) para este tópico
- E qual a diferença entre concorrência e paralelismo?

Concorrência é lidar com várias coisas ao mesmo tempo. 
Código concorrente não necessariamente é paralelo.
O paralelismo vai ser algo realizado pelo compilador

Concorrência é sobre a execução sequencial e disputada de um conjunto de tarefas independentes. 
Sob o ponto de vista de um sistema operacional, o responsável por esse gerenciamento é o escalonador de processos. 
Já sob o ponto de vista de concorrência em uma linguagem de programação como Go, por exemplo, 
o responsável é o scheduler interno da linguagem. 
Escalonadores preemptivos (como é o caso dos sistemas operacionais modernos) favorecem a concorrência pausando 
e resumindo tarefas (no caso de sistemas operacionais estamos falando de processos e threads 
no que chamamos de trocas de contexto) para que todas tenham a oportunidade de serem executadas.

Paralelidmo é fazer várias coisas ao mesmo tempo.
Paralelismo é sobre a execução paralela de tarefas, ou seja, mais de uma por vez (de forma simultânea), 
a depender da quantidade de núcleos (cores) do processador. 
Quanto mais núcleos, mais tarefas paralelas podem ser executadas. 
É uma forma de distribuir processamento em mais de um núcleo.
*/

/*
- Goroutines!
- O que são goroutines? São "threads."
- O que são threads? [WP](https://pt.wikipedia.org/wiki/Thread_...))
- Na prática: go func.
- Exemplo: código termina antes da go func executar.
- Ou seja, precisamos de uma maneira pra "sincronizar" isso.
- Ah, mas então... não.
- Qualé então? sync.WaitGroup:
- Um WaitGroup serve para esperar que uma coleção de goroutines termine sua execução.
    - func Add: "Quantas goroutines?"
    - func Done: "Deu!"
    - func Wait: "Espera todo mundo terminar."
- Ah, mas então... sim!
- Só pra ver: runtime.NumCPU() & runtime.NumGoroutine()
*/