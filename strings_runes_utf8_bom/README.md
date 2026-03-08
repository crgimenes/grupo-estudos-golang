# Strings, runas, UTF-8 e BOM

Este tópico cobre diferença entre bytes e runas, iteração segura com `range`, e detecção de BOM UTF-8.

## Executar

```bash
go run .
```

## Testar

```bash
go test ./...
```

## Pontos-chave

- `len(string)` mede bytes, não caracteres.
- `range` percorre runas e fornece índice em bytes.
- `utf8.RuneCountInString` conta runas.
- BOM UTF-8 (`EF BB BF`) pode aparecer no início de arquivos e causar parsing inesperado.
