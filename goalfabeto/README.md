# Go Alfabeto

    Go Alfabeto é a versão em Go do zzalfabeto.
    Funções ZZ é um conjunto dos mais variados aplicativos, escritos em shell script, 
	com as mais variadas aplicações. Dentre ele, o zzalfabeto, objeto deste estudo.

## O ZZAlfabeto!
    Abaixo, temos a codificação original, do miniaplicativo.
    
```sh
zzalfabeto ()
{
	zzzz -h alfabeto "$1" && return

	local char letra

	local coluna=1
	local dados="\
A:Alpha:Apples:Ack:Ace:Apple:Able/Affirm:Able:Aveiro:Alan:Adam:.-
B:Bravo:Butter:Beer:Beer:Beer:Baker:Baker:Bragança:Bobby:Boy:-...
C:Charlie:Charlie:Charlie:Charlie:Charlie:Charlie:Charlie:Coimbra:Charlie:Charles:-.-.
D:Delta:Duff:Don:Don:Dog:Dog:Dog:Dafundo:David:David:-..
E:Echo:Edward:Edward:Edward:Edward:Easy:Easy:Évora:Edward:Edward:.
F:Foxtrot:Freddy:Freddie:Freddie:Freddy:Fox:Fox:Faro:Frederick:Frank:..-.
G:Golf:George:Gee:George:George:George:George:Guarda:George:George:--.
H:Hotel:Harry:Harry:Harry:Harry:How:How:Horta:Howard:Henry:....
I:India:Ink:Ink:Ink:In:Item/Interrogatory:Item:Itália:Isaac:Ida:..
J:Juliet:Johnnie:Johnnie:Johnnie:Jug/Johnny:Jig/Johnny:Jig:José:James:John:.---
K:Kilo:King:King:King:King:King:King:Kilograma:Kevin:King:-.-
L:Lima:London:London:London:Love:Love:Love:Lisboa:Larry:Lincoln:.-..
M:Mike:Monkey:Emma:Monkey:Mother:Mike:Mike:Maria:Michael:Mary:--
N:November:Nuts:Nuts:Nuts:Nuts:Nab/Negat:Nan:Nazaré:Nicholas:Nora:-.
O:Oscar:Orange:Oranges:Orange:Orange:Oboe:Oboe:Ovar:Oscar:Ocean:---
P:Papa:Pudding:Pip:Pip:Peter:Peter/Prep:Peter:Porto:Peter:Paul:.--.
Q:Quebec:Queenie:Queen:Queen:Queen:Queen:Queen:Queluz:Quincy:Queen:--.-
R:Romeo:Robert:Robert:Robert:Roger/Robert:Roger:Roger:Rossio:Robert:Robert:.-.
S:Sierra:Sugar:Esses:Sugar:Sugar:Sugar:Sugar:Setúbal:Stephen:Sam:...
T:Tango:Tommy:Toc:Toc:Tommy:Tare:Tare:Tavira:Trevor:Tom:-
U:Uniform:Uncle:Uncle:Uncle:Uncle:Uncle:Uncle:Unidade:Ulysses:Union:..-
V:Victor:Vinegar:Vic:Vic:Vic:Victor:Victor:Viseu:Vincent:Victor:...-
W:Whiskey:Willie:William:William:William:William:William:Washington:William:William:.--
X:X-ray/Xadrez:Xerxes:X-ray:X-ray:X-ray:X-ray:X-ray:Xavier:Xavier:X-ray:-..-
Y:Yankee:Yellow:Yorker:Yorker:Yoke/Yorker:Yoke:Yoke:York:Yaakov:Young:-.--
Z:Zulu:Zebra:Zebra:Zebra:Zebra:Zebra:Zebra:Zulmira:Zebedee:Zebra:--.."

# Escolhe o alfabeto a ser utilizado
case "$1" in
	--militar | --radio | --fone | --telefone | --otan | --nato | --icao | --itu | --imo | --faa | --ansi)
		coluna=2 ; shift ;;
	--romano | --latino           ) coluna=1  ; shift ;;
	--royal | --royal-navy        ) coluna=3  ; shift ;;
	--signalese | --western-front ) coluna=4  ; shift ;;
	--raf24                       ) coluna=5  ; shift ;;
	--raf42                       ) coluna=6  ; shift ;;
	--raf43 | --raf               ) coluna=7  ; shift ;;
	--us41 | --us                 ) coluna=8  ; shift ;;
	--pt | --portugal             ) coluna=9  ; shift ;;
	--name | --names              ) coluna=10 ; shift ;;
	--lapd                        ) coluna=11 ; shift ;;
	--morse                       ) coluna=12 ; shift ;;
esac

if test "$1"
then
	# Texto informado, vamos fazer a conversão
	# Deixa uma letra por linha e procura seu código equivalente
	echo "$*" |
		zzmaiusculas |
		sed 's/./&\ /g' |
		while IFS='' read -r char
		do
			letra=$(echo "$char" | sed 's/[^A-Z]//g')
			if test -n "$letra"
			then
				echo "$dados" | grep "^$letra" | cut -d : -f $coluna
			else
				test -n "$char" && echo "$char"
			fi
		done
else
	# Apenas mostre a tabela
	echo "$dados" | cut -d : -f $coluna
fi
}
```
## Exemplo de Implementação

```go
package main

import (
	"os"
	"strings"

	"github.com/paulopraxedes/estudos/go/goalfabeto"
)
/*
 * O programa deverá funcionar da seguinte forma:
    > O usuário poderá consultar toda a tabela de alfabetos = go run main.go
    > O usuário poderá consultar a tabela de um específico alfabeto = go run main.go <tipo_alfabeto>
    > O usuário poderá consultar a codificação resultante, baseado no tipo do alfabeto e a palavra a ser codificada = go run main.go <tipo_alfabeto> <nome_a_ser_codificado>
*/
func main() {
	if len(os.Args) == 1 { // Sem parâmetro
		goalfabeto.MostraTabela()
	} else if len(os.Args) == 2 { // <tipo_alfabeto>
		goalfabeto.MostraAlfabeto(strings.ToLower(os.Args[1]))
	} else if len(os.Args) == 3 { // <tipo_alfabeto> <nome_a_ser_codificado>
		goalfabeto.MostraAlfabetoFormatado(strings.ToLower(os.Args[1]), strings.ToUpper(os.Args[2]))
	}
}
```
---
O Projeto ZZ está disponível no [GitHub](https://github.com/funcoeszz/funcoeszz).
Para contribuições, leiam o [README.md](https://github.com/funcoeszz/funcoeszz/blob/master/README.md)
