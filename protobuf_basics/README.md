# Protobuf: introdução moderna

Este tópico cobre geração de código, contratos e quando usar Protobuf no lugar de JSON.

## Fonte reintegrada

- `nao_tratados/protobuf/e1/user/user.proto`
- `nao_tratados/protobuf/e1/user/user.pb.go`

## Fluxo típico

```bash
protoc --go_out=. --go_opt=paths=source_relative user.proto
```

## Quando preferir Protobuf

- contratos rígidos entre serviços;
- payloads menores e serialização eficiente;
- evolução de schema com compatibilidade.
