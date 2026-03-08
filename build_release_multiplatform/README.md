# Build, release e multiplataforma

Este tópico cobre `GOOS`, `GOARCH` e organização de artefatos.

## Exemplos

```bash
GOOS=linux GOARCH=amd64 go build -o dist/app-linux-amd64 ./cmd/app
GOOS=windows GOARCH=amd64 go build -o dist/app-windows-amd64.exe ./cmd/app
GOOS=darwin GOARCH=arm64 go build -o dist/app-darwin-arm64 ./cmd/app
```

## Boas práticas

- nomear artefatos por plataforma;
- registrar versão e commit no binário;
- assinar release quando aplicável.
