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
