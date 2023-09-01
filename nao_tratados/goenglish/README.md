# Go English

Go English é a versão em Go do zzenglish.

Funções ZZ é um conjunto dos mais variados aplicativos, escritos em shell script, com as mais variadas aplicações. Dentre eles, o zzenglish, objeto deste estudo.

## O ZZEnglish

Abaixo, temos a codificação original, do miniaplicativo.
    
```sh
zzenglish ()
{
	zzzz -h english "$1" && return

	test -n "$1" || { zztool -e uso english; return 1; }

	local cinza verde amarelo fecha
	local url="http://www.dict.org/bin/Dict?Form=Dict2&Database=*&Query=$1"

	if test $ZZCOR -eq 1
	then
		cinza=$(  printf '\033[0;34m')
		verde=$(  printf '\033[0;32;1m')
		amarelo=$(printf '\033[0;33;1m')
		fecha=$(  printf '\033[m')
	fi

	zztool dump "$url" | zzutf8 |
		sed "
			/Questions or comments about this site./d
			# pega o trecho da página que nos interessa
			/[0-9]\{1,\} definitions\{0,1\} found/,/ *[_-][_-][_-][_-][_-]* *$/!d
			s/_____*//
			s/-----*//
			# protege os colchetes dos sinônimos contra o cinza escuro
			s/\[syn:/@SINONIMO@/g
			# aplica cinza escuro em todos os colchetes (menos sinônimos)
			s/\[/$cinza[/g
			# aplica verde nos colchetes dos sinônimos
			s/@SINONIMO@/$verde[syn:/g
			# 'fecha' as cores de todos os sinônimos
			s/\]/]$fecha/g
			# # pinta a pronúncia de amarelo - pode estar delimitada por \\ ou //
			s/\\\\[^\\]\{1,\}\\\\/$amarelo&$fecha/g
			s|/[^/]\{1,\}/|$amarelo&$fecha|g
			# cabeçalho para tornar a separação entre várias consultas mais visível no terminal
			/[0-9]\{1,\} definitions\{0,1\} found/ {
				H
				s/.*/==================== DICT.ORG ====================/
				p
				x
			}" |
		zztrim -V -r |
		zzsqueeze -l
}
```

## Exemplo de Implementação

```go
package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
			body, err := ioutil.ReadAll(response.Body)
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

// tip [@crgimenes] (https://github.com/crgimenes) 
// [estudos-go] (https://youtu.be/eEU9CwVkJt8)
func closer(c io.Closer) { 
	err := c.Close()
	if err != nil {
		fmt.Println("[closer function] - Closer Error: " + err.Error())
		return
	}
}
//formatBody - removes html's tags of body
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
```

---
O Projeto ZZ está disponível no [GitHub](https://github.com/funcoeszz/funcoeszz).
Para contribuições, leiam o [README.md](https://github.com/funcoeszz/funcoeszz/blob/master/README.md)
