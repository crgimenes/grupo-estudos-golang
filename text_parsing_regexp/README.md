# Parsing textual e regex

Texto solto aparece em log, arquivo de configuração simples, saída de comando e payload legado. Antes de sair enfiando `strings.Split` em tudo, vale separar duas coisas: leitura linha a linha e extração do pedaço que interessa.

Este tópico usa só biblioteca padrão:

- `bufio.Scanner` para ler o texto uma linha por vez;
- `regexp` para reconhecer linhas no formato `chave=valor`;
- `strings` para limpar espaços e ignorar comentários simples;
- `slices` para imprimir o resultado em ordem estável.

## Pré-requisitos

- Go instalado.
- Conhecimento básico de mapas e funções.
- Noção de que regex serve bem para formato pequeno e previsível. Para gramática de verdade, use um parser de verdade. Regex não precisa apanhar por uma tarefa que não é dela.

## Exemplo

O exemplo lê um texto com linhas válidas, comentário e uma linha inválida:

```text
app=go-study
mode=dev
# ignored
invalid-line
workers=4
```

A função `ParseKeyValues` devolve um mapa com as linhas que casam com `chave=valor`. Linhas vazias, comentários e lixo ficam de fora.

```go
parsed, err := ParseKeyValues(data)
if err != nil {
    panic(err)
}
```

Mesmo lendo de uma `string`, o exemplo trata `scanner.Err()`. É hábito bom: quando a fonte muda para arquivo, pipe ou rede, o erro já tem caminho para aparecer.

## Rodar

```bash
go run .
```

Saída esperada:

```text
app=go-study
mode=dev
workers=4
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

## Pontos de atenção

`bufio.Scanner` é ótimo para linha pequena. O limite padrão de token é 64 KiB; se você pretende ler linha gigante, configure `scanner.Buffer` ou use `bufio.Reader`.

Regex também tem limite prático. O padrão deste tópico aceita chave com letras, números e `_`, e valor sem outro `=`. Isso é uma regra didática, não um formato universal.

## Erros comuns

- Usar `Split(line, "=")` e quebrar valor que contém `=` sem perceber.
- Ignorar `scanner.Err()` porque o exemplo começou lendo uma string.
- Fazer regex grande demais para compensar um formato mal definido.
- Depender da ordem de iteração do mapa. Mapa em Go não tem ordem estável; por isso o exemplo ordena as chaves antes de imprimir.

## Próximos passos

- Trocar a entrada de `string` por um arquivo aberto com `os.Open`.
- Aceitar espaços ao redor do `=` se o formato pedir isso.
- Comparar este exemplo com `encoding/csv` quando o dado tiver colunas de verdade.
