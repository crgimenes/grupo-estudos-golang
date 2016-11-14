# error

Go tem um sistema muito parecido com C de tratamento de erros.

```go
f, err := os.Open("filename.ext")
if err != nil {
    panic(err)
}
```
[Playground](https://play.golang.org/p/5hEzgU5pvy)

Podemos gerar nossos erros :smile:, usando o pacote nativo **errors**

```go
err := errors.New("Novo erro")
```

---
[Inicio](../README.md)

[< interface](../interface/) - [goroutines >](../goroutines/)
