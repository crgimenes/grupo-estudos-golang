package main

import (
	"archive/zip"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func GzipBytes(data []byte) ([]byte, error) {
	var out bytes.Buffer
	zw := gzip.NewWriter(&out)

	_, err := zw.Write(data)
	if err != nil {
		return nil, err
	}

	err = zw.Close()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func UngzipBytes(data []byte) ([]byte, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	out, readErr := io.ReadAll(zr)
	closeErr := zr.Close()
	if readErr != nil {
		return nil, readErr
	}
	if closeErr != nil {
		return nil, closeErr
	}

	return out, nil
}

func ZipSingleFile(name string, content []byte) ([]byte, error) {
	var out bytes.Buffer
	zw := zip.NewWriter(&out)

	w, err := zw.Create(name)
	if err != nil {
		return nil, err
	}

	_, err = w.Write(content)
	if err != nil {
		return nil, err
	}

	err = zw.Close()
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func ZipFileNames(data []byte) ([]string, error) {
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(zr.File))
	for _, file := range zr.File {
		names = append(names, file.Name)
	}

	return names, nil
}

func main() {
	payload := []byte("go study group\n")

	compressed, err := GzipBytes(payload)
	if err != nil {
		panic(err)
	}

	restored, err := UngzipBytes(compressed)
	if err != nil {
		panic(err)
	}

	archive, err := ZipSingleFile("note.txt", payload)
	if err != nil {
		panic(err)
	}

	names, err := ZipFileNames(archive)
	if err != nil {
		panic(err)
	}

	fmt.Printf("gzip roundtrip: %q\n", restored)
	fmt.Printf("gzip bytes: %d -> %d\n", len(payload), len(compressed))
	fmt.Printf("zip files: %v\n", names)
}
