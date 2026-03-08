# Práticas de segurança em Go

Este tópico consolida privilégios mínimos, validação de entrada e gestão de segredos.

## Princípios

- executar com menor privilégio possível;
- validar toda entrada externa;
- usar timeout/cancelamento em I/O de rede;
- proteger segredos com variáveis de ambiente e rotação.

## Prática de privilégios

`syscall.Setuid` e `syscall.Setgid` ajudam a reduzir impacto em caso de exploração.
