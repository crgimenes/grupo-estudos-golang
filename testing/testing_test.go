package testing

import (
	"fmt"
	"log"
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
