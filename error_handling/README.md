# Tratamento de erros

Erro em Go é valor. Essa frase parece mantra até cair no primeiro `if err != nil` copiado sem pensar. O ponto deste tópico é separar três casos que aparecem em código real: erro conhecido, erro com contexto e erro com dado estruturado.

O exemplo usa só biblioteca padrão:

- `errors.Is` para reconhecer um erro sentinela mesmo depois de `%w`;
- `errors.As` para extrair um tipo de erro com campos úteis;
- `fmt.Errorf` com `%w` para preservar a causa original.

## Pré-requisitos

Você precisa saber declarar funções, retornar múltiplos valores e ler um `if` simples. Não precisa de goroutine, interface própria ou framework de logging.

## Rodar

```bash
go run .
```

Saída esperada:

```text
id=1 -> user=gopher
id=2 -> not found
id=0 -> validation error on id
age="25" -> 25
age="-1" -> validation error: age: must be non-negative
age="abc" -> parse age "abc": strconv.Atoi: parsing "abc": invalid syntax
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes não tentam comparar a string inteira do erro. Isso é de propósito. A mensagem pode mudar quando você adiciona contexto, mas o contrato que importa aqui é outro: `errors.Is` ainda encontra `ErrUserNotFound`, e `errors.As` ainda encontra `ValidationError` com o campo certo.

## Pontos de atenção

Não use `==` para comparar erro embrulhado com `%w`; use `errors.Is`. O `==` só vê o valor do topo, então quebra exatamente quando você começa a escrever mensagens melhores.

Não transforme todo erro em string cedo demais. Se uma função faz `return fmt.Errorf("%v", err)`, ela perdeu a cadeia de erro. Quem chama não consegue mais usar `errors.Is` nem `errors.As`.

Também não crie erro sentinela para tudo. `ErrUserNotFound` faz sentido porque o chamador pode tomar uma decisão específica. Para validação, um tipo com `Field` e `Msg` carrega mais informação do que meia dúzia de variáveis globais.

## Próximos passos

Depois deste exemplo, vale testar `errors.Join` quando uma operação pode falhar em mais de um lugar. Use com calma: juntar erro só para evitar escolher uma causa principal costuma deixar o chamador com mais trabalho, não menos.
