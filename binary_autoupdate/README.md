# Automação e atualização de binários

Este tópico cobre release, versionamento, distribuição e riscos de autoatualização.

## Riscos principais

- atualização interrompida pode corromper binário;
- assinatura/verificação de origem é obrigatória em produção;
- rollback precisa ser planejado.

## Testar

```bash
go test ./...
```
