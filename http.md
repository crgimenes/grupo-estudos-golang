# http

### Pilhas inclusas!

Go vem com um conjunto enorme de bibliotecas padrão e mais milhares de bibliotecas criadas pela comunidade.

por exemplo para criar um servidor http basta o seguinte:


```go
package main

import (
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Olá mundo HTTP!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9999", nil)
}
```

### Servindo arquivos

Esse é um exemplo bem simples de como servir arquivos. Nesse caso o diretório *assets* contem os arquivos que queremos servir.

```go
package main

import (
  "log"
  "net/http"
)

func main() {
  assets := http.FileServer(http.Dir("assets/"))
  http.Handle("/", assets)

  log.Fatal(http.ListenAndServe(":9999", nil))
}
```

### Servindo arquivos e adicionando cabeçalho

Esse exemplo mostra como adicionar cabeçalhos na resposta do servidor antes de entregar os arquivos. Isso é importante porque você pode manipular o cache no cliente ou mesmo modificar completamente como a requisição vai responder.

Como bonus adicionamos dois parâmetros para configurar a porta e o diretório com os assets.

```go
package main

import (
	"flag"
	"log"
	"net/http"
)

func setHSTSHeader(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		h.ServeHTTP(w, r)
	}
}

func main() {

	var port = flag.String("port", "9999", "Define what TCP port to bind to")
	var root = flag.String("root", "assets/", "Define the root filesystem path")

	flag.Parse()

	assets := setHSTSHeader(http.FileServer(http.Dir(*root)))
	http.Handle("/", assets)

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
```

Esse ultimo exemplo fica melhor se compilarmos com:

```
go build http3.go
```

E então execute com:

```
./http3 --help
```





---
[Inicio](README.md)

[< godoc](godoc.md) - [tratando sinais >](signals.md)
