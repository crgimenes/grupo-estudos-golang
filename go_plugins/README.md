# Sistema de plugins em Go

Este tópico cobre pacote `plugin`, limitações e alternativas.

## Limitações importantes

- disponível em Linux/macOS;
- exige `-buildmode=plugin`;
- acoplamento de versões entre host e plugin.

## Build básico

```bash
go build -buildmode=plugin -o greeter.so ./plugin
```

## Quando evitar

Prefira interfaces + injeção de dependência ou scripting (Lua) quando portabilidade for prioridade.
