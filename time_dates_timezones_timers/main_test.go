package main

import (
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil {
		t.Fatal(err)
	}
	v, err := ParseDate("2006-01-02 15:04", "2026-03-08 10:00", loc)
	if err != nil {
		t.Fatal(err)
	}
	if v.Location().String() != "America/Sao_Paulo" {
		t.Fatalf("location = %s", v.Location())
	}
}

func TestNewTimerDone(t *testing.T) {
	select {
	case <-NewTimerDone(2 * time.Millisecond):
	case <-time.After(100 * time.Millisecond):
		t.Fatal("timer timeout")
	}
}
