# Compressão e arquivamento

Este tópico separa duas ideias que costumam aparecer juntas, mas não são a mesma coisa.

`compress/gzip` comprime um fluxo de bytes. Ele não guarda nome de arquivo, diretório ou lista de entradas. `archive/zip` cria um arquivo com entradas nomeadas; cada entrada pode ter seu próprio conteúdo comprimido. Se você só quer mandar um payload menor pela rede, gzip basta. Se precisa empacotar `relatorio.csv`, `log.txt` e `metadata.json` no mesmo arquivo, aí é zip.

O exemplo fica em memória para não esconder o assunto atrás de caminho de arquivo. O payload é pequeno de propósito: `go study group\n`, 15 bytes. Depois do cabeçalho do gzip, ele vira 39 bytes. Compressão de texto minúsculo pode crescer. Isso não é bug, é overhead.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Terminal no diretório `compression_archiving`.
- Noção básica de `[]byte` e tratamento de erro.

## Rodar

```bash
go run .
```

Saída esperada:

```text
gzip roundtrip: "go study group\n"
gzip bytes: 15 -> 39
zip files: [note.txt]
```

A primeira linha confirma o roundtrip: comprimir e descomprimir devolve os mesmos bytes. A segunda mostra o tamanho antes e depois do gzip. A terceira lista o nome salvo dentro do zip.

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes não tentam provar que gzip "comprime bem". Eles travam o comportamento didático: gzip precisa voltar ao payload original, e o zip precisa carregar a entrada `note.txt`.

## Pontos de atenção

- Feche `gzip.Writer` e `zip.Writer`. Sem `Close`, bytes finais podem ficar no buffer e o arquivo sai inválido.
- Zip guarda nomes. Nunca extraia um nome vindo de fora direto no disco sem validar caminho; `../../arquivo` é uma armadilha velha e ainda aparece.
- Gzip não é criptografia. Quem lê o arquivo comprimido também lê o conteúdo.
- Para arquivo grande, não monte tudo em `bytes.Buffer` como aqui. Use `io.Copy` entre arquivos ou streams.

## Próximos passos

- Ler um arquivo com `os.Open` e gravar o `.gz` com `os.Create`.
- Criar um zip com mais de uma entrada.
- Medir com benchmark a diferença entre payload pequeno e payload repetitivo grande.
