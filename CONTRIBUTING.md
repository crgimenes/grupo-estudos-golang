Esse material foi criado inicialmente para um workshop de duas horas e com o passar do tempo foi usando em outras demonstrações de Go e rapidamente começou a receber solicitações de itens que haviam sido deixados de fora porque não cabiam no formato de duas horas e também contribuições de outros programadores que acharam o material util.

Como crescer é inevitável parece uma boa ideia abraçar essa oportunidade e extender o material para se tornar um repositório de código e exemplos que possa ser usado por todos para estudo e referencia.

## Primeiros passos

De uma olhada nas issues e veja se o que você quer fazer já não esta sendo discutido, se estiver ótimo, participe da discussão e de suas ideias caso não esteja você pode criar uma nova issue para discutirmos ou se preferir pode também mandar diretamente um _pull request_, só não esqueça de descrever muito bem o que você quer mudar/adicionar.

## Formato

Material novo idealmente deve ser organizados em diretórios e ter três arquivos, um README.md descrevendo o exemplo, um exemplo.go contendo o exemplo propriamente dito e um exemplo_test.go.

```
exemplo/README.md
exemplo/exemplo.go
exemplo/exemplo_test.go
```

O material antigo esta aos poucos sendo reformatado para esse novo padrão.

## Cuidado com o código

Antes de mandar um _pull request_ formate seu código com _go fmt_. Também é uma boa ideia passar uma ferramenta de analize estatica como golint por exemplo.

## Variaveis 

Inicialmente todas as variáveis eram escritas em ingles porque fica um código mais idiomaticamente homogêneo, mas para quem esta iniciando na linguagem é fácil confundir variáveis com palavras reservadas, então para resolver esse problema adotamos o português para nomes de variáveis. Fora do âmbito desse material é uma boa pratica nomear as variáveis em ingles.


## Não se prolongue demais

Uma das características que é importante manter são exemplos curtos e de fácil entendimento, tenta fazer com que os exemplos caibam em uma tela, tudo bem se não for possível, apenas mantenha isso em mente.

Exemplos rápidos, curtos, diretos e de fácil entendimento são nossa meta.

## Arquivo README.md e Rodapé

Discuta com os outros desenvolvedores o melhor lugar para colocar o seu tutorial na lista do arquivo README.md, a ideia é que os tutoriais estejam em ordem complexidade.

No rodapé do seu tutorial coloque links para o inicio, tutorial anterior e proximo tutorial como nesse exemplo:

```
---
[Inicio](../README.md)

[< Struct](../struct/) - [Loop for >](../for/)
```

## Enviando uma contribuição

- Faça um _fork_ do projeto.
- Crie uma _branch_ com as suas modificações `git checkout -b exemplo-fantastico`.
- Faça _commit_ das suas alterações `git commit -m 'Implementação do novo exemplo fantástico'`.
- Faça um _push_ na sua _branch_ `git push origin exemplo-fantastico`.
- Faça um _pull request_ com suas alterações.
