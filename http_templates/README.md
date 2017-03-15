# http

### Uso de Templates!

Templates em Go são utilizados para gerar websites com conteúdo dinâmicos. 

Iremos neste exemplo, criar um webserver que irá manipular requisições web, apresentando um cabeçalho padrão e respondendo a solicitações de páginas html

Para utilização de templates devemos importar o pacote html/template

```go
package main

import (
	"fmt"
	"net/http"
	"html/template"
)

type Page struct {
	Name string
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	
		p := Page{Name : "GopherBrasil"}
		
		if name := r.FormValue("name"); name != ""{
			p.Name = name;
		}
	
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
	})
	
	fmt.Println(http.ListenAndServe(":8080", nil))

}
```

Neste nosso exemplo, devemos criar um diretório chamado "templates" e dentro dele um arquivo html "index.html" com o seguinte conteúdo:
```
<html>
	<head>
		<body>
			Hello, {{.Name}}
		</body>
	</head>
</html>
```

Criamos um objeto para representar os dados que queremos exibir na página
```go
type Page struct {
	Name string
}
```


Utilizamos um manipulador da biblioteca html/template para realizar o parse do nosso template (templates/index.html)
```go
templates := template.Must(template.ParseFiles("templates/index.html"))
```

Ainda dentro da função principal iremos criar um handler para tratar os eventos de request e response
```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	
		//Criamos um objeto que será passado como parâmetro para nosso template
		p := Page{Name : "GopherBrasil"}
		
		//Associamos um elemento de formulário a este objeto. Se o name não estiver vazio, atualiza a página com o valor estipulado
		if name := r.FormValue("name"); name != ""{
			p.Name = name;
		}
	
		//Dentro do nosso manipulador vamos tratar os erros relacionados ao template
		//ExecuteTemplate irá executar diretamente nosso template 
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			//Alerta o usuário para os erros encontrados através da mensagem e código padrão do erro
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
	})
```

Realizando o setup para ouvir na porta 8080
```go
fmt.Println(http.ListenAndServe(":8080", nil))
```

E por fim note que podemos utilizar a passagem de parâmetros diretamente via url
http://localhost:8080/?name=Edward

---
[Inicio](../README.md)

[< Cliente Http](../http_get/) - [enviando e-mail via SMTP >](../smtp/)
