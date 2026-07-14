# Fundamentos de HTTP

Este tópico mostra o mínimo útil de HTTP no Go: um `http.Handler`, uma resposta JSON, headers e status codes.

A ideia aqui não é montar um framework caseiro. `net/http` já entrega o contrato principal: uma função recebe `*http.Request`, escreve em `http.ResponseWriter` e decide o status antes de mandar o corpo. Para muita API interna pequena, esse desenho aguenta mais tempo do que parece.

## Pré-requisitos

- Go instalado.
- Noção básica de funções e structs.

## Exemplo

O arquivo `main.go` tem um handler de health check:

```go
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(message{Status: "ok"})
}
```

Três detalhes importam mais do que o tamanho do código:

- o método HTTP é validado antes de escrever o corpo;
- o header `Content-Type` sai junto da resposta JSON;
- quando o método é recusado, o handler informa `Allow: GET`, que é o que um cliente decente espera ler num `405 Method Not Allowed`.

## Testar

Use `httptest` para testar o handler sem abrir porta local:

```bash
go test -timeout 30s -count 1 ./...
```

O teste cobre o happy path (`GET /health`) e o erro de método (`POST /health`). Isso força o exemplo a tratar status, header e corpo como parte do contrato, não como detalhe visual do navegador.

## Pontos de atenção

- Depois que o corpo começa a ser escrito, mudar status code já virou tarde demais.
- Header precisa ser definido antes da primeira escrita no `ResponseWriter`.
- `http.Error` escreve texto simples. Se a API inteira promete JSON, o erro também precisa seguir esse formato. Este exemplo não faz isso para não misturar middleware e negociação de erro no primeiro contato com HTTP.

## Próximos passos

- Adicionar um `http.ServeMux` com duas rotas pequenas.
- Testar query string com `r.URL.Query()`.
- Separar o formato de erro JSON quando o tópico avançar para APIs.
