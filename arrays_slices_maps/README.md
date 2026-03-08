# Arrays, slices e mapas

Este tópico destaca diferenças de uso, cópia, compartilhamento de memória e armadilhas comuns.

## Executar

```bash
go run .
```

## Testar

```bash
go test ./...
```

## Pontos-chave

- Array tem tamanho fixo no tipo (`[3]int`).
- Slice referencia um array subjacente.
- `copy` evita compartilhamento acidental de memória.
- Mapas retornam zero value para chave ausente.
