# Goroutines e WaitGroup

`goroutine` é a unidade básica de concorrência em Go: você chama uma função com `go` e ela passa a rodar em paralelo com o fluxo atual. O problema aparece logo depois: se o `main` terminar antes dessas funções, o processo acaba e o trabalho some junto.

`sync.WaitGroup` resolve esse caso pequeno: contar tarefas, marcar cada término e esperar a contagem voltar para zero.

## Pré-requisitos

- saber declarar slices e funções;
- entender que concorrência não promete ordem de execução;
- ter Go 1.26 ou mais novo instalado.

## O exemplo

O programa calcula quadrados em goroutines separadas e guarda cada resultado no mesmo índice da entrada. Esse detalhe importa: a ordem em que as goroutines terminam varia, mas o slice final continua previsível.

```go
values := []int{2, 3, 4, 5}
fmt.Printf("in=%v out=%v\n", values, ConcurrentSquare(values))
```

Saída verificada com `go run .`:

```text
in=[2 3 4 5] out=[4 9 16 25]
```

## O que observar no código

- `wg.Add(len(values))` acontece antes do loop. Chamar `Add` dentro da goroutine cria uma corrida entre "comecei a esperar" e "avisei que existe trabalho".
- `defer wg.Done()` fica no começo da função anônima. Se alguém adicionar um `return` mais tarde, a contagem ainda fecha.
- O loop passa `i` e `v` como argumentos da função anônima. Fica explícito qual valor aquela goroutine deve usar.

## Executar

```bash
go run .
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

## Erros comuns

- esquecer `Done` e deixar `Wait` preso para sempre;
- chamar `Add` depois de iniciar a goroutine;
- escrever em estruturas compartilhadas sem separar os índices ou sem sincronização;
- esperar que a ordem de término das goroutines seja igual à ordem do loop.

Este exemplo não tenta limitar o número de goroutines. Para quatro itens isso é ótimo. Para 40 mil, o próximo assunto é worker pool ou semáforo com canal.
