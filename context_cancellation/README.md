# Context

`context.Context` é o jeito normal de avisar uma goroutine que o trabalho perdeu o sentido: a requisição caiu, o prazo acabou, ou o chamador decidiu cancelar. Sem esse fio de cancelamento, código concorrente continua rodando no escuro e você só percebe quando o teste fica lento ou o processo começa a juntar goroutine parada.

Este exemplo mostra um trabalho pequeno que termina sozinho ou para quando o contexto fecha.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Noções básicas de goroutines, canais e `select`.

## Executar

```bash
go run .
```

Saída esperada:

```text
work finished without cancellation
canceled with: context deadline exceeded
```

## O que observar no código

`Work` recebe o contexto como primeiro argumento. Essa convenção parece detalhe, mas ajuda bastante quando a função cresce: quem chama enxerga logo que ela pode ser cancelada.

O `select` espera duas coisas: o timer do trabalho ou `ctx.Done()`. Quando o prazo de 5ms vence antes do trabalho de 50ms, a função devolve `context deadline exceeded`. O erro não é embrulhado aqui de propósito, porque `errors.Is(err, context.DeadlineExceeded)` precisa continuar funcionando no teste.

Também tem uma escolha pequena no timer: o exemplo usa `time.NewTimer` com `defer timer.Stop()` em vez de `time.After`. Para um programa minúsculo os dois funcionam. Em código que roda em loop, parar o timer deixa a intenção mais clara e evita carregar trabalho morto até o runtime limpar.

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes cobrem o caminho cancelado por deadline e o caminho que termina antes do contexto fechar.

## Erros comuns

- Criar um contexto novo dentro da função em vez de receber o contexto do chamador. Isso corta a propagação do cancelamento.
- Ignorar `ctx.Err()` e devolver `nil` no cancelamento. O chamador perde a diferença entre trabalho concluído e trabalho abandonado.
- Guardar `context.Context` em struct para reutilizar depois. Contexto carrega prazo e cancelamento de uma operação específica, não configuração global.

## Próximos passos

- Trocar o timer por uma chamada HTTP com `http.NewRequestWithContext`.
- Combinar `context.WithCancel` com uma goroutine produtora e conferir se ela para quando o consumidor desiste.
