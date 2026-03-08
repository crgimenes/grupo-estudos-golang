# Padrão mínimo para tópicos

Este documento define o padrão obrigatório para novos tópicos do repositório.

## Objetivo

Garantir consistência de estrutura, clareza didática e facilidade de manutenção.

## Estrutura recomendada

Cada tópico deve usar um diretório em inglês, minúsculo, com underscore quando necessário.

Exemplo:

```text
topic_name/
  README.md
  go.mod              # quando o tópico for executado de forma isolada
  main.go             # quando houver exemplo único
  main_test.go        # quando houver comportamento testável
  examples/           # opcional, quando houver múltiplos exemplos
```

## Regras obrigatórias

- `README.md` em português do Brasil.
- Código-fonte, comentários e nomes técnicos em inglês.
- Exemplos pequenos, completos e executáveis.
- Tratamento explícito de erros quando aplicável.
- Sem dependências externas sem justificativa didática clara.

## Requisitos de conteúdo do README

Cada tópico deve incluir:

- objetivo do tópico;
- pré-requisitos;
- comandos completos de execução;
- pontos de atenção e erros comuns;
- próximos passos sugeridos.

## Critério para sair de `nao_tratados/`

Um conteúdo antigo só pode ser removido de `nao_tratados/` quando:

- o material equivalente já existir em um tópico navegável;
- o novo README cobrir os conceitos principais do conteúdo antigo;
- os exemplos do novo tópico compilarem;
- os links para o novo tópico estiverem no `README.md` principal e no `src/_indice.md`.

## Validação obrigatória para tópicos com código Go

Execute no diretório do tópico:

```sh
go fmt ./...
go fix ./...
go vet ./...
go test -timeout 30s -count 1 ./...
```
