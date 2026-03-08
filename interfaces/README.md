# Interfaces

Este tópico cobre satisfação implícita, interfaces pequenas, `any`, type assertion, type switch e interfaces padrão.

## Executar

```bash
go run .
```

## Testar

```bash
go test ./...
```

## Pontos-chave

- Em Go, tipo satisfaz interface por método compatível.
- Prefira interfaces pequenas e orientadas a comportamento.
- `any` é alias de `interface{}`.
- Type switch ajuda a tratar diferentes implementações com segurança.
