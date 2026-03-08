# Um Crawler simples
Nesse exemplo iremos aprender a usar um parser HTML, o [GoQuery](https://github.com/PuerkitoBio/goquery), para coletar
as últimas notícias da [UOL](https://noticias.uol.com.br/ultimas/index.htm). Esse crawler é apenas um exemplo 
didático, não visa prejudicar o site da Uol.

Com o GoQuery, conseguimos buscar os elementos da página, buscando classes, ids etc, de forma "parecida" com jQuery.

Queremos pegar essas informações:
- Data da publicação;
- Descrição;
- Fonte;
- Imagem;
- Título da notícia;
- URL.

A primeira coisa que fazemos quando vamos capturar algum dado de uma página HTML é inspecionar a página e ver aonde
estão as informações que queremos. Deixarei como exercício você fazer isso :D

Instale a lib que iremos usar:

`go get github.com/PuerkitoBio/goquery`

Mãos à obra!
