package goalfabeto

import (
	"fmt"
	"sort"
)

/**
 * mapDados = Objeto do tipo Mapa, que representa os alfabetos possíveis
 */
var mapDados = map[int][]string{
	0:  {"A", "Alpha", "Apples", "Ack", "Ace", "Apple", "Able/Affirm", "Able", "Aveiro", "Alan", "Adam", ".-"},
	1:  {"B", "Bravo", "Butter", "Beer", "Beer", "Beer", "Baker", "Baker", "Bragança", "Bobby", "Boy", "-..."},
	2:  {"C", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Coimbra", "Charlie", "Charles", "-.-."},
	3:  {"D", "Delta", "Duff", "Don", "Don", "Dog", "Dog", "Dog", "Dafundo", "David", "David", "-.."},
	4:  {"E", "Echo", "Edward", "Edward", "Edward", "Edward", "Easy", "Easy", "Évora", "Edward", "Edward", "."},
	5:  {"F", "Foxtrot", "Freddy", "Freddie", "Freddie", "Freddy", "Fox", "Fox", "Faro", "Frederick", "Frank", "..-."},
	6:  {"G", "Golf", "George", "Gee", "George", "George", "George", "George", "Guarda", "George", "George", "--."},
	7:  {"H", "Hotel", "Harry", "Harry", "Harry", "Harry", "How", "How", "Horta", "Howard", "Henry", "...."},
	8:  {"I", "India", "Ink", "Ink", "Ink", "In", "Item/Interrogatory", "Item", "Itália", "Isaac", "Ida", ".."},
	9:  {"J", "Juliet", "Johnnie", "Johnnie", "Johnnie", "Jug/Johnny", "Jig/Johnny", "Jig", "José", "James", "John", ".---"},
	10: {"K", "Kilo", "King", "King", "King", "King", "King", "King", "Kilograma", "Kevin", "King", "-.-"},
	11: {"L", "Lima", "London", "London", "London", "Love", "Love", "Love", "Lisboa", "Larry", "Lincoln", ".-.."},
	12: {"M", "Mike", "Monkey", "Emma", "Monkey", "Mother", "Mike", "Mike", "Maria", "Michael", "Mary", "--"},
	13: {"N", "November", "Nuts", "Nuts", "Nuts", "Nuts", "Nab/Negat", "Nan", "Nazaré", "Nicholas", "Nora", "-."},
	14: {"O", "Oscar", "Orange", "Oranges", "Orange", "Orange", "Oboe", "Oboe", "Ovar", "Oscar", "Ocean", "---"},
	15: {"P", "Papa", "Pudding", "Pip", "Pip", "Peter", "Peter/Prep", "Peter", "Porto", "Peter", "Paul", ".--."},
	16: {"Q", "Quebec", "Queenie", "Queen", "Queen", "Queen", "Queen", "Queen", "Queluz", "Quincy", "Queen", "--.-"},
	17: {"R", "Romeo", "Robert", "Robert", "Robert", "Roger/Robert", "Roger", "Roger", "Rossio", "Robert", "Robert", ".-."},
	18: {"S", "Sierra", "Sugar", "Esses", "Sugar", "Sugar", "Sugar", "Sugar", "Setúbal", "Stephen", "Sam", "..."},
	19: {"T", "Tango", "Tommy", "Toc", "Toc", "Tommy", "Tare", "Tare", "Tavira", "Trevor", "Tom", "-"},
	20: {"U", "Uniform", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Unidade", "Ulysses", "Union", "..-"},
	21: {"V", "Victor", "Vinegar", "Vic", "Vic", "Vic", "Victor", "Victor", "Viseu", "Vincent", "Victor", "...-"},
	22: {"W", "Whiskey", "Willie", "William", "William", "William", "William", "William", "Washington", "William", "William", ".--"},
	23: {"X", "X-ray/Xadrez", "Xerxes", "X-ray", "X-ray", "X-ray", "X-ray", "X-ray", "Xavier", "Xavier", "X-ray", "-..-"},
	24: {"Y", "Yankee", "Yellow", "Yorker", "Yorker", "Yoke/Yorker", "Yoke", "Yoke", "York", "Yaakov", "Young", "-.--"},
	25: {"Z", "Zulu", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zulmira", "Zebedee", "Zebra", "--.."},
}

/**
ordenaTabela - função que ordena a lista, de modo crescente, para uma melhor leitura
*/
func ordenaTabela(tabela map[int][]string) (keys []int) {
	for valor := range tabela {
		keys = append(keys, valor)
	}
	sort.Ints(keys)
	return
}

/*MostraTabela - função que imprime a tabela gerada.
Tendo uma única coluna, imprime o valor do índice
Tendo mais que uma coluna, imprime o slice inteiro
*/
func MostraTabela(tabela map[int][]string) {
	if tabela == nil {
		tabela = mapDados
	}
	for valor := range ordenaTabela(tabela) {
		if len(tabela[valor]) > 1 {
			fmt.Printf("%+s\r\n", tabela[valor])
		} else {
			fmt.Printf("%+s\r\n", tabela[valor][0])
		}
	}
}

/**
populaMapa - função que popula um mapa reduzido com o resultado do mapa maior
*/
func populaMapa(tabela map[int][]string, index int) (mapa map[int][]string) {
	mapa = map[int][]string{}
	for key := range tabela {
		var vet []string
		vet = append(vet, tabela[key][index])
		mapa[key] = vet
	}
	return mapa
}

/**
convertePalavra - função que converte a palavra desejada ao alfabeto desejado
*/
func convertePalavra(tabela map[int][]string, palavra string) {
	for _, caractere := range palavra {
		for index := range mapDados {
			if string(caractere) == mapDados[index][0] {
				fmt.Printf("%s\r\n", tabela[index][0])
			}
		}
	}
}

/*MontaAlfabeto - função que monta o mapa com alfabeto desejado
 */
func MontaAlfabeto(tipo string, valor string) {
	tabela := map[int][]string{}
	switch tipo {
	case "--romano", "--latino":
		tabela = populaMapa(mapDados, 0)
	case "--militar", "--radio", "--fone", "--telefone", "--otan", "--nato", "--icao", "--itu", "--imo", "--faa", "--ansi":
		tabela = populaMapa(mapDados, 1)
	case "--royal", "--royal-navy":
		tabela = populaMapa(mapDados, 2)
	case "--signalese", "--western-front":
		tabela = populaMapa(mapDados, 3)
	case "--raf24":
		tabela = populaMapa(mapDados, 4)
	case "--raf42":
		tabela = populaMapa(mapDados, 5)
	case "--raf43", "--raf":
		tabela = populaMapa(mapDados, 6)
	case "--us41", "--us":
		tabela = populaMapa(mapDados, 7)
	case "--pt", "--portugal":
		tabela = populaMapa(mapDados, 8)
	case "--name", "--names":
		tabela = populaMapa(mapDados, 9)
	case "--lapd":
		tabela = populaMapa(mapDados, 10)
	case "--morse":
		tabela = populaMapa(mapDados, 11)
	}

	if valor != "" {
		convertePalavra(tabela, valor)
	} else {
		MostraTabela(tabela)
	}
}

//GETMap - função GET que retorna a tabela de alfabetos disponíveis
func GETMap() map[int][]string {
	return mapDados
}
