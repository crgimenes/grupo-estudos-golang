package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)


func main() {

	// Obtém o feed de rss através da url
	resp, err := http.Get("http://www.techworld.com/news/rss")
	
	if err != nil{
		s := err.Error()
		fmt.Printf("%q\n", s)	
		return
	}
	
	// Sinaliza que a última ação a ser feita no programa é o fechamento da resposta
	defer resp.Body.Close()
	
	// Verifica se o código de status é 200, indicando assim o sucesso da solicitação
	if resp.StatusCode != 200 {
		println(resp.StatusCode)	
		return
	}

	// Por fim escreve o conteúdo do feed de rss
	bodyBytes, err2 := ioutil.ReadAll(resp.Body) 
	
	if err2 != nil{
		s := err2.Error()
		fmt.Printf("%q\n", s)	
		return
	}
	
    bodyString := string(bodyBytes) 
	println(bodyString)

}
