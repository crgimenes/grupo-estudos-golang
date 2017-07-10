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
type Localidade struct {
	X     int    `json:"valor_x"`
	Y     int    `json:"valor_y"`
	Nome  string `json:"nome_da_localidade"`
	valor int    // campo que inicia com letra minuscula é "private"
}
```

Você também pode adicionar etiquetas (tags) aos seus campos para gerar/ler arquivos JSON, para salvar/ler documentos do MongoDB, campos de banco de dados com sqlx. Esses recursos leem essas etiquetas e automaticamente fazem o binding/unbinding dos dados da Struct para esses estruturas externas a aplicação. A instrução *omitempty* é para que o binder ignore esse campo quando o valor dele for vazio

```go
type Localidade struct {
	X     int    `json:"valor_x" bson:"valor_x" db:"valorX"`
	Y     int    `json:"valor_y" bson:"valor_y" db:"valorY"`
	Nome  string `json:"nome_da_localidade,omitempty" bson:"nome_da_localidade,omitempty" db:"nomeDaLocalidade,omitempty"`
	valor int    // campo que inicia com letra minuscula é "private"
}
```

Você pode criar funções para atuar especificamente sobre sua struct como funções membro de uma classe.

```go
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
```
[Playground](https://play.golang.org/p/_TiiFcxGx_)

---
[Inicio](../README.md)

[< Funções](../funcoes/) - [Loop for >](../for/)
