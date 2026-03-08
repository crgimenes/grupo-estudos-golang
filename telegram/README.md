# Integrações com APIs externas

Este tópico evolui o diretório `telegram` para práticas gerais de clientes HTTP reutilizáveis.

## Regras de segurança

- usar variáveis de ambiente para segredos;
- nunca logar token completo;
- aplicar timeout e retry com backoff;
- tratar status de erro e limite de rate.

## Fluxo recomendado

1. criar `http.Client` com timeout;
2. encapsular chamadas em funções pequenas;
3. validar payload de entrada e saída;
4. mapear erros de rede e resposta HTTP.

## Fontes

- `main.go`
- `help.md`
- `context.txt`
