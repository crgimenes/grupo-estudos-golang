# Observabilidade básica

Log não é decoração. Quando um serviço dá problema às 2h da manhã, uma linha como `erro ao processar requisição` não ajuda ninguém. O mínimo útil é registrar o evento com campos que sobrevivem a grep, filtro por JSON e painel: rota, status e duração.

Este tópico usa `log/slog`, que entrou na biblioteca padrão no Go 1.21. A ideia aqui é pequena de propósito: produzir um log estruturado e testar os campos que importam. Métrica, tracing distribuído e OpenTelemetry ficam fora; são assuntos próprios e pesam demais para este primeiro passo.

## Objetivo

- criar um logger JSON com `log/slog`;
- registrar uma requisição com campos nomeados;
- manter a saída do exemplo determinística para teste e documentação.

## Pré-requisitos

- Go instalado;
- noções básicas de `go test` e mapas JSON.

## Executar

```bash
go run .
```

Saída esperada:

```text
{"level":"INFO","msg":"request","path":"/health","status":200,"duration_ms":3}
```

O exemplo remove o campo `time` do handler. Em serviço real o timestamp costuma ser obrigatório; aqui ele atrapalharia porque mudaria a cada execução e deixaria teste e README brigando com o relógio.

## Testar

```bash
go test ./...
```

O teste decodifica o JSON e confere os campos um por um. É melhor do que procurar uma substring: se `status` virar string sem querer, o teste quebra onde deve quebrar.

## Pontos de atenção

- Use nomes de campos estáveis. `duration_ms` é chato, mas é fácil de filtrar.
- Não coloque segredo em log. Token em log estruturado continua sendo token vazado.
- Evite montar mensagem gigante quando campos resolvem. `slog` existe para isso.

## Próximos passos

- adicionar `request_id` quando houver uma borda HTTP real;
- comparar `slog.TextHandler` e `slog.JSONHandler`;
- estudar métricas separadamente com `expvar` ou outro tópico dedicado.
