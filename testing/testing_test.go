package testing

import (
	"log"
	"os"
	"testing"
)

func TestDivideInteiros(t *testing.T) {
	for _, test := range []struct {
		// Struct que define os dados de entrada e saida necessarios para os testes
		dividendo int
		divisor   int
		quociente int
		resto     int
		err       error
	}{
		// Casos de teste para a função
		{10, 0, 0, 0, errDivisaoInvalida},
		{10, 2, 5, 0, nil},
		{7, 2, 3, 1, nil},
	} {
		// faz interação sobre os casos de testes
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

func TestMain(m *testing.M) {
	log.Println("Start tests")
	code := m.Run()
	log.Println("Stop tests")
	os.Exit(code)
}
