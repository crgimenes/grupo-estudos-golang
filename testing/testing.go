package testing

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
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

func sum(a, b int) (ret int) {
	ret = a + b
	return
}

func leitor(r io.Reader) (ret string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r) // nolint
	ret = buf.String()
	return
}

func leEFecha(r io.ReadCloser) (ret string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)

	// note que agora alem de ler
	// vamos fechar o descritor
	r.Close()

	ret = buf.String()
	return
}

func responde(w http.ResponseWriter) {
	ret := struct {
		Msg string `json:"message"`
	}{
		Msg: "algo muito importante",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ret) // nolint
}
