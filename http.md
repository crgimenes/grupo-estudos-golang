# http

Pilhas inclusas!

Go vem com um conjunto enorme de bibliotecas padrão e mais milhares de bibliotecas criadas pela comunidade.

por exemplo para criar um servidor http basta o seguinte:


```
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
	http.ListenAndServe(":8080", nil)
}
```

---
[Inicio](README.md)
