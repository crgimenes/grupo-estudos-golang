# build tags

Build tags é a forma do go fazer compilação condicional ou seja mudar a forma como o código é compilado dependendo da tag que foi passada para o build.

Com esse recurso você pode por exemplo fazer com que o compilador use código especifico para a plataforma destino, como um código especifico para Windows e outro especifico para Linux e outro especifico para macOS. 

Você também pode criar suas próprias tags e também pode usar criar condições booleanas por exemplo colocando `!` (_not_) antes da tag.

## exemplos

Esse arquivo sera ignorado na compilação:

```go
// +build ignore

...
```

---

Esse arquivo sera compilado apenas quando a tag `minhatag` for passada para o comando _build_.

```go
// +build minhatag

...
func algumaFuncao() {
    ...
}
```

E esse próximo arquivo só sera compilado quando a tag `outratag` for passada para o _build_.

```go
// +build outratag

...
func algumaFuncao() {
    ...
}
```

---

Para passar tags para o comando build você faz assim:

```console
go build -tags 'minhatag'
```

---

Exite muito mais coisas interessantes sobre tags e outros recursos do build, para aprender mais de uma olhada nos links abaixo.

- [Documentação do pacote build](https://golang.org/pkg/go/build/)
- [O comando go](https://golang.org/cmd/go/)
