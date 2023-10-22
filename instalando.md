---
title: "Instalando Go no Windows"
description: "Como instalar a linguagem Go no Windows."
tags:
    - golang
    - grupo-estudos-golang
date: "2023-10-21T14:39:08-03:00"
weight: 6
draft: false
---

# Instalando

## Download

Faça o download do instalador correspondente ao seu sistema operacional em [https://go.dev/dl/](https://go.dev/dl/).

### Microsoft Windows

Apos executar o instalador, Go estará instalada em

```
C:\Go
```

**Lembre!** Você está instalando um compilador, desative qualquer software antivirus antes de compilar seus projetos. Muitos antivirus simplesmente bloqueiam o compilador silenciosamente, sem nenhum alerta e daí o compilador não vai conseguir gerar o executável dos exemplos.

---
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

Após isso o Go estará instalado em:
```
/usr/local/bin/go
```

---
### Linux

Apos executar o instalador Go estará instalada no diretório:

```
/usr/local/go
```

---
## Git

Para seguir os exemplos que envolvem baixar pacotes diretamente do github é necessário ter o git instalado, você pode baixar uma versão para o seu sistema no site oficial.

[https://git-scm.com](https://git-scm.com)


---
# Ambiente de desenvolvimento

## Atom

Você pode usar qualquer editor de texto puro para editar seu código fonte em Go, um que eu gosto é o Atom, é free e muito bom. Depois de instalar vá nas configurações do editor e procure pelos plugins de Go que achar interessante.

[https://atom.io](https://atom.io)

Sugestões de plugins:

[https://atom.io/packages/go-plus](https://atom.io/packages/go-plus)

## Visual Studio Code

Outro ótimo editor é o Visual Studio Code, assim como o Atom, depois de instalar o editor é necessário adicionar os plugins de Go.

[https://code.visualstudio.com](https://code.visualstudio.com)

## Sublime Text

Se você busca o máximo de simplicidade, o Sublime poderá lhe atender. Além de ser um ótimo editor, ele também permite a intalação de plugins que auxiliam no desenvolvimento em Go.

[https://www.sublimetext.com/](https://www.sublimetext.com/)

Sugestões de plugins:

[https://github.com/DisposaBoy/GoSublime](https://github.com/DisposaBoy/GoSublime)

---
### Mas eu não quero instalar nada!

Se você não quiser fazer a instalação para testar os exemplos você pode usar o [The Go Playground](https://play.golang.org)

---

[< Inicio](README.md) - [Configurando >](configurando.md)
