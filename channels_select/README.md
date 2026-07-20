# Canais e select

Este exemplo mostra uma situação pequena e comum: duas goroutines podem responder, mas o programa só precisa da primeira resposta. O `select` espera o primeiro canal que ficar pronto e também dá uma saída de segurança com timeout.

O detalhe que vale reparar não está só no `select`. Os canais `fast` e `slow` têm buffer de 1. Sem esse buffer, se o timeout ganhasse antes de uma goroutine enviar o valor, ela poderia ficar presa tentando escrever para um canal que ninguém mais lê. Para um exemplo de estudo isso parece pequeno; em serviço rodando por dias, esse tipo de vazamento vira fila de goroutine parada.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Noção básica de goroutine e canal.

## Executar

```bash
go run .
```

Saída esperada:

```text
fast
```

O retorno é `fast` porque essa goroutine dorme 10ms, enquanto a outra dorme 50ms. Se o timeout passado para `FirstSignal` for menor que o primeiro envio, o resultado passa a ser `timeout`.

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes cobrem dois caminhos: a primeira resposta vencendo e o timeout vencendo. Eles ainda usam tempos pequenos, então servem bem para enxergar o comportamento, mas não são um modelo para teste pesado de concorrência.

## Pontos de atenção

- `select` não consulta os casos em ordem; quando mais de um canal está pronto, a escolha é pseudoaleatória.
- `time.After` é ótimo para exemplos curtos. Em loops ou caminhos quentes, prefira `time.NewTimer` e pare o timer quando ele não for mais usado.
- Canal sem buffer exige alguém lendo no mesmo ritmo de quem escreve. Buffer pequeno pode ser uma decisão consciente, não um remendo.

## Próximos passos

Troque os tempos em `main.go` e rode os testes de novo. Depois, tente remover o buffer dos canais e pense no que acontece quando o timeout vence antes dos envios.
