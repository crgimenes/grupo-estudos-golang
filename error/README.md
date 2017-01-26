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
## Testando erros


Quando estiver tratando erros não compare strings compare erros, a forma mais simples de fazer isso é instanciando uma variável para conter o erro no inicio package e usar essa instancia nas comparações e retornos.

```go
var ErrPanettoneDeChocolate = errors.New("Panettone tem que ser apenas com passas e frutas cristalizadas")
```
...
```go
err := p.VerificaPanettone()
if err == ErrPanettoneDeChocolate {
	panic(err)
}
```
O package *os* cria seu próprio objeto error para retornar não apenas uma string mas mais informações isso é possível porque error é uma interface e desde que você implemente uma função Error() que retorne uma string você pode fazer qualquer struct ser um objeto *error*.

```go
if _, err := os.Stat(nomeDoArquivo); err != nil {
    if os.IsNotExist(err) {
        return false
    }
}
```


---
[Inicio](../README.md)

[< interface](../interface/) - [goroutines >](../goroutines/)
