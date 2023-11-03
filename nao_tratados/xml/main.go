package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

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

func closer(body io.Closer) {
	err := body.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	resp, err := http.Get("http://pizzadedados.com/feed.xml")
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	defer closer(resp.Body)

	// Mostra o XML
	fmt.Println(string(body))
	fmt.Println("-=-=-=-=-=-=-=-=-=-=-=-=-=-=-")

	rss := RSS{}
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		fmt.Println(err)
	}

	// lista o feed
	for k, i := range rss.ItemList {
		fmt.Println(k, i.Title)
	}
}
