# http Uso do Get

Dentro do pacote HTTP, Go fornece o método Get que permite facilmente sua utilização para consumo de requisições web.

No exemplo abaixo podemos ver o uso do método Get para uma requisição de um determinado feed de noticias:


```go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {

	// Obtém o feed de rss através da url
	resp, err := http.Get("http://www.techworld.com/news/rss")
	
	if err != nil{
		s := err.Error()
		fmt.Printf("%q\n", s)	
		return
	}
	
	// Sinaliza que a última ação a ser feita no programa é o fechamento da resposta
	defer resp.Body.Close()
	
	// Verifica se o código de status é 200, indicando assim o sucesso da solicitação
	if resp.StatusCode != 200 {
		println(resp.StatusCode)	
		return
	}

	// Por fim escreve o conteúdo do feed de rss
	bodyBytes, err2 := ioutil.ReadAll(resp.Body) 
	
	if err2 != nil{
		s := err2.Error()
		fmt.Printf("%q\n", s)	
		return
	}
	
    bodyString := string(bodyBytes) 
	println(bodyString)

}


```

Quando o método Get retorna, recebemos um ponteiro para um valor do tipo Response. 
```go
resp, err := http.Get("")
```

Após verificar se houveram erros, o método Close da resposta deve ser escalonado.
```go
defer resp.Body.Close()
```

Também realizamos a verificação no atributo StatusCode para saber se recebemos o código 200 associado ao sucesso da requisição.
```go
if resp.StatusCode != 200
```

Em nosso exemplo, qualquer valor diferente de 200 deve ser tratado como um erro. Se o valor não for 200, retornamos o código resultante da requisição.
```go
println(resp.StatusCode)
```

Após esta verificação, fazemos a leitura da resposta em um processo de conversão byte para string com o auxílio da biblioteca ioutil.
```go
bodyBytes, err2 := ioutil.ReadAll(resp.Body)
bodyString := string(bodyBytes)
```
