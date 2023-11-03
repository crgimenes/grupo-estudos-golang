package testing

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Start tests")
	code := m.Run()
	log.Println("Stop tests")
	os.Exit(code)
}

func Test_sum(t *testing.T) {
	type args struct {
		a int
		b int
	}
	type expected struct {
		ret int
	}
	tests := []struct {
		name     string
		args     args
		want     expected
		setup    func()
		tearDown func()
	}{
		{
			name: "test 1+1",
			args: args{
				a: 1,
				b: 1,
			},
			want: expected{
				ret: 2,
			},
			setup: func() {
				fmt.Println("setup do teste 1+1")
			},
			tearDown: func() {
				fmt.Println("tearDown do teste 1+1")
			},
		},
		{
			name: "test 2+2",
			args: args{
				a: 2,
				b: 2,
			},
			want: expected{
				ret: 4,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup()
			}
			ret := sum(tt.args.a, tt.args.b)
			if ret != tt.want.ret {
				t.Errorf("sum() = %v, want %v", ret, tt.want.ret)
			}
			if tt.tearDown != nil {
				tt.tearDown()
			}
		})
	}
}

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

func Test_leitor(t *testing.T) {
	expected := "hello world"

	r := bytes.NewReader([]byte(expected))

	got := leitor(r)
	if got != expected {
		t.Errorf("leitor() = %v, want %v", got, expected)
	}
}

func Test_leEFecha(t *testing.T) {
	expected := "hello world"

	r := io.NopCloser(bytes.NewReader([]byte(expected)))

	got := leEFecha(r)
	if got != expected {
		t.Errorf("leEFecha() = %v, want %v", got, expected)
	}
}

func Test_responde(t *testing.T) {
	w := httptest.NewRecorder()
	expected := "{\"message\":\"algo muito importante\"}\n"

	responde(w)
	b := w.Body.Bytes()

	if string(b) != expected {
		t.Errorf("responde(): expected %q, but got %q", expected, string(b))
	}

	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("responde(): expected Content-Type application/json, but got %q", w.Header().Get("Content-Type"))
	}
}

func Test_clienteHttp(t *testing.T) {
	serverOk := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintln(w, "ok")
			}))
	defer serverOk.Close()

	serverErr := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "error", http.StatusInternalServerError)
				fmt.Fprintln(w, "error")
			}))
	defer serverErr.Close()

	type args struct {
		token  string
		method string
		url    string
		body   io.Reader
	}
	tests := []struct {
		name       string
		args       args
		wantStatus int
		wantErr    bool
	}{
		{
			name:       "StatusOK",
			wantErr:    false,
			wantStatus: http.StatusOK,
			args: args{
				token:  "test",
				method: http.MethodGet,
				url:    serverOk.URL,
			},
		},
		{
			name:       "StatusInternalServerError",
			wantErr:    false,
			wantStatus: http.StatusInternalServerError,
			args: args{
				token:  "test",
				method: http.MethodGet,
				url:    serverErr.URL,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := clienteHttp(tt.args.token, tt.args.method, tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("clienteHttp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResp.StatusCode != tt.wantStatus {
				t.Errorf("clienteHttp() = %v, want %v", gotResp.StatusCode, tt.wantStatus)
			}
		})
	}
}

/*
Testando a função closer() usando interfaces
*/

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
		out, err = io.ReadAll(r)
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
