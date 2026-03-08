package main

import "testing"

func TestRuneCount(t *testing.T) {
	if got := runeCount("ação"); got != 4 {
		t.Fatalf("runeCount = %d, want 4", got)
	}
}

func TestHasUTF8BOM(t *testing.T) {
	with := []byte{0xEF, 0xBB, 0xBF, 'G', 'o'}
	without := []byte{'G', 'o'}
	if !hasUTF8BOM(with) {
		t.Fatal("expected BOM")
	}
	if hasUTF8BOM(without) {
		t.Fatal("did not expect BOM")
	}
}
