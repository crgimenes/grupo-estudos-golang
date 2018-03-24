# Testes

> Todo arquivo de testes deve ter o sufixo `_test.go` para o `go test` (ferramenta do go pra executar testes) enxergar o arquivo e suas
funções devem ter a assinatura `func Test...(t *testing.T)` para serem consideradas testes.

## Exemplo

- Função a ser testada:

```go
package testing

import (
	"errors"
)

var errDivisaoInvalida = errors.New("divisão invalida")

func divideInteiros(dividendo, divisor int) (quociente int, resto int, err error) {
	if divisor == 0 {
		err = errDivisaoInvalida
		return
	}
	quociente = dividendo / divisor
	resto = dividendo % divisor
	return
}
```

- Teste:

```go
func Test_divideInteiros(t *testing.T) {
	type args struct {
		dividendo int
		divisor   int
	}
	type expected struct {
		quociente int
		resto     int
		err       error
	}
	tests := []struct {
		name string
		args args
		want expected
	}{
		// Casos de teste para a função
		{
			name: "divide por zero",
			want: expected{
				err: errDivisaoInvalida,
			},
			args: args{
				dividendo: 10,
				divisor:   0,
			},
		},
		{
			name: "divisão sem resto",
			want: expected{
				err:       nil,
				resto:     0,
				quociente: 5,
			},
			args: args{
				dividendo: 10,
				divisor:   2,
			},
		},
		{
			name: "divisão com resto",
			want: expected{
				err:       nil,
				resto:     1,
				quociente: 3,
			},
			args: args{
				dividendo: 7,
				divisor:   2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQuociente, gotResto, err := divideInteiros(tt.args.dividendo, tt.args.divisor)
			if err != tt.want.err {
				t.Errorf("divideInteiros() error = %v, wantErr %v", err, tt.want.err)
				return
			}
			if gotQuociente != tt.want.quociente {
				t.Errorf("divideInteiros() gotQuociente = %v, want %v", gotQuociente, tt.want.quociente)
			}
			if gotResto != tt.want.resto {
				t.Errorf("divideInteiros() gotResto = %v, want %v", gotResto, tt.want.resto)
			}
		})
	}
}
```

> Como podemos ver no exemplo, em Go só precisamos descrever nos testes os casos de falha, se algum caso de falha for satisfeito o código entrará no if e o teste falhará

## TestMain

> É possivel criar uma função `Main` para nossos testes com isso conseguimos testar recursos globais de nossa aplicação e criar um `setup`e um `teardown`global para nossa base de testes.

### Exemplo

```go
func TestMain(m *testing.M) {
	log.Println("Start tests")
	code := m.Run()
	log.Println("Stop tests")
	os.Exit(code)
}
```

---

## Mock

Existem algumas definições para *mock* mas a que mais se encaixa no que queremos é essa daqui:

> *adjective*
> 1. not authentic or real, but without the intention to deceive

Criar um mocks para simular situações e partes de código é uma parte importante dos testes e aqui vamos ver algumas formas de fazer isso.

### Mock usando interfaces

#### io.Reader

É normal quando estamos trabalhando com serviços receber um reader que vamos ler como um array de bytes, felizmente o Go prove uma interface pronta para isso.

```go
func leitor(r io.Reader) (ret string)  {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	ret = buf.String()
	return
}
```

Como `r` é uma interface podemos passar qualquer elemento que implemente a interface *io.Reader*.

```go
r := bytes.NewReader([]byte("hello world"))
```

#### io.ReaderCloser

Vamos ver novamente o exemplo anterior mas agora com a interface *io.ReadCloser*

```go
func leEFecha(r io.ReadCloser) (ret string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// note que agora alem de ler
	// vamos fechar o descritor
	r.Close()

	ret = buf.String()
	return
}
```

E precisamos implementar a interface de acordo, para isso no pacote *ioutil* a biblioteca padrão já fornece uma interface que faz mock do closer que a nossa simples string não vai implementar. Veja como fica:

```go
r := ioutil.NopCloser(bytes.NewReader([]byte("hello world")))
```

#### Mock http.ResponseWriter

Muito parecido com o exemplo anterior podemos querer fazer mock de alguma resposta que enviamos via http para o cliente. Veja a função abaixo:

```go
func responde(w http.ResponseWriter) {
	ret := struct {
		Msg string `json:"message"`
	}{
		Msg: "algo muito importante",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret) // nolint
}
```

Como a função recebe a interface *http.ResponseWriter* podemos passar qualquer interface que implemente as funções necessárias e o pacote httptest já fornece uma implementação para usarmos.

```go
w := httptest.NewRecorder()
```

#### Mock de servidor

O pacote httptest também uma forma de fazer mock de servidores dessa forma podemos facilmente testar os nossos clientes.

```go
	serverOk := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}))
	defer serverOk.Close()
```

E então basta passar para o cliente que esta sendo testado a URL do servidor de testes que no caso do nosso exemplo esta em `serverOk.URL`


#### Mock usando interfaces

Um exemplo mais completo mas ainda simples de mock usando interfaces é o teste da função *closer()*

```go
func closer(body io.Closer) {
	err := body.Close()
	if err != nil {
		log.Errorln(err)
	}
}
```

Criamos duas interfaces, uma para o caso de erro e outra para sucesso e então testamos cada uma.

```go
type closerSuccess struct {
}

func (c closerSuccess) Close() (err error) {
	return
}

type closerError struct {
}

func (c closerError) Close() (err error) {
	err = errors.New("closer error")
	return
}
```

Como closer não retorna nada a unica forma de validar seu funcionamento é capturando o stdout, existe um exemplo muito completo de teste capturando o stdout no pacote nuveo/log em http://github.com/nuveo/log

Veja como ficou o teste

```go
func Test_closer(t *testing.T) {

	getStdout := func(obj io.Closer) (out []byte, err error) {
		rescueStdout := os.Stdout
		defer func() { os.Stdout = rescueStdout }()
		r, w, err := os.Pipe()
		if err != nil {
			return nil, err
		}
		os.Stdout = w

		closer(obj)

		err = w.Close()
		if err != nil {
			return
		}
		out, err = ioutil.ReadAll(r)
		return
	}

	cs := closerSuccess{}
	ce := closerError{}

	type args struct {
		body io.Closer
	}
	type expected struct {
		err bool
	}
	tests := []struct {
		name string
		args args
		want expected
	}{
		{
			name: "success",
			args: args{
				body: cs,
			},
			want: expected{
				err: false,
			},
		},
		{
			name: "error",
			args: args{
				body: ce,
			},
			want: expected{
				err: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := getStdout(tt.args.body)
			if err != nil {
				t.Error(err)
				return
			}

			if (len(out) > 0) != tt.want.err {
				fmt.Printf("out: %q\n", string(out))
				t.Errorf("closer() unexpected log %q", string(out))
			}
		})
	}
}
```

---
[Inicio](../README.md)

[< tratando sinais](../signals/) - [plugin >](../plugin/)
