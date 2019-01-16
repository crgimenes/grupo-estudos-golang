package simples

import "testing"

func TestRetornaSempreTrue(t *testing.T) {
	if !RetornaSempreTrue() {
		t.Fatal("RetornaSempreTrue devia retornar true mas retornou false!")
	}
}
