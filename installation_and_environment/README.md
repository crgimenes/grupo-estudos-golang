# Instalação e ambiente do Go

Este tópico mostra como instalar o Go, configurar o ambiente e validar a instalação em macOS, Linux e Windows.

## Objetivo

Ao final, você deve conseguir:

- instalar o Go na sua plataforma;
- configurar o PATH corretamente;
- validar o ambiente com `go version` e `go env`;
- executar comandos básicos no terminal.

## Pré-requisitos

- acesso ao terminal (PowerShell, Terminal, zsh ou bash);
- permissão para instalar software na máquina.

## Download oficial

Use sempre o site oficial:

- [go.dev/dl](https://go.dev/dl/)

## Instalação por sistema operacional

### Windows

1. Baixe o instalador `.msi` em [go.dev/dl](https://go.dev/dl/).
2. Execute o instalador.
3. Verifique se `C:\Go\bin` está no PATH.

Comando para validar no PowerShell:

```powershell
go version
go env GOROOT GOPATH
```

### macOS

Opção recomendada - instalador oficial:

1. Baixe o pacote em [go.dev/dl](https://go.dev/dl/).
2. Instale o pacote.
3. Confirme se `/usr/local/go/bin` está no PATH.

Validação:

```bash
go version
go env GOROOT GOPATH
```

Opção alternativa com Homebrew:

```bash
brew update
brew install go
go version
```

### Linux

Instalação manual com tarball oficial:

```bash
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.26.0.linux-amd64.tar.gz
```

Adicione ao shell (`~/.bashrc` ou `~/.zshrc`):

```bash
export PATH=$PATH:/usr/local/go/bin
```

Recarregue o shell e valide:

```bash
source ~/.bashrc
# ou
source ~/.zshrc

go version
go env GOROOT GOPATH
```

## Comandos iniciais úteis

```bash
go version
go env
go help
go help build
go help test
```

## Verificação final do ambiente

Se a instalação estiver correta, os comandos abaixo devem funcionar sem erro:

```bash
go version
go env GOOS GOARCH GOROOT GOPATH
```

## Erros comuns

- PATH não configurado corretamente.
- Mais de uma instalação de Go ativa na máquina.
- Terminal antigo aberto antes da alteração de PATH.

## Próximo tópico recomendado

- [Olá, mundo!](../hello_world/README.md)
