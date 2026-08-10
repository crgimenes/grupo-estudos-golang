# Defer, panic e recover

`defer` é útil para fechar arquivo, liberar lock e registrar uma ação que precisa acontecer quando a função sair. O detalhe que pega muita gente é a ordem: vários `defer` rodam em pilha, o último registrado roda primeiro.

`panic` e `recover` entram em outro lugar. `panic` não é substituto para `error`; é para estado impossível, bug de programação ou uma borda que precisa derrubar a execução daquele fluxo. `recover` só funciona dentro de uma função chamada por `defer`, então ele deve ficar perto da fronteira que sabe transformar o pânico em erro.

## Objetivo

Este tópico mostra três coisas pequenas:

- a ordem real de execução de `defer`;
- como um `panic` interrompe a função atual;
- como `recover` captura esse pânico e devolve um `error` para o chamador.

O exemplo usa divisão por zero de propósito. Em código normal, o melhor seria validar `b == 0` e retornar erro direto. Aqui o pânico existe para deixar `recover` visível sem inventar uma falha maior.

## Pré-requisitos

- Go instalado.
- Saber rodar comandos básicos com `go run` e `go test`.
- Já ter visto funções que retornam `error`.

## Executar

```bash
go run .
```

Saída esperada:

```text
10/2 = 5
recovered: recovered panic: division by zero
defer order: start,first,second,third
```

A linha do `defer order` mostra o ponto principal: o corpo da função adiciona `start`, depois os `defer` rodam de trás para frente.

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes conferem se `safeDivide(10, 0)` vira erro depois do `recover` e se a ordem dos `defer` continua exata.

## Pontos de atenção

`recover` chamado fora de um `defer` não captura nada útil. Também não adianta recuperar um pânico e seguir como se nada tivesse acontecido; se o programa perdeu consistência, devolver um erro claro é o mínimo.

Outro detalhe: `defer` executa no retorno da função onde foi registrado, não no fim do bloco. Se você registrar `defer file.Close()` dentro de um loop grande, pode segurar muitos arquivos abertos até a função inteira terminar.

## Próximos passos

- Trocar a divisão por um exemplo com arquivo e `defer Close`.
- Testar a ordem dos `defer` com três funções que imprimem nomes diferentes.
- Comparar esse exemplo com tratamento explícito de erro, sem `panic`.
