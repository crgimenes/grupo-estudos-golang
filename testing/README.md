# Testes

> Todo arquivo de testes deve ter o sufixo `_test.go` para o `go test` (ferramenta do go pra executar testes) enxergar o arquivo e suas
funções devem ter a assinatura `func Test...(t *testing.T)` para serem consideradas testes.

# Exemplo

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
package testing

import (
	"testing"
)

func TestDivideInteiros(t *testing.T) {
	for _, test := range []struct {
		dividendo int
		divisor   int
		quociente int
		resto     int
		err       error
	}{
		{10, 0, 0, 0, errDivisaoInvalida},
		{10, 2, 5, 0, nil},
		{7, 2, 3, 1, nil},
	} {
		q, r, erro := divideInteiros(test.dividendo, test.divisor)
		if q != test.quociente {
			t.Errorf("Esperava como quociente %d e obiteve %d\n", test.quociente, q)
		}
		if r != test.resto {
			t.Errorf("Esperava como resto %d e obiteve %d\n", test.resto, r)
		}
		if erro != test.err {
			t.Errorf("Esperava como err %v e obiteve %v\n", test.err, erro)
		}
	}
}
```

> Como podemos ver no exemplo, em Go só precisamos descrever nos testes os casos de falha, 
se algum caso de falha for satisfeito o código entrará no if e o teste falhará