# Integração com C e C++ (cgo)

Este tópico explica limites de portabilidade, custo de manutenção e critérios de uso.

## Fontes reintegradas

- `nao_tratados/go_com_c`
- `nao_tratados/cpp`

## Quando usar cgo

- biblioteca legada obrigatória;
- integração com hardware/SDK sem API Go;
- necessidade explícita de interoperabilidade.

## Riscos

- builds mais lentos;
- complexidade de toolchain;
- diferenças entre plataformas.
