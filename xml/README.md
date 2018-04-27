# XML


Este é um exemplo de como ler um feed RSS2 e interpretar o retorno para extrair os dados que queremos.

Primeiro criamos as structs que vão conter os dados vindos do XML

```go
// RSS contem a base da estrutura do feed
type RSS struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Title       string   `xml:"channel>title"`
	Link        string   `xml:"channel>link"`
	Description string   `xml:"channel>description"`
	PubDate     string   `xml:"channel>pubDate"`
	ItemList    []Item   `xml:"channel>item"`
}

// Item contem os posts no feed
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Content     string `xml:"encoded"`
	PubDate     string `xml:"pubDate"`
	Comments    string `xml:"comments"`
}
```

Então baixamos o feed da internet

```go
resp, err := http.Get("http://pizzadedados.com/feed.xml")
if err != nil {
	fmt.Println(err)
	return
}

body, err := ioutil.ReadAll(resp.Body)
if err != nil {
	fmt.Println(err)
}
defer closer(resp.Body)
```

Finalmente interpretamos os dados o armazenamos na struct

```go
rss := RSS{}
err = xml.Unmarshal(body, &rss)
if err != nil {
	fmt.Println(err)
}
```
Pronto agora só precisamos exibir os dados

```go
for k, i := range rss.ItemList {
	fmt.Println(k, i.Title)
}
```