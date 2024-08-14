#  Grupo de Estudos de Golang <a rel="license" href="http://creativecommons.org/licenses/by-sa/4.0/"><img alt="Licença Creative Commons" style="border-width:0" src="https://i.creativecommons.org/l/by-sa/4.0/88x31.png" /></a>


## Exemplos antigos

Os exemplos estão sendo reescritos, muita coisa antiga precisa ser ajustada e melhorada. Os exemplos antigos estão no diretório "nao_tratados".

--- 

[Sobre o Grupo de Estudos de Go](#um-pouco-de-história)

# Ementa do Curso de Programação em Go

## Módulo 1: Introdução ao Go

### 1: Bem-vindo ao Go

- [História e filosofia do Go](#história-e-filosofia-do-go)
- [Instalação e configuração do ambiente](#instalação-e-configuração-do-ambiente)
- [Verificando a instalação](#verificando-a-instalação)
- [Primeiro programa em Go: "Hello, World!"](#primeiro-programa-em-go-hello-world)

### 2: Fundamentos da Linguagem

- [Sintaxe básica](#sintaxe-básica)
- Tipos de dados primitivos
- Variáveis e constantes
- Estruturas de controle (if, else, switch)

### 3: Funções

- Declaração e chamada de funções
- Parâmetros e retorno de funções
- Funções anônimas e closures

### 4: Estruturas de Dados

- Arrays e slices
- Mapas
- Structs

## Módulo 2: Intermediário

### 5: Programação Orientada a Objetos em Go

- Structs e métodos
- Interfaces
- Polimorfismo

### 6: Manipulação de Erros

- Tratamento de erros
- Padrões de erro
- Customização de erros

### 7: Concorrência em Go

- Goroutines
- Canais (channels)
- Select e sincronização

### 8: Input e Output

- Leitura e escrita de arquivos
- Manipulação de JSON e XML
- Uso de pacotes padrão para I/O

## Módulo 3: Avançado

### 9: Rede e Web

- Desenvolvimento de servidores HTTP
- Construção de APIs RESTful
- Manipulação de sockets e comunicação em rede

### 10: Banco de Dados

- Conexão e manipulação de bancos de dados SQL (usando SQLite como exemplo)
- Transações e migrações
- ORM com GORM

### 11: Testes em Go

- Testes unitários
- Benchmarking

### 12: Ferramentas e Boas Práticas

- Uso do Go Modules
- Linters e formatação de código
- Profiling e otimização de desempenho

## Módulo 4: Projetos Práticos

### 13: Projeto 1 - Aplicação CLI

- Desenvolvimento de uma aplicação de linha de comando
- Parsing de argumentos e flags
- Automação de tarefas

### 14: Projeto 2 - Web Application

- Construção de uma aplicação web completa
- Templates HTML e manipulação de formulários
- Autenticação e autorização

### 15: Projeto 3 - Sistema de BBS

- Integração com KotobaHub
- Implementação de funcionalidades de chat, fórum e wiki
- Implementação de acesso e controle de usuários

### 16: Projeto 4 - Bot do Telegram

- Criação e configuração de um bot do Telegram
- Uso da API do ChatGPT-4 para interações
- Manutenção de contexto nas conversas

## Módulo 5: Tópicos Especiais

### 17: Desenvolvimento para Microserviços

- Arquitetura de microserviços
- Comunicação entre serviços
- Deploy e escalabilidade

### 18: Integração Contínua e Deploy Contínuo (CI/CD)

- Ferramentas de CI/CD (por exemplo, GitHub Actions)
- Automação de builds e testes
- Deploy automático

### 19: Segurança em Aplicações Go

- Práticas de segurança
- Proteção contra vulnerabilidades comuns
- Uso de pacotes seguros

### 20: Performance e Escalabilidade

- Técnicas avançadas de profiling
- Otimização de código
- Estratégias de escalabilidade

## Recursos Adicionais

- [Perguntas frequentes](#perguntas-frequentes)
- Leituras recomendadas
- Exercícios práticos e desafios
- [Como contribuir](#como-contribuir)
- [Como fazer boas perguntas e dar boas respostas](#como-fazer-boas-perguntas-e-dar-boas-respostas)
- Comunidade de desenvolvedores Go
    - [Telegram](https://crg.eti.br/grupo-estudos-golang/telegram)
    - [Discord](https://crg.eti.br/grupo-estudos-golang/discord)
    - [YouTube](https://crg.eti.br/grupo-estudos-golang/youtube)
- Sessões de mentoria
- Licença de uso

--- 

## Um pouco de história

Por volta de 2015 eu estava procurando uma linguagem de programação moderna que atendesse a alguns critérios que considero importantes, deveria ser simples, compilada, abrangente e com uma boa comunidade. Com capacidade para criar executáveis multiplataforma e que aproveitasse as características de multi-processamento das máquinas modernas.

Go caiu como uma luva e comecei a usar nos meus projetos. No início de 2016 no ABC Makerspace criei um pequeno curso, o Go-Hands-On e de lá iniciamos o que acabaria se tornando o grupo de estudos. 

Desde aquela época os objetivos do grupo continuam praticamente os mesmos, estudar a linguagem de programação Go e tecnologias relacionadas, dar apoio para quem está iniciando assim como manter atualizados os programadores mais experientes, criar conteúdo em português falado no Brasil.

No dia 11 de maio de 2016 criei o repositório do Grupo de Estudos de Golang no GitHub, essa se tornou nossa data de início oficial. De lá para cá fizemos centenas de encontros, praticamente toda semana, escrevemos muitos exemplos, gravamos muitos vídeos, mas o mais importante é que conhecemos muita gente boa disposta a ajudar e compartilhar o que sabe.

Esse material doi criado com ajuda de inumeras pessoas, sinta-se a vontade para usar, compartilhar e melhorar.

[CRG](https://crg.eti.br)

---

## Perguntas frequentes

### Qual framework, ORM, banco de dados, etc. usar?

- Go é uma linguagem de programação minimalista, não possui um framework padrão como Django, Rails ou Laravel. Antes de procurar algo externos tente usar as bibliotecas padrão, elas são muito boas e completas.

- ORM é um conceito que não é muito usado em Go, idealmente você deve usar um banco de dados que tenha um driver nativo para Go, como PostgreSQL, MySQL, SQLite, etc. e escrever seus proprios SQLs.

- Frameworks web são muitos, alguns populares são o Gin, Echo, Fiber, Beego, etc. mas você pode usar o pacote `net/http` da biblioteca padrão, ele provavelmente vai resolver seu problema sem adicionar complexidade ou dependências.

### O que é migration?

- Migration é uma ferramenta que permite que você controle as mudanças no banco de dados, geralmente é usado para criar tabelas, adicionar colunas, etc. Existem muitas ferramentas de migration, você pode usar uma delas ou escrever suas proprias migrations.

### Qual paradigma/padrão de projeto usar em Go.

- Go esta mais para programação estruturada/procedural. Go não possui um padrão de projeto como MVC, MVVM, etc. Você pode usar o que achar melhor. Orientação a objetos é possivel em Go, mas não é muito comum.

### Porque o tratamento de erro é tão verborragico?

- Go não possui exceções, o tratamento de erro é feito com o retorno de um erro, isso é uma escolha de design da linguagem. O tratamento de erro é explicito e isso é uma das coisas que torna Go uma linguagem simples e fácil de entender. Os erros são tratados imediatamente onde eles acontecem tornando o debug mais fácil. Uma dica importante é sempre tratar todos os erros, muitas vezes um erro que parece impossível de acontecer, acontece.

---

## Como fazer boas perguntas e dar boas respostas

Algumas dicas para melhorar a interação com o grupo quando for enviar perguntas ou respostas.

**Especifique o Problema:**

- Deixe claro qual é o problema que você está enfrentando. Uma descrição precisa auxiliará os outros a entenderem e oferecerem soluções.

**Demonstração do Código:**

- Forneça um trecho de código, e inclua a saída esperada e atual. Isso dará aos outros uma ideia melhor do que está acontecendo.

- Você pode usar o [The Go Playground](https://go.dev/play/) para compartilhar exemplos de código. ou então use o markdown para formatar o código. Geralmente basta colocar o código entre três crases.

**Ambiente de Desenvolvimento:**

- Mencione a versão do Go que você está usando e outras configurações de ambiente que possam ser relevantes para o problema, como o sistema operacional, por exemplo.

**Formato é importante:**

- Formate sua pergunta de maneira clara e legível, com uso adequado de marcações, para facilitar a leitura por outros membros do grupo.

**Revise o Texto:**

- Antes de perguntar utilize as ferramentas de busca. Mesmo que não encontre a resposta, isso ajudará a formular sua dúvida de maneira mais clara.

- Consulte nosso [FAQ](#faq) e o [código de conduta](#codigo-de-conduta).

- Ao formular uma pergunta, inclua as etapas que já tentou para resolver o problema.

- Mantenha a pergunta concisa; perguntas mais curtas tendem a ser respondidas mais rapidamente.

- É útil demonstrar o problema com um trecho de código. Evite enviar imagens da tela; em vez disso, crie um pequeno exemplo que ilustre a dúvida ou problema em poucas linhas.

- Se desejar, use o [The Go Playground](https://go.dev/play/). No entanto, esteja ciente de que o playground grava os exemplos compartilhados, então evite colocar dados ou códigos sensíveis.

- Se ninguém responder sua pergunta, não se preocupe. Pode ser que ninguém tenha a resposta no momento. Não leve para o lado pessoal.

- Respeite a opinião dos outros. Nós programadores temos fortes opiniões sobre o que é melhor e adoramos discutir o melhor de tudo: melhor procedimento, melhor editor, melhor sistema operacional, etc. Seja objetivo e não prolongue discussões.

- Quando escrever, evite anglicismos mas tente se manter próximo da forma popular da língua. Por exemplo: nós adotamos alguns anglicismos como mouse, kernel, etc. são comuns e aceitáveis. Aqui tem [algumas substituições úteis](https://www.ime.usp.br/~kon/ResearchStudents/traducao.html).

- Quando responder ao grupo, seja exemplar e lembre-se de que você é a referência que outros programadores seguirão.

---

## Como contribuir

### Primeiros passos

De uma olhada nas issues e veja se o que você quer fazer já não esta sendo discutido, se estiver ótimo, participe da discussão e de suas ideias caso não esteja você pode criar uma nova issue para discutirmos ou se preferir pode também mandar diretamente um _pull request_, só não esqueça de descrever muito bem o que você quer mudar/adicionar.

### Formato

Material novo idealmente deve ser organizados em diretórios e ter três arquivos, um README.md descrevendo como executar o exemplo, um exemplo.go contendo o exemplo propriamente dito e um exemplo_test.go. (alternativamente você pode usar main.go e main_test.go). Também é necessário fazer uma descrição do exemplo no arquivo README.md do diretório principal.

```
exemplo/README.md
exemplo/exemplo.go
exemplo/exemplo_test.go
```

### Cuidado com o código

Antes de mandar um _pull request_ formate seu código com _go fmt_. Também é uma boa ideia passar uma ferramenta de analize estatica como golint por exemplo.

### Não se prolongue demais

Uma das características que é importante manter são exemplos curtos e de fácil entendimento, tenta fazer com que os exemplos caibam em uma tela, tudo bem se não for possível, apenas mantenha isso em mente.

Exemplos rápidos, curtos, diretos e de fácil entendimento são nossa meta.

### Inclua seus exemplos no The Go Playground

Inclua um link sob os exemplos que você escrever apontando para o https://play.golang.org, dessa fora o leitor pode testar seu exemplo imediatamente. Obviamente isso não é possivel com qualquer exemplo, mas onde for possivel é bom colocar.

### Arquivo README.md

Discuta com os outros desenvolvedores o melhor lugar para colocar o seu tutorial na lista do arquivo README.md, a ideia é que os tutoriais estejam em ordem complexidade.

### Enviando uma contribuição

- Faça um _fork_ do projeto.
- Crie uma _branch_ com as suas modificações `git checkout -b exemplo-fantastico`.
- Faça _commit_ das suas alterações `git commit -m 'Implementação do novo exemplo fantástico'`.
- Faça um _push_ na sua _branch_ `git push origin exemplo-fantastico`.
- Faça um _pull request_ com suas alterações.

---

## Licença de uso

Esse material é livre, sob a licença Creative Commons, CC BY-SA. você pode usar da forma que preferir, apenas lembre de citar a fonte, não apenas copie, no lugar disso acrescente, melhore e compartilhe, assim todos ganhamos.

[Cesar Gimenes](crg.eti.br)

---

# Bem-vindo ao Go

## História e filosofia do Go

Go, foi criado no Google em 2007 por Robert Griesemer, Rob Pike e Ken Thompson. 

A primeira versão estável, Go 1.0, foi lançada em março de 2012. 

O objetivo do Go é ser uma linguagem simples, eficiente e com uma curva de aprendizado suave, mantendo um equilíbrio entre desempenho e facilidade de uso. 

Go foi projetado para resolver problemas encontrados em linguagens tradicionais como C++ e Java, oferecendo concisão, simplicidade e segurança na concorrência.

Os mantenedores da linguagem Go se comprometem a manter a [retrocompatibilidade com a versão 1.0](https://go.dev/doc/go1compat), garantindo que programas escritos para versões anteriores continuem funcionando sem alterações. Esse compromisso é uma das razões pelas quais Go é uma linguagem popular para desenvolvimento de sistemas de larga escala.

Outro ponto importante para lembrar é que a retrocompatibilidade não é binária, ou seja, não é garantido que um programa escrito em Go 1.0 gere o mesmo binário em versões futuras. O que é garantido é que o código fonte continuará compilando. Um exemplo onde retrocompatibilidade do binário é importante é se você precisar suportar versões antigas de sistemas operacionais ou hardware, seu programa continua compilando sem erros mas não funciona na plataforma antiga.


## Instalação e configuração do ambiente

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

#### Anti-virus

Caso esteja usando algum anti-virus, é necessário fazer ajustes para que o compilador não seja bloqueado, é comum que o anti-virus bloqueie a execução de programas compilados com Go e as vezes de forma silenciosa.

#### Ajuste do PATH

Existem várias formas de configurar a variável PATH no Windows, uma das formas mais simples é usar o Painel de Controle. Você precisa que o diretório `C:\Go\bin` seja adicionado ao PATH.

No Windows 11, você pode usand a tecla de atalho `Win + R` e digitar `sysdm.cpl` e clicar em `OK`. Na janela que abrir, clique na aba `Avançado` e depois no botão `Variáveis de Ambiente`. Na janela que abrir, clique duas vezes na variável `Path` e adicione o diretório `C:\Go\bin` no final da lista.

### Mac OS X

Apos executar o instaldor, Go estará instalada em:

```
/usr/local/go
```

#### Mac OS X - Alternativa - Usando Homebrew

Preferencialmente use o instalador oficial.

#### Ajuste do PATH

Adicione o diretório `bin` do Go ao PATH. Abra o arquivo `~/.bash_profile` ou `~/.zshrc` e adicione a seguinte linha:

```bash
export PATH=$PATH:/usr/local/go/bin
```

#### Mac OS X - Alternativa - Usando Homebrew

Se você usa o Homebrew, o Go pode ser instalado com dois simples comandos:

```
brew update
brew install go
```

No caso da instalação via Homebrew, o PATH é configurado automaticamente.

### Linux

Você pode descompactar o arquivo tar.gz em `/usr/local` ou `$HOME` e configurar a variável PATH.

```
tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

#### Ajuste do PATH

Adicione o diretório `bin` do Go ao PATH. Abra o arquivo `~/.bashrc` ou `~/.zshrc` e adicione a seguinte linha:

```bash
export PATH=$PATH:/usr/local/go/bin
```

Uma dica importante, quando você quiser atualizar o Go é bom remover a versão antiga antes de instalar a nova, os instaladores para Mac e Windows fazem isso automaticamente, mas no Linux é necessário fazer manualmente.

Como super usuário, execute o comando:

```bash
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
```

Isso ira remover a versão antiga e instalar a nova.

## Verificando a instalação

Para verificar se a instalação foi bem sucedida, abra um terminal e execute o comando:

```bash
go version
```

Você deve ver algo como:

```
go version go1.22.0 linux/amd64
```

## Primeiro programa em Go: "Hello, World!"

"Hello, World!" é um clássico primeiro programa, é extremamente simples, mas é ótimo para comprovar que o ambiente de desenvolvimento está funcionando corretamente.

Em Go o programa inicia na função `main` e o pacote principal é `main`.

```go
package main

func main() {
    println("Hello, World!")
}
```

A primeira linha do programa é o pacote, em Go os programas são organizados em pacotes, o pacote principal é `main`, e o nome do pacote deve ser o mesmo que o nome do diretório.

A segunda linha é a função `main`, ela é a função chamada quando o programa é executado.

E finalmente, terceira linha é a chamada para a função `println`, que imprime a mensagem "Hello, World!" no terminal.

Para executar o programa, salve o código em um arquivo chamado `hello.go` e execute o comando:

```bash
go run hello.go
```

Você deve ver a mensagem "Hello, World!" impressa no terminal.

Go não tem versão interpretada, mesmo esse teste usando `go run` é compilado antes de ser executado, o executável é temporário e é removido após a execução.

---

## 2: Fundamentos da Linguagem

### Sintaxe básica

Go é uma linguagem de programação com sintaxe minimalista parecida com C e linguagens derivadas. Se você já programou em C, C++, Java ou JavaScript, vai sentir muita familiaridade com Go.

Aqui vamos fazer uma passagem rápida sobre a sintaxe, você pode encontrar mais detalhes na [documentação oficial](https://golang.org/doc/).

#### Pacotes

No decorrer deste texto falaremos muito sobre pacotes, eles são uma parte fundamental da organização de código em Go.

Um pacote é composto por um ou mais arquivos fonte em Go e é organizado em diretórios. idealmente o nome do pacote deve ser o mesmo que o nome do diretório.

Não é permitido ter mais de um pacote em um mesmo diretório.

Todo programa Go começa na função `main` do pacote `main`.

Você pode criar múltiplos pacotes para organizar seu código e pode importar pacotes de outros pacotes. Falaremos mais sobre importação de pacotes de terceiros quando falarmos de ferramentas como o `go mod`.

Os pacotes são importados no início do arquivo com a palavra-chave `import`.

Exemplo:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Nesse exemplo, foi importado o pacote `fmt` que é um pacote da biblioteca padrão de Go, ele é usado para formatar a saída do programa.

Você pode também importar múltiplos pacotes com um único comando import.

```go
import (
    "fmt"
    "os"
)
```

Um ponto importante para ter em mente quando estiver organizando seu projeto é que as dependências do Go crescem em apenas uma direção, em outras palavras, um pacote não pode importar um pacote que o importa, isso criaria uma dependência circular e resultaria em um erro de compilação.

#### Comentários

Comentários em Go são feitos com `//` para comentários de linha e `/* */` para comentários multiplas linhas, igual a C e C++ e tantas outras linguagens `C-like`.

```go
// Este é um comentário de linha

/*
Este é um comentário
com múltiplas 
linhas
*/
```


#### Variáveis e constantes

#### Tipos de dados primitivos

#### Funções

#### Estruturas de controle (if, else, switch)






