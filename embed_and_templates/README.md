# Embed e templates

Este tópico junta duas coisas que aparecem cedo em aplicação web pequena: guardar arquivos dentro do binário com `embed` e renderizar texto com `text/template` ou HTML com `html/template`.

O ponto principal não é fazer um site. É enxergar uma diferença que pega gente distraída: `text/template` só substitui texto; `html/template` entende contexto de HTML e escapa valores antes de escrever no documento.

## Pré-requisitos

- Go instalado.
- Noção básica de funções, strings e tratamento de erro.
- Um terminal aberto neste diretório.

## O exemplo

O arquivo `assets/hello.txt` entra no binário por causa desta diretiva:

```go
//go:embed assets/*.txt
var assets embed.FS
```

A partir daí o programa pode ler `assets/hello.txt` sem depender do arquivo solto no diretório de execução. Isso é útil para templates, páginas estáticas pequenas, mensagens padrão e fixtures didáticas.

O exemplo também renderiza a mesma ideia por dois caminhos:

- `RenderText` usa `text/template` para montar texto comum.
- `RenderHTML` usa `html/template` para escrever HTML com escape automático.

Quando a entrada é `<script>alert(1)</script>`, o HTML gerado não injeta uma tag real. O resultado impresso é este:

```text
Hello, gopher
<p>&lt;script&gt;alert(1)&lt;/script&gt;</p>
```

## Executar

```bash
go run .
```

## Testar

```bash
go test -timeout 30s -count 1 ./...
```

Os testes cobrem o escape de conteúdo HTML e um caso específico de URL insegura. `html/template` troca `javascript:...` por `#ZgotmplZ` quando percebe que o valor foi parar em um `href`. Esse marcador é feio de propósito: é melhor quebrar o link do que entregar um HTML perigoso com cara de válido.

## Pontos de atenção

`embed` resolve arquivos em tempo de compilação. Se você mudar `assets/hello.txt`, precisa rodar `go run .` ou `go test` de novo para compilar outro binário.

`html/template` não valida regra de negócio. Ele ajuda na fronteira de saída para HTML, mas não decide se uma URL deveria ser aceita, se um usuário pode ver uma página, ou se o dado faz sentido.

Evite usar `template.HTML` para “desligar” escape. Quase sempre é um atalho ruim. Só use quando o HTML já foi produzido por uma parte confiável do próprio programa e essa decisão estiver clara no código.

## Próximos passos

- Trocar o template inline por um arquivo `.html` dentro de `assets/`.
- Comparar `template.ParseFS` com `template.ParseFiles`.
- Servir um arquivo embutido com `http.FileServer(http.FS(...))` em um exemplo separado.
