package main

import (
	"fmt"
//	"sync"
)

type pessoa struct {
	nome string
	idade int
	sexo string
}

func (p *pessoa) falar() {
	//var wg sync.WaitGroup
	fmt.Println("Olá meu nome é ", p.nome, ", e eu tenho ", p.idade, " anos, de idade, prazer te conhecer")
	//wg.Done()
}

type humanos interface {
	falar()
}

func dizerAlgo(h humanos) {
	h.falar()
}

func main() {
	p1 := pessoa{"João", 20, "M"}
	p2 := pessoa{"Maria", 25, "F"}

	(&p1).falar()
	(&p2).falar()

	dizerAlgo(&p1)
	dizerAlgo(&p2)
}
