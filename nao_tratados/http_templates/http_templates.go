package http_templates

import (
	"fmt"
	"net/http"
	"html/template"
)

//Criamos um objeto para representar os dados que queremos exibir na página
type Page struct {
	Name string
}


func New(){main()}

func main() {
	//Utilizamos um manipulador da biblioteca html/template para realizar o parse do nosso template 
	templates := template.Must(template.ParseFiles("templates/index.html"))


	//Utilizamos outro manipulador para tratar os eventos de request e response
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
	
		//Criamos um objeto que será passado como parâmetro para nosso template
		p := Page{Name : "GopherBrasil"}
		
		//Associamos um elemento de formulário a este objeto. Se o name não estiver vazio, atualiza a página com o valor estipulado
		if name := r.FormValue("name"); name != ""{
			p.Name = name;
		}
	
		//Dentro do nosso manipulador vamos tratar os erros relacionados ao template
		//ExecuteTemplate irá executar diretamente nosso template 
		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			//Alerta o usuário para os erros encontrados através da mensagem e código padrão do erro
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	
	})
	
	//realizando o setup para ouvir na porta 8080
	fmt.Println(http.ListenAndServe(":8080", nil))

}
