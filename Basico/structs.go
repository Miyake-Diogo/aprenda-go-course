package main

import (
	"fmt"
)

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

// Structs embutidos

type pessoa struct {
	nome  string
	idade int
}

type profissao struct {
	pessoa
	area    string
	cargo   string
	salario float32
}

func main() {
	cliente1 := cliente{
		nome:      "Jorge",
		sobrenome: "Capivara",
		fumante:   false,
	}

	cliente2 := cliente{"Tomé", "Souza", true}
	cliente3 := cliente{"Tulio", "Balde", false}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
	fmt.Println(cliente3)

	//Structs Embutidos
	pessoa1 := pessoa{
		nome:  "Marcelo",
		idade: 30,
	}

	pessoa2 := profissao{
		pessoa: pessoa{
			nome:  "Bruce",
			idade: 25,
		},
		area:    "Heroi",
		cargo:   "Batman",
		salario: 115000000.99,
	}
	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	pessoa3 := pessoa{"Mauricio", 40}
	pessoa4 := profissao{pessoa{"Vanderlei", 50}, "Política", "Deputado",10000000.66}

	fmt.Println(pessoa1.nome)
	fmt.Println(pessoa2.pessoa.nome)
	fmt.Println(pessoa3.nome)
	fmt.Println(pessoa4)

	//Structs anonimos
	x := struct {
		nome  string
		idade int
	}{nome:  "Mario",
		idade: 50}

	fmt.Println(x)

}

/*
- Struct é um tipo de dados composto que nos permite armazenar valores de tipos *diferentes.*
- Seu nome vem de "structure," ou estrutura.
- Declaração: type x struct { y: z }
- Acesso: x.y
- Exemplo: nome, idade, fumante.

Structs Embutidos
- Structs dentro de structs dentro de structs.
- Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) *e tambem* um competidor (nome, equipe, pontos).


- É importante se familiarizar com a documentação da linguagem Go.
- Neste vídeo vamos ver um pouco sobre o que a documentação diz sobre structs.
- Veremos:
    - ref/spec
        - Já vimos mais da metade dos tipos em Go!
        - Struct types.
            - x, y int
            - anonymous fields
            - promoted fields


- São structs sem identificadores.
- x := struct { name type }{ name: value }


*/
