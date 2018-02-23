package testing

import (
	"fmt"
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

func TestSum(t *testing.T) {
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

func TestMain(m *testing.M) {
	log.Println("Start tests")
	code := m.Run()
	log.Println("Stop tests")
	os.Exit(code)
}
