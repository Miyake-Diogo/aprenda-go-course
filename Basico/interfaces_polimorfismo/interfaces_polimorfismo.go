package main
import (
	"fmt"
)
type Especie struct {
	Nome string
	Raca string

}
type Animal interface {
	comida(c string)
	dormir()
	rugido()

}
func (e Especie) comida(c string) {
	fmt.Println("A comida é", c)
}
func (e Especie) dormir() {
	fmt.Println("O animal está dormindo")
}
func (e Especie) rugido() {
	if e.Nome == "Cachorro" {
		fmt.Println("AUAUAUAU")
	} else {
		fmt.Println("Miauuu")
	}
}

func main() {
	cachorro := Especie{
		Nome: "Cachorro",
		Raca: "Vira-lata",
	}
	gato := Especie{
		Nome: "Gato",
		Raca: "Persa",
	}
	cachorro.comida("Carne")
	gato.comida("Ração")
	cachorro.dormir()
	gato.dormir()
	cachorro.rugido()
	gato.rugido()
	


}
/*

- Em Go, valores podem ter mais que um tipo.
- Uma interface permite que um valor tenha mais que um tipo.
- Declaração: keyword identifier type → type x interface
- Após declarar a interface, deve-se definir os métodos necessários para implementar essa interface.
- Se um tipo possuir todos os métodos necessários (que, no caso da interface{}, pode ser nenhum) então esse tipo implicitamente implementa a interface.
- Esse tipo será o seu tipo *e também* o tipo da interface.

- Exemplos:
    - Os tipos *profissão1* e *profissão2* contem o tipo *pessoa*
    - Cada um tem seu método *oibomdia()*, e podem dar oi utilizando *pessoa.oibomdia()*
    - Implementam a interface *gente*
    - Ambos podem acessar o função *serhumano()* que chama o método *oibomdia()* de cada *gente*
    - Tambem podemos no método *serhumano()* tomar ações diferentes dependendo do tipo:
        switch pessoa.(type) { case profissão1: fmt.Println(h.(profissão1).valorquesóexisteemprofissão1) [...] }* 
*/