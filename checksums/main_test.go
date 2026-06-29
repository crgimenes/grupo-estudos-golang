package main

import "testing"

func TestChecksums(t *testing.T) {
	payload := []byte("invoice:42:paid")

	if got := CRC32(payload); got != 0xd57bb496 {
		t.Fatalf("CRC32() = %08x, want d57bb496", got)
	}

	if got := CRC64(payload); got != 0x5ca2c9f6a56b90c2 {
		t.Fatalf("CRC64() = %016x, want 5ca2c9f6a56b90c2", got)
	}

	if got := SHA256(payload); got != "5caa2109d58ce3fefe483e8b7176b80a00485ecbb3d92b18e48204a6ed4fe876" {
		t.Fatalf("SHA256() = %s", got)
	}
}
