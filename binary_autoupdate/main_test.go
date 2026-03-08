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

func TestCompareVersionsInvalid(t *testing.T) {
	if _, err := CompareVersions("1.0", "1.0.1"); err == nil {
		t.Fatal("expected invalid version error")
	}
}
