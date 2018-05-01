# godoc

Go tem uma ferramenta muito poderosa para visualizar documentação.

Exemplos:

```sh
godoc fmt
godoc github.com/crgimenes/rotateString
```

Você pode facilmente exportar a documentação em formato html:

```sh
godoc -html github.com/crgimenes/rotateString > rotateString.html
```

Ou ainda subir a documentação toda em um servidor html.

```sh
godoc -http=:6060
```
