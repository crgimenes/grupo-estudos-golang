# Instalando

## Download

### Microsoft Windows
Faça o download do arquivo [go1.7.3.windows-amd64.msi](https://storage.googleapis.com/golang/go1.7.3.windows-amd64.msi).

Apos executar o instalador, Go estará instalada em

```
C:\Go
```

**Lembre!** Você está instalando um compilador, desative qualquer software antivirus antes de compilar seus projetos. Muitos antivirus simplesmente bloqueiam o compilador silenciosamente, sem nenhum alerta e daí o compilador não vai conseguir gerar o executável dos exemplos.

---
### Mac OS X
Faça o download do arquivo [go1.7.3.darwin-amd64.pkg](https://storage.googleapis.com/golang/go1.7.3.darwin-amd64.pkg).

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
[go1.7.3.linux-amd64.tar.gz](https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz)

Apos executar o instalador Go estará instalada no diretório:

```
/usr/local/go
```
---
### Direto pelos fontes
Você também pode baixar os fontes e instalar você mesmo.

[go1.7.3.src.tar.gz](https://storage.googleapis.com/golang/go1.7.3.src.tar.gz)

---
## Git

Para seguir os exemplos que envolvem baixar pacotes diretamente do github é necessário ter o git instalado, você pode baixar uma versão para o seu sistema no site oficial.

[https://git-scm.com](https://git-scm.com)


---
# Ambiente de desenvolvimento

## Atom

Você pode usar qualquer editor de texto puro para editar seu código fonte em Go, um que eu gosto é o Atom, é free e muito bom. Depois de instalar vá nas configurações do editor e procure pelos plugins de Go que achar interessante.

[https://atom.io](https://atom.io)

## Visual Studio Code

Outro ótimo editor é o Visual Studio Code, assim como o Atom, depois de instalar o editor é necessário adicionar os plugins de Go.

[https://code.visualstudio.com](https://code.visualstudio.com)

---
### Mas eu não quero instalar nada!

Se você não quiser fazer a instalação para testar os exemplos você pode usar o [The Go Playground](https://play.golang.org)

---

[< Inicio](README.md) - [Configurando >](configurando.md)
