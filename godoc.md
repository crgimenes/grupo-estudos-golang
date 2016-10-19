# godoc

Go tem uma ferramenta muito poderosa para visualizar documentação.

Exemplos:

```
godoc fmt
godoc github.com/crgimenes/rotateString
```

Você pode facilmente exportar a documentação em formato html:

```
godoc -html github.com/crgimenes/rotateString > rotateString.html
```

Ou ainda subir a documentação toda em um servidor html.

```
godoc -http=:6060
```

[Início](README.md)

[< time](time.md) - [http >](http.md)
