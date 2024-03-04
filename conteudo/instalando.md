---
title: "Instalando Go"
description: "Como instalar a linguagem Go."
tags:
    - golang
    - grupo-estudos-golang
date: "2023-10-21T14:39:08-03:00"
weight: 6
draft: false
---

# Instalando

A instalação do Go é bastatnte simples, o caminho mais direto é baixar o instalador para a sua plataforma e executá-lo, ou se preferir baixar os binarios.

Em ambienes UNIX/Linux geralmente é comum instalar o Go em `/usr/local/go` ou `$HOME/go`, e em ambientes Windows em `C:\Go`.

Apos a instalação é necessário configurar a variável PATH para que o sistema operacional possa encontrar o compilador.

Versõs antigas do Go exigiam que outras variaveis fossem condiguradas, hoje essas configurações são opcionais.

**Lembre!** Caso esteja usando algum anti-virus, é necessário adicionar o diretório de instalação do Go e seu diretório de projetos na lista de exceções. É comum que o anti-virus bloqueie a execução de programas compilados com Go e as vezes de forma silenciosa.

## Download

Faça o download do instalador correspondente ao seu sistema operacional ou dos binarios em [https://go.dev/dl/](https://go.dev/dl/).

### Microsoft Windows

Apos executar o instalador, Go estará instalada em

```
C:\Go
```

### Mac OS X

Apos executar o instaldor, Go estará instalada em:

```
/usr/local/go
```

#### Mac OS X - Alternativa - Usando Homebrew

Se você usa o Homebrew, o Go pode ser instalado com dois simples comandos:

```
brew update
brew install go
```

### Linux

Você pode descompactar o arquivo tar.gz em `/usr/local` ou `$HOME` e configurar a variável PATH.

```
tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

