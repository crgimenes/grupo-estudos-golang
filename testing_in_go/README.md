# Testes em Go

Este tópico mostra três peças que aparecem cedo em teste Go: teste em tabela, subtest com `t.Run` e `t.Helper` para tirar ruído da linha de erro.

O exemplo é pequeno de propósito. `NormalizeSpace` recebe uma string, troca qualquer sequência de whitespace por um espaço ASCII e remove sobra no começo e no fim. Isso dá casos bons para testar sem montar servidor, banco ou fixture.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Terminal aberto neste diretório: `testing_in_go`.

## O que observar

O teste fica em `package space_test`, não no mesmo pacote do código. Isso força o exemplo a usar a API pública, do mesmo jeito que outro pacote usaria.

`assertEqual` chama `t.Helper()`. Sem isso, uma falha aponta para dentro do helper. Com isso, ela aponta para a linha do subtest que chamou o helper. Parece detalhe pequeno, mas em tabela com 20 casos economiza tempo.

O `ExampleNormalizeSpace` também é teste. O `go test` executa o exemplo e compara a saída com o comentário `// Output:`. Se a saída mudar, o teste quebra.

## Executar os testes

```bash
go test -timeout 30s -count 1 ./...
```

Saída esperada:

```text
ok  	example.com/testing-in-go	...
```

O tempo no fim varia de máquina para máquina.

## Ver os subtests

```bash
go test -v ./...
```

Você deve ver nomes como:

```text
=== RUN   TestNormalizeSpace/extra_spaces
=== RUN   TestNormalizeSpace/tabs_and_newlines
```

Esses nomes vêm do campo `name` da tabela. Quando um caso quebra, o nome do caso quebrado aparece no relatório.

## Erros comuns

- Esquecer `t.Helper()` em helpers de asserção e acabar depurando a linha errada.
- Usar `t.Fatal` dentro de goroutine criada pelo teste. O teste pode continuar de um jeito que não parece óbvio para iniciante.
- Testar só o happy path. Aqui o caso com tab, newline e espaços nas pontas existe para pegar exatamente a parte chata da função.

## O que este tópico não cobre

Não entra em mock, teste HTTP, benchmark nem fuzzing. Cada um desses assuntos merece um exemplo próprio; misturar tudo aqui deixaria o primeiro contato com `testing` mais confuso do que útil.
