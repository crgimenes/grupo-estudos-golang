# Struct

Structs são coleções de campos.

```go
type Aluno struct {
	Nome  string
	Idade int
}
```

Você também pode "decorar" sua declaração para tornar mais fácil converter para JSON por exemplo.

```go
type Ponto struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_do_ponto"`
	valor int    // campo que inicia com letra minuscula é "private"
}
```

Você pode criar funções para atuar especificamente sobre sua struct como funções membro de uma classe.

```go
type Ponto struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_do_ponto"`
	valor int    // campo que inicia com letra minuscula é "private"
}

func (p *Ponto) Soma(ponto Ponto) {
	p.X += ponto.X
	p.Y += ponto.Y
}
```

## Exemplo

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
}

type Ponto struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_do_ponto"`
	valor int    // campo que inicia com letra minuscula é "private"
}

func (p *Ponto) Soma(ponto Ponto) {
	p.X += ponto.X
	p.Y += ponto.Y
}

func main() {
	// Declarando e atribuindo valores usando os nomes dos campos
	aluno := Aluno{
		Nome:  "Cesar",
		Idade: 42,
	}

	fmt.Println("Aluno:", aluno)

	// Declarando e atribuindo valores diretamente
	meuPonto := Ponto{3, 4, "casa", 500}

	// Atribuindo uma estrutura vazia
	estruturaVazia := Aluno{}

	// Mais uma vez declarando e atribuindo usando os nomes dos campos
	escolaPonto := Ponto{
		Y:    100,
		X:    200,
		Nome: "escola",
	}

	// Alocando e em seguida atribuindo valores
	var outroPonto Ponto
	outroPonto.X = 10
	outroPonto.Y = 20
	outroPonto.Nome = "trabalho"

	fmt.Println("meu ponto:", meuPonto)
	fmt.Println("outro ponto:", outroPonto)
	fmt.Println("escola:", escolaPonto)
	fmt.Println("aluno:", aluno)
	fmt.Printf("estrutura vazia: %q\r\n", estruturaVazia)

	meuPonto.Soma(outroPonto)

	fmt.Println("meu ponto somado com outro ponto", meuPonto)

	fmt.Printf("meu ponto: %v\r\n", meuPonto)
	fmt.Printf("meu ponto: %+v\r\n", meuPonto)

	// Brincando com JSON
	j, err := json.Marshal(meuPonto)
	if err != nil {
		panic(err)
	}

	fmt.Println("meu ponto json:", string(j))

	novoPonto := Ponto{}
	err = json.Unmarshal(j, &novoPonto)
	if err != nil {
		panic(err)
	}

	fmt.Println("pondo depois do Unmarshal:", novoPonto)

}
```

---
[Inicio](README.md)

[< Variáveis](variaveis.md) - [Funções >](funcoes.md)
