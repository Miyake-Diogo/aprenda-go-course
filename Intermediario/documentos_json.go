package main

import (
	"encoding/json"
	"fmt"
	"os"
)
type pessoa struct
{
	Nome string
	Sobrenome string
	Idade int
	Profissao string
	ContaBancaria float64
}

// para Unmarshal

type informacoes struct {
	Nome          string  `json:"Nome"`
	Sobrenome     string  `json:"Sobrenome"`
	Idade         int     `json:"Idade"`
	Profissao     string  `json:"Profissao"`
	Contabancaria float64 `json:"Contabancaria"`
}

func main() {
	// Marshal precisa de variável para receber os dados
	fmt.Println("Marshal")
	jamesBond := pessoa{
		Nome: "James",
		Sobrenome: "Bond",
		Idade: 32,
		Profissao: "Agente Secreto",
		ContaBancaria: 350000.00,

	}
	darthVader := pessoa{
		Nome: "Anakin",
		Sobrenome: "Skywalker",
		Idade: 60,
		Profissao: "Sith",
		ContaBancaria: 4890000000.00,
	}
	j, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println("Erro ao gerar o JSON:", err)
	}
	d, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println("Erro ao gerar o JSON:", err)
	}
	fmt.Println("Todas infos: ", string(j))
	fmt.Println("Todas infos: ", string(d))

	// UnMarshal
	fmt.Println("UnMarshal")
	sb := []byte(`{"Nome":"James","Sobrenome":"Bond","Idade":40,"Profissao":"Agente Secreto","Contabancaria":1000000.5}`)
	var jamesbond informacoes

	error2 := json.Unmarshal(sb, &jamesbond)
	if err != nil {
		fmt.Println("error:", error2)
	}

	fmt.Println("Informações: ", jamesbond)
	fmt.Println("Profissão: ", jamesbond.Profissao)
	
	// encoder vai dretamente
	fmt.Println("Encoder")
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesBond)

}