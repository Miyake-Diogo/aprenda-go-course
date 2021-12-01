package main

import (
	"fmt"
)

func main() {
	x := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "Doce",
		ondetem: []string{"Holanda", "Shopping", "Na casa do tio rico"},
		vaibemcom: map[string]string{
			"Café da manhã": "Chá",
			"Almoço":        "Cafezinho"},
	}

	fmt.Println(x)


	y := struct {
		nome      string
		sabor     string
		ondetem   []string
		vaibemcom map[string]string
	}{
		nome:    "Stroopwafel",
		sabor:   "doce",
		ondetem: []string{"na Holanda", "na casa do seu tio rico"},
		vaibemcom: map[string]string{
			"no café da manhã": "chá",
			"no almoço":        "cafezinho",
			"na janta":         "não vai bem pq comer doce a noite engorda",
		},
	}
	
	fmt.Println(y)

}
/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/