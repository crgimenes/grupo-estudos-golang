# error

Go tem um sistema muito parecido com C de tratamento de erros.

```go
f, err := os.Open("filename.ext")
if err != nil {
    panic(err)
}
```

Podemos gerar nossos erros :smile:, usando o pacote nativo **errors**

```go
err := errors.New("Novo erro")
```

---
[Inicio](README.md)

[< interface](interface.md) - [goroutines >](goroutines.md)
