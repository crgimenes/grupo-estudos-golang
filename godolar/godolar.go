package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

/*
inspirado por http://funcoeszz.net/man.html#zzdolar
*/

func main() {
	response, err := http.Get("https://ptax.bcb.gov.br/ptax_internet/consultarUltimaCotacaoDolar.do")
	if err != nil {
		fmt.Println(err)
	}

	var contents []byte
	defer response.Body.Close()
	contents, err = ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	re := regexp.MustCompile("[1-9],[0-9][0-9][0-9][0-9]")

	valores := re.FindAll(contents, -1)

	fmt.Println(" Preço do dólar Banco Central.")
	fmt.Println(" R$:", string(valores[0][:]), "para compra.")
	fmt.Println(" R$:", string(valores[1][:]), "para Venda.")

}
