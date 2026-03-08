package main

import "testing"

func TestIsNewerVersion(t *testing.T) {
	if !IsNewerVersion("1.0.0", "1.0.1") {
		t.Fatal("expected update")
	}
	if IsNewerVersion("1.0.0", "1.0.0") {
		t.Fatal("did not expect update")
	}
}
