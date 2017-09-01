package main

import "fmt"

//var dados = map[string][]string{"A": {"a", "b"}}

var dados = map[string][]string{
	"A": {"Alpha", "Apples", "Ack", "Ace", "Apple", "Able/Affirm", "Able", "Aveiro", "Alan", "Adam", ".-", "Anton"},
	"B": {"Bravo", "Butter", "Beer", "Beer", "Beer", "Baker", "Baker", "Bragança", "Bobby", "Boy", "-...", "Berta"},
	"C": {"Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Charlie", "Coimbra", "Charlie", "Charles", "-.-.", "Casar"},
	"D": {"Delta", "Duff", "Don", "Don", "Dog", "Dog", "Dog", "Dafundo", "David", "David", "-..", "Dora"},
	"E": {"Echo", "Edward", "Edward", "Edward", "Edward", "Easy", "Easy", "Évora", "Edward", "Edward", ".", "Emil"},
	"F": {"Foxtrot", "Freddy", "Freddie", "Freddie", "Freddy", "Fox", "Fox", "Faro", "Frederick", "Frank", "..-.", "Friedrich"},
	"G": {"Golf", "George", "Gee", "George", "George", "George", "George", "Guarda", "George", "George", "--.", "Gustav"},
	"H": {"Hotel", "Harry", "Harry", "Harry", "Harry", "How", "How", "Horta", "Howard", "Henry", "....", "Heinrich"},
	"I": {"India", "Ink", "Ink", "Ink", "In", "Item/Interrogatory", "Item", "Itália", "Isaac", "Ida", "..", "Ida"},
	"J": {"Juliet", "Johnnie", "Johnnie", "Johnnie", "Jug/Johnny", "Jig/Johnny", "Jig", "José", "James", "John", ".---", "Julius"},
	"K": {"Kilo", "King", "King", "King", "King", "King", "King", "Kilograma", "Kevin", "King", "-.-", "Kaufmann/Konrad"},
	"L": {"Lima", "London", "London", "London", "Love", "Love", "Love", "Lisboa", "Larry", "Lincoln", ".-..", "Ludwig"},
	"M": {"Mike", "Monkey", "Emma", "Monkey", "Mother", "Mike", "Mike", "Maria", "Michael", "Mary", "--", "Martha"},
	"N": {"November", "Nuts", "Nuts", "Nuts", "Nuts", "Nab/Negat", "Nan", "Nazaré", "Nicholas", "Nora", "-.", "Nordpol"},
	"O": {"Oscar", "Orange", "Oranges", "Orange", "Orange", "Oboe", "Oboe", "Ovar", "Oscar", "Ocean", "---", "Otto"},
	"P": {"Papa", "Pudding", "Pip", "Pip", "Peter", "Peter/Prep", "Peter", "Porto", "Peter", "Paul", ".--.", "Paula"},
	"Q": {"Quebec", "Queenie", "Queen", "Queen", "Queen", "Queen", "Queen", "Queluz", "Quincy", "Queen", "--.-", "Quelle"},
	"R": {"Romeo", "Robert", "Robert", "Robert", "Roger/Robert", "Roger", "Roger", "Rossio", "Robert", "Robert", ".-.", "Richard"},
	"S": {"Sierra", "Sugar", "Esses", "Sugar", "Sugar", "Sugar", "Sugar", "Setúbal", "Stephen", "Sam", "...", "Samuel/Siegfried"},
	"T": {"Tango", "Tommy", "Toc", "Toc", "Tommy", "Tare", "Tare", "Tavira", "Trevor", "Tom", "-", "Theodor"},
	"U": {"Uniform", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Uncle", "Unidade", "Ulysses", "Union", "..-", "Ulrich"},
	"V": {"Victor", "Vinegar", "Vic", "Vic", "Vic", "Victor", "Victor", "Viseu", "Vincent", "Victor", "...-", "Viktor"},
	"W": {"Whiskey", "Willie", "William", "William", "William", "William", "William", "Washington", "William", "William", ".--", "Wilhelm"},
	"X": {"X-ray/Xadrez", "Xerxes", "X-ray", "X-ray", "X-ray", "X-ray", "X-ray", "Xavier", "Xavier", "X-ray", "-..-", "Xanthippe/Xavier"},
	"Y": {"Yankee", "Yellow", "Yorker", "Yorker", "Yoke/Yorker", "Yoke", "Yoke", "York", "Yaakov", "Young", "-.--", "Ypsilon"},
	"Z": {"Zulu", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zebra", "Zulmira", "Zebedee", "Zebra", "--..", "Zacharias/Zurich"},
}

/*

--militar | --radio | --fone | --telefone | --otan | --nato | --icao | --itu | --imo | --faa | --ansi)
				coluna=2 ; shift ;;
			--romano | --latino           ) coluna=1     ; shift ;;
			--royal | --royal-navy        ) coluna=3     ; shift ;;
			--signalese | --western-front ) coluna=4     ; shift ;;
			--raf24                       ) coluna=5     ; shift ;;
			--raf42                       ) coluna=6     ; shift ;;
			--raf43 | --raf               ) coluna=7     ; shift ;;
			--us41 | --us                 ) coluna=8     ; shift ;;
			--pt | --portugal             ) coluna=9     ; shift ;;
			--name | --names              ) coluna=10    ; shift ;;
			--lapd                        ) coluna=11    ; shift ;;
			--morse                       ) coluna=12    ; shift ;;
			--german                      ) coluna=13    ; shift ;;

*/




func main() {

	if len(os.Args)>1 {
		if 
	}




	fmt.Println(dados["C"][3])
}
