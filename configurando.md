# Configurando

Para que o ambiente de desenvolvimento funcione corretamente é necessário configurar algumas variáveis de ambiente.

**GOROOT** - Deve apontar para o diretório de instalação do Go, você só precisa se preocupar com essa variável se você preferiu instalar a linguagem em outro diretório, o padrão é /usr/local/go/bin ou C:\\Go no caso do Windows.

**PATH** - Deve apontar para o diretório onde os binários foram instalados normalmente /usr/local/go/bin ou C:\\Go\\bin.

**GOPATH** - Deve apontar para seu diretório de trabalho.

## Windows

Instalando pelo MSI o sistema já deve fazer o ajuste no PATH, mas caso seja necessário alguma alteração basta ir no "Painel de Controle" -> "Sistema" -> "Avançado" -> Variáveis de ambiente.
Em algumas versões do Windows você deve ir em "Configurações avançadas do sistema" ->  "Variáveis de ambiente".

## Mac OS X e Linux

No Mac, Linux e BSD você pode adicionar essas variáveis no arquivo de configuração do shell que você estiver usando como .profile no caso do bash ou .zshrc no caso do zsh.

Exemplo:

```bash
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=~/projeto
```

**Obs:** Caso ainda não esteja você também precisa apontar o git na variavel PATH do seu sistema.

---
## Testando a instalação

Depois de instalado e configurado você pode verificar se Go esta respondendo corretamente pelo comando

```bash
go version
```

Para testar o git execute:

```bash
git --version
```
---
[Inicio](README.md)

[< Instalando](instalando.md) - [Olá Mundo >](ola_mundo.md)
