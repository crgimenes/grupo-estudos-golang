# CSV, TSV e dados tabulares

Arquivo separado por vírgula parece simples até aparecer campo com vírgula dentro, quebra de linha no meio de aspas ou TSV exportado de planilha. Fazer `strings.Split` direto funciona no exemplo de duas colunas e quebra quando o dado fica minimamente real.

Este tópico usa `encoding/csv`, da biblioteca padrão, para ler dados tabulares com delimitador configurável. O mesmo leitor cobre CSV e TSV: muda o `Comma`, não o parser inteiro.

## Objetivo

- Ler registros CSV e TSV sem escrever parser manual.
- Mostrar por que `FieldsPerRecord = -1` aceita linhas com tamanhos diferentes.
- Tratar erro de parsing em vez de devolver uma tabela meio lida como se estivesse tudo certo.

## Pré-requisitos

- Go instalado.
- Saber o básico de slices e strings.
- Ter visto tratamento de erros com `if err != nil`.

## Executar

```bash
go run .
```

Saída esperada:

```text
name | city | total
Ana | Santos | 42.50
Bruno | Recife | 19.90
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes cobrem leitura com vírgula, leitura com tab, erro de aspas quebradas e a formatação usada pelo `main`.

## Pontos de atenção

`encoding/csv` não é só para vírgula. Para TSV, configure `reader.Comma = '\t'`. Continuar chamando a função de `ParseCSV` quando ela também lê TSV vira ruído rápido; por isso o exemplo usa `ParseDelimited`.

`FieldsPerRecord = -1` deixa o leitor aceitar registros com quantidades diferentes de campos. Isso é útil para logs exportados e dados sujos, mas não valida contrato de schema. Se o arquivo precisa ter exatamente 3 colunas, cheque isso depois de ler.

O erro de `Read` precisa ser tratado dentro do loop. `io.EOF` é o fim normal do arquivo; qualquer outro erro é dado quebrado.

## Próximos passos

- Ler de um `os.File` em vez de uma string.
- Validar quantidade mínima de colunas por linha.
- Converter campos numéricos com `strconv.ParseFloat` depois do parsing.
