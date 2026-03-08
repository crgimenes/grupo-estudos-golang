package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		response, err := http.PostForm("http://www.dict.org/bin/Dict/", url.Values{"Form": {"Dict1"}, "Query": {os.Args[1]}, "Strategy": {"*"}, "Database": {"*"}, "submit": {"Submit query"}})
		if err != nil {
			fmt.Println("[mai] - HTTP Post Form Error: " + err.Error())
			return
		}
		defer closer(response.Body)
		if response.StatusCode == 200 {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				fmt.Println("[main] - Failure into reading payload body. " + err.Error())
				return
			}
			fmt.Println("==================== DICT.ORG ====================")
			fmt.Println(formatBody(body))
		} else {
			fmt.Println("[main] - Server Status Error Description: " + response.Status)
			return
		}
	} else {
		fmt.Println("[main] - Informe uma palavra, como parâmetro, para pesquisar no DICT.ORG")
	}
}

func closer(c io.Closer) {
	err := c.Close()
	if err != nil {
		fmt.Println("[closer function] - Closer Error: " + err.Error())
		return
	}
}
func formatBody(body []byte) (b string) {
	var lg bool
	for _, c := range body {
		if c == '<' {
			lg = true
		} else if c == '>' && lg == true {
			lg = false
		} else if lg == false {
			b += string(c)
		}
	}
	return b[strings.Index(b, "definitions")-5 : strings.LastIndex(b, "Questions")]
}
