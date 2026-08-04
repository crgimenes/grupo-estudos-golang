# Sincronização e race conditions

Um contador compartilhado entre goroutines parece inofensivo até o primeiro `go test -race`. `n++` não é uma operação atômica: lê, soma e grava. Sem coordenação, duas goroutines podem ler o mesmo valor e uma atualização some no caminho.

Este tópico usa `sync.Mutex` para proteger uma seção crítica pequena e `sync.WaitGroup` para esperar as goroutines terminarem. Também mostra `sync.Map`, mas sem vender milagre: ele serve para caches e mapas com acesso concorrente específico, não para substituir `map` em todo código.

## Pré-requisitos

- Go 1.26 ou mais novo.
- Saber criar goroutines.
- Ter visto `go test` pelo menos uma vez.

## Executar

```bash
go run .
```

Saída esperada:

```text
counter=100
cache[go]=2
```

O número `100` importa. Se o contador fosse incrementado sem o `Mutex`, o teste poderia passar em uma máquina e falhar em outra. Race condition tem esse gosto ruim: o bug existe mesmo quando a execução do dia resolveu colaborar.

## Testar

```bash
go test ./...
go test -race ./...
```

Use `go test -race` quando mexer em código concorrente. Ele deixa a execução mais lenta, mas encontra acessos compartilhados sem sincronização que uma asserção comum não enxerga.

## Pontos de atenção

- Proteja o menor trecho possível com `Lock` e `Unlock`. Segurar o `Mutex` durante I/O é pedir fila à toa.
- Use `defer Unlock()` quando a função tem retorno ou erro no meio. No incremento de uma linha, destravar logo abaixo é mais direto.
- `sync.Map` não preserva ordem. Se precisar imprimir ou comparar valores em ordem estável, copie as chaves para um slice e ordene.
- `WaitGroup` espera trabalho terminar; ele não protege memória compartilhada. São ferramentas diferentes.

## Próximos passos

Troque o `Mutex` por `sync.RWMutex` em uma versão com muitas leituras e poucas escritas. Depois rode com `go test -race` de novo; se o exemplo não passa no race detector, ele não está pronto para ensinar concorrência.
