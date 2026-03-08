# SQLite e PostgreSQL: diferenças práticas

Este tópico compara preparo de ambiente, DSN e características de uso.

## SQLite

- ideal para aplicações locais e embarcadas;
- banco em arquivo;
- setup simples para estudos.

Exemplo de abertura:

```go
sql.Open("sqlite", "file:app.db")
```

## PostgreSQL

- indicado para aplicações distribuídas e concorrentes;
- servidor dedicado;
- controle avançado de transações e extensões.

Exemplo de DSN:

```text
postgres://user:pass@localhost:5432/app?sslmode=disable
```

## Comparação objetiva

- SQLite: simplicidade operacional.
- PostgreSQL: robustez para produção.
