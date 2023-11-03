# JSON Dicas e Truques

## Convertendo e validando JSON

É sempre bom saver se o JSON é valido antes, isso poupa tempo de debug procurando um problema que não esta no seu código, para isso uma boa ferramenta on-line é o [JSONLint](https://jsonlint.com)

Outra ferramenta on-line muito boa é o [JSON-to-Go](https://mholt.github.io/json-to-go/) para converter arquivos JSON em struct Go, ele não vai converter partes repetitivas em sub-structs mas ajuda muito, principalmente quando a estrutura é muito grande e complexa.

## Percorrendo o JSON com Recursão

Funções recursivas são deliciosas, são extremamente importantes para computação e entender o seu funcionamento pode não ser intuitivo então vale a pena dedicar algum tempo para estudando recursividade, um bom ponto de partida é o otimo video [Programming Loops vs Recursion do canal Computerphile](https://www.youtube.com/watch?v=HXNhEYqFo0o).

Vamos ver como carregar um arquivo JSON em uma map e dai percorre cada um dos elementos do mapa, muito útil quando se quer fazer algum processamento com os valores dos campos.

### Lendo um arquivo para uma variável

Usamos a função ReadFile do pacote ioutil para carregar ler o arquivo, a vantagem dessa função é que ela é muito pratica, mas você precisa tomar cuidado porque dependendo do tamanho do arquivo não é uma boa ideia carregar tudo para a RAM de uma vez.

```go
data, err := os.ReadFile("../payload.json")
if err != nil {
 fmt.Println(err)
 return
}
```

Agora temos o conteúdo do arquivo em um array de bytes na RAM, a variável `data`.

### Convertendo para map

Para converter o array de bytes para um map vamos usar a função `Unmarshal` do pacote `json`.

```go
payload := make(map[string]interface{})
err = json.Unmarshal(data, &payload)
if err != nil {
 fmt.Println(err)
 return
}
```

Nesse momento a variável payload contem todos os dados em um formato bastante conveniente, todos os nomes dos campos do JSON agora são as chaves no map e todos os objetos estão na interface, lembrando que é sempre bom tomar cuidado com interfaces vazias `interface{}`.

### Verificando o map

finalmente vamos chamar a função `chkMap` que vamos ver mais abaixo para verificar se algum valor esta abaixo do limite.

```go
if chkMap(payload) {
 fmt.Println("Um ou mais itens abaixo do limite")
}
fmt.Println("fim")
```

### As funções recursivas

A função `chkMap` vai percorrer cada campo do map, se algum dos objetos conter os campos `Limit` e `Value` comparamos os valores e dependendo do resultado já retornamos, caso contrário vamos para o loop que percorre todos os campos desse nível procurando por campos do tipo `map` ou do tipo `slice` e se encontrar chama a função adequada. Uma ressalva para esse código é que criamos uma função auxiliar para tornar o código mais legível, alem da `chkMap` que percorre mapas também criamos a `chkSlice` para percorrer os `slices`, dessa forma separamos essas duas estruturas de memória em funções especializadas. Seria perfeitamente possível escrever esse código em uma única função mas perderíamos a chance de demonstrar o uso de mais de uma função para recursividade.

```go
func chkMap(payload map[string]interface{}) (ret bool) {
 limit, lmtOk := payload["Limit"]
 value, valOk := payload["Value"]
 if lmtOk && valOk {
  if value.(float64) < limit.(float64) {
   ret = true
   return
  }
 }

 for _, v := range payload {
  switch v.(type) {
  case []interface{}:
   ret = chkSlice(v.([]interface{}))
  case map[string]interface{}:
   ret = chkMap(v.(map[string]interface{}))
  }
  if ret {
   return
  }
 }
 return
}

func chkSlice(pauload []interface{}) (ret bool) {
 for _, v := range pauload {
  switch v.(type) {
  case []interface{}:
   ret = chkSlice(v.([]interface{}))
  case map[string]interface{}:
   ret = chkMap(v.(map[string]interface{}))
  }
  if ret {
   return
  }
 }
 return
}
```

## Manipulando JSON com struct

Outra forma de tratar JSON que usamos bastante é converter para uma struct, manipular os dados como queremos e dai converter novamente para JSON.

### omitempty

Go fornece algumas ferramentas úteis como a tag `omitempty` que podemos usar para avisar o parser JSON que se o conteudo de um determinado campo estiver vazio ele deve ser omitido na hora de gerar o JSON.

```go
type metadata struct {
 SystemID  int    `json:"SystemID,omitempty"`
 FileID    string `json:"FileID,omitempty"`
 SubModule string `json:"SubModule,omitempty"`
}
```

Nesse exemplo sempre que o campo `SystemID` for zero, ou `FileID` for uma string vazia `""` ou então `SubModule` for uma string vazia o campo sera omitido na hora de gerar o JSON.

### Campos como ponteiros

Mais uma forma de tratar dados usando struct é transformando o campo em um ponteiro, basta colocar um asterisco `*` na frente do tipo do campo. Não se preocupe, se o `C` deixou você traumatizado com ponteiros, Go é muito mais gentil com isso.

Se um campo da nossa struct for um ponteiro e o valor for `nil` quando a função `Marshal` converter a struct para JSON esse campo vai aparecer como um campo `Null`, raramente é o que queremos, mas se adicionarmos a tag `omitempty` como vimos anteriormente esse campo vai desaparecer completamente.

E claro se você quiser retornar um objeto vazio é só no lugar de usar um ponteiro igualar o campo a uma instancia fazia, no caso de `metadata`
 seria `metadata{}`, e cada campo dentro da struct vai obedecer as suas tags como já vimos.

### Structs parciais

Muitas vezes não queremos todos os dados do JSON, alguns campos de alguma parte pode ser suficiente e não tem motivo para fazer [uma enorme struct como a do exemplo](https://github.com/go-br/estudos/blob/master/json/struct/main.go) se queremos apenas alguns dados.

Vamos supor por exemplo que queremos apenas o campo `metadata` do [JSON do nosso exemplo](https://github.com/go-br/estudos/blob/master/json/payload.json)

Não precisa declarar a struct inteira, para pegar apenas esse campo podemos declarar a seguinte struct.

```go
type apenasMetadata struct {
 Payload []struct {
  Result struct {
   Metadata metadata
  }
 }
}
```

Nesse exemplo declaramos apenas a parte da struct que queremos, precisa seguir o mesmo caminho dos campos superiores, mas diminuiu muito o tamanho da struct e o parser JSON quando fizer `Unmarshal` dos dados vai alegremente ignorar todo o resto e retornar apenas o que esta representado na struct.

## Structs dentro de funções

E uma dica para o uso de structs pequenas especializadas é que você nem precisa declarar ela fora da função que vai usar, muito útil para não ter estruturas perdidas pelo seu código.

No exemplo abaixo a struct `produtos` só existe dentro da função `structInetna`

```go
func structInetna() {
 type produtos struct {
  Nome  string
  Valor float64
    }
    ...
}
```

E se você for usar ela apenas uma vez nem mesmo é necessário declarar um novo tipo, você pode simplesmente já declarar e instanciar ao mesmo como no proximo exemplo.

```go
func structInetna() {
 produtos := struct {
  Nome  string
  Valor float64
    }{}
    ...
}
```

## Um erro comum com structs e JSON

Um erro muito comum quando estamos usando structs e usando o parser de JSON é que campos com letras minúsculas são `private` para o Go, apenas campos com letras maiúsculas são visíveis e isso se aplica também ao parser de JSON, então se os nomes dos campos da struct iniciarem com letra minúscula eles serão sumariamente ignorados pelo parser tanto para `Unmarshal` como para `Marshal`.

Caso você queira que o campo tenha letra minúscula quando gerar o JSON mude o nome na tag como no exemplo.

```go
type device struct {
 Limit int `json:"limit"`
 Value int `json:"value"`
```

Note que o campo na struct tem letra maiúscula o que faz com que ele seja visível para o parser mas na tag json tem letra minúscula ou seja condo for convertido de e para array de bytes vai usar letras minúsculas.

## Links úteis

- [Código fonte](https://github.com/go-br/estudos/tree/master/json)
- [Repositório do nosso grupo](https://github.com/go-br/estudos)
- [E você encontra mais exemplos aqui](https://github.com/go-br)
- [Pagina do grupo de estudos](https://gopher.pro.br)
