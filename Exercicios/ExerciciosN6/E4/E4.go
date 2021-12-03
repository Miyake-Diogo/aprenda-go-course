package main

import (
	"fmt"
)
// criação de structs
type comida struct {
	nome string
	preco float64
	tipo string
}
// criação dos métodos getNome, getPreco e getTipo
func (nomeComida comida) getNome() string {
	return nomeComida.nome
}

func (precoComida comida) getPreco() float64 {
	return precoComida.preco
}

func (tipoComida comida) getTipo() string {
	return tipoComida.tipo
}
// função principal
func main() {
	// criação de variáveis
	vegana := comida{
		nome : "curry",
		tipo: "indiana",
		preco: 10.0,
	}
	fmt.Println("O Nome da comida é: ", vegana.getNome())
	fmt.Println("O Preço da comida é: R$", vegana.getPreco())
	fmt.Println("O Tipo da comida é: ", vegana.getTipo())
}
/*
- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
- Crie um método para "pessoa" que demonstre o nome completo e a idade;
- Crie um valor de tipo "pessoa";
- Utilize o método criado para demonstrar esse valor.
*/