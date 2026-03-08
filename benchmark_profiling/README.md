# Benchmarks e profiling

Este tópico conecta benchmark e profiling a decisões de otimização.

## Teste e benchmark

```bash
go test ./...
go test -bench=. -benchmem ./...
```

## Perfil de CPU

```bash
go test -bench=. -cpuprofile cpu.out ./...
go tool pprof cpu.out
```
