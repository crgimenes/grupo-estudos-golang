package main

import (
	"encoding/json"
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
}

type Localidade struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_da_localidade"`
	valor int    // campo que inicia com letra minuscula é "private"
}

func (p *Localidade) Soma(l Localidade) {
	p.X += l.X
	p.Y += l.Y
}

func main() {
	// Declarando e atribuindo valores usando os nomes dos campos
	aluno := Aluno{
		Nome:  "Cesar",
		Idade: 42,
	}

	fmt.Println("Aluno:", aluno)

	// Declarando e atribuindo valores diretamente
	minhaCasa := Localidade{3, 4, "casa", 500}

	// Atribuindo uma estrutura vazia
	estruturaVazia := Aluno{}

	// Mais uma vez declarando e atribuindo usando os nomes dos campos
	escolaLocalidade := Localidade{
		Y:    100,
		X:    200,
		Nome: "escola",
	}

	// Alocando e em seguida atribuindo valores
	var outraLocalidade Localidade
	outraLocalidade.X = 10
	outraLocalidade.Y = 20
	outraLocalidade.Nome = "trabalho"

	fmt.Println("minha casa:", minhaCasa)
	fmt.Println("outra localidade:", outraLocalidade)
	fmt.Println("escola:", escolaLocalidade)
	fmt.Println("aluno:", aluno)
	fmt.Printf("estrutura vazia: %q\r\n", estruturaVazia)

	minhaCasa.Soma(outraLocalidade)

	fmt.Println("localidade minha casa somada com outra localidade", minhaCasa)

	fmt.Printf("minha casa: %v\r\n", minhaCasa)
	fmt.Printf("minha casa: %+v\r\n", minhaCasa)

	// Brincando com JSON
	j, err := json.Marshal(minhaCasa)
	if err != nil {
		panic(err)
	}

	fmt.Println("minha casa json:", string(j))

	novaLocalidade := Localidade{}
	err = json.Unmarshal(j, &novaLocalidade)
	if err != nil {
		panic(err)
	}

	fmt.Println("pondo depois do Unmarshal:", novaLocalidade)

}
