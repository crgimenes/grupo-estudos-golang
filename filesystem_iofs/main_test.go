package main

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestListRegularFiles(t *testing.T) {
	fsys := fstest.MapFS{
		"a.txt":       {Data: []byte("a")},
		"notes":       {Mode: fs.ModeDir},
		"notes/b.txt": {Data: []byte("b")},
	}

	files, err := ListRegularFiles(fsys, ".")
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{"a.txt", "notes/b.txt"}
	if !reflect.DeepEqual(files, expected) {
		t.Fatalf("files = %#v, want %#v", files, expected)
	}
}

func TestFileSize(t *testing.T) {
	fsys := fstest.MapFS{
		"message.txt": {Data: []byte("hello\n")},
	}

	size, err := FileSize(fsys, "message.txt")
	if err != nil {
		t.Fatal(err)
	}

	if size != 6 {
		t.Fatalf("size = %d, want 6", size)
	}
}

func TestFileSizeMissingFile(t *testing.T) {
	fsys := fstest.MapFS{}

	_, err := FileSize(fsys, "missing.txt")
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatalf("err = %v, want fs.ErrNotExist", err)
	}
}
