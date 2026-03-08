# Variáveis, constantes e tipos

Este tópico apresenta declaração curta, inferência de tipo, zero values, conversão explícita, `iota` e visibilidade de identificadores.

## Executar

```bash
go run .
```

Saída esperada:

```text
short=10 inferred=20.5 zeroInt=0 converted=20 level=info
```

## Testar

```bash
go test ./...
```

## Pontos-chave

- `:=` declara e inicializa variáveis locais.
- Zero value depende do tipo (`int` vira `0`, `string` vira `""`).
- Conversão numérica em Go é explícita (`int(x)`).
- `iota` ajuda a criar enums simples com tipo nomeado.
- Identificadores com inicial maiúscula são exportados.
