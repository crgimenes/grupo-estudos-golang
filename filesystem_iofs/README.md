# Diretórios e `io/fs`

`io/fs` é o ponto em que o código para de depender de "arquivo no disco" e passa a depender de um filesystem. Parece detalhe, mas muda o desenho do exemplo: a mesma função consegue ler um diretório real com `os.DirFS`, arquivos embutidos com `embed.FS` ou um filesystem falso no teste com `fstest.MapFS`.

Este tópico usa `fs.WalkDir` para percorrer uma árvore, `fs.Stat` para ler metadados e `fstest.MapFS` para deixar o exemplo pequeno. Nada de criar arquivo temporário só para provar que a lógica funciona.

## Objetivo

- Percorrer diretórios com `fs.WalkDir`.
- Listar apenas arquivos regulares, ignorando diretórios.
- Ler tamanho de arquivo por `fs.Stat` sem acoplar a função a `os.Stat`.

## Pré-requisitos

- Go 1.26 ou mais recente.
- Saber o básico de slices e tratamento de erros.
- Ter visto interfaces, porque `fs.FS` é uma interface pequena.

## Executar

```bash
go run .
```

Saída esperada:

```text
assets/logo.txt 2 bytes
docs/readme.md 6 bytes
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes usam `testing/fstest.MapFS` para montar uma árvore em memória com `a.txt`, `notes/` e `notes/b.txt`. Isso cobre o caso que costuma dar bug em exemplo apressado: `WalkDir` visita diretórios também, então filtrar por `IsDir` ou por arquivo regular não é opcional.

## Pontos de atenção

`fs.WalkDir` trabalha com caminhos separados por `/`, mesmo no Windows. Não misture isso com `filepath.Join` dentro de uma função que recebe `fs.FS`; `path` e os caminhos do próprio pacote `io/fs` são o encaixe certo ali.

`os.DirFS(".")` é útil para ligar a abstração ao disco, mas ele não prende o código dentro do diretório se alguém passar caminho com `..`. Para ferramenta local didática, tudo bem. Para limite de segurança, não trate `DirFS` como sandbox.

`DirEntry.Type()` nem sempre tem todos os bits de tipo preenchidos. Neste exemplo, `IsRegular` funciona com `fstest.MapFS` e com arquivos comuns. Se o código precisar seguir symlink, inspecionar permissão ou diferenciar dispositivos, chame `Info()` e lide com o erro.

## Próximos passos

- Trocar `fstest.MapFS` por `os.DirFS(".")` e listar arquivos do diretório atual.
- Usar `embed.FS` para empacotar arquivos de exemplo no binário.
- Adicionar um filtro por extensão usando `path.Ext`, não `filepath.Ext`, quando o dado vem de `fs.FS`.
