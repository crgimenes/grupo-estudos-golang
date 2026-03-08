package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type News struct {
	Description string
	ImageURL    string
	Source      string
	Time        string
	Title       string
	URL         string
}

// Faz a requisição de um site dada uma URL.
func getResponse(url string) *http.Response {
	// Em crawlers que irão fazer muitas requisições, é uma boa prática dar um sleep para não sobrecarregar o servidor.
	// Não é o caso desse crawler, mas fica como dica, espera um tempo aleatório de 0 a 5 segundos.
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if res.StatusCode != 200 {
		panic(fmt.Sprintf("status code error: %d %s", res.StatusCode, res.Status))
	}

	return res
}

// Retorna as últimas notícias do Uol:
func getLastNews() []News {
	url := "https://noticias.uol.com.br/ultimas/index.htm"
	res := getResponse(url)
	lastNews := []News{}

	// Carrega o documento HTML:
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".thumbnail-standard-wrapper").Find("a").Each(func(i int, selection *goquery.Selection) {
		news := News{}

		// Pegando o link da notícia, se não houver o Attr, has será false:
		link, has := selection.Attr("href")
		if has {
			news.URL = link
		}

		// Iremos pegar as imagens agora. Como na árvore de elementos HTML, .thumbnail-standard-wrapper tem
		// .thumb-layer como filho que por sua vez tem um img como filho, podemos pegar essa image:
		link, has = selection.Find(".thumb-layer").Find("img").Attr("src")
		if has {
			news.ImageURL = link
		}

		// Mesmo esquema da imagem mas agora queremos o texto que o elemento tem:
		news.Source = selection.Find(".thumb-caption").Find(".thumb-kicker").Text()
		news.Title = selection.Find(".thumb-caption").Find(".thumb-title").Text()
		news.Description = selection.Find(".thumb-caption").Find(".thumb-description").Text()
		news.Time = selection.Find(".thumb-caption").Find(".thumb-time").Text()

		lastNews = append(lastNews, news)
	})

	return lastNews
}

func main() {
	lastNews := getLastNews()

	for i, news := range lastNews {
		fmt.Printf("Notícia %d\n %+v\n", i+1, news)
	}
}
