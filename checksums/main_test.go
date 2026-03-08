package main

import "testing"

func TestChecksums(t *testing.T) {
	if CRC32([]byte("go")) == 0 {
		t.Fatal("unexpected zero crc32")
	}
	if CRC64([]byte("go")) == 0 {
		t.Fatal("unexpected zero crc64")
	}
}
