# Go com Lua

Este tópico consolida scripting embarcado com Lua para extensão controlada de aplicações.

## Objetivo

Usar Lua para configuração e regras dinâmicas sem recompilar a aplicação Go.

## Estrutura

- `example1/`: leitura de configuração via script.
- `example2/`: integração simples de runtime.
- `launcher/`: execução controlada com script externo.

## Dependência

```bash
go get github.com/yuin/gopher-lua
```

## Boas práticas

- expor apenas APIs necessárias ao script;
- validar entradas vindas do script;
- limitar operações críticas fora do runtime Go.
