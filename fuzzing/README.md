# Fuzzing

Este tópico mostra fuzz test com `go test -fuzz` para entradas arbitrárias.

## Testar

```bash
go test ./...
```

## Rodar fuzzing

```bash
go test -fuzz=FuzzParseCSVLine -run=^$ ./...
```
