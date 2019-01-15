package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

/*inspirado por http://funcoeszz.net/man.html#zzdolar*/

var errTratamento = errors.New("Não foi possivel encontrar valores")

//BuscaRequest Busca pagina BC
func BuscaRequest() (contents []byte, err error) {
	response, err := http.Get("https://ptax.bcb.gov.br/ptax_internet/consultarUltimaCotacaoDolar.do")
	if err != nil {
		return
	}
	defer response.Body.Close()
	contents, err = ioutil.ReadAll(response.Body)
	return
}

//TrataRequest Recebe o conteudo da pagina e retorna os valores de compra e venda
func TrataRequest(contents []byte) (compra, venda string, err error) {
	re := regexp.MustCompile("[1-9],[0-9][0-9][0-9][0-9]")
	valores := re.FindAll(contents, -1)
	if len(valores) == 2 {
		compra = string(valores[0][:])
		venda = string(valores[1][:])
	} else {
		err = errTratamento
	}
	return

}

func main() {
	contents, err := BuscaRequest()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	compra, venda, err := TrataRequest(contents)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(" Preço do dólar Banco Central.")
	fmt.Println(" R$:", compra, "para Compra.")
	fmt.Println(" R$:", venda, "para Venda.")
}
