# Ponteiros, escopo e shadowing

Este tópico explica mutação por ponteiro, passagem por valor e bug comum de shadowing com `:=`.

## Executar

```bash
go run .
```

## Testar

```bash
go test ./...
```

## Pontos-chave

- Go sempre passa argumentos por valor.
- Para mutar o dado original, passe ponteiro.
- `:=` em escopo interno pode esconder variável externa.
- Shadowing é fonte comum de bugs em tratamento de erro.
