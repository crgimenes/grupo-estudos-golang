package main

import (
	"fmt"
	"io/fs"
	"log"
	"testing/fstest"
)

func ListRegularFiles(fsys fs.FS, root string) ([]string, error) {
	files := []string{}
	err := fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.Type().IsRegular() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func FileSize(fsys fs.FS, path string) (int64, error) {
	info, err := fs.Stat(fsys, path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

func main() {
	demoFS := fstest.MapFS{
		"assets/logo.txt": {Data: []byte("go")},
		"docs/readme.md":  {Data: []byte("hello\n")},
	}

	files, err := ListRegularFiles(demoFS, ".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		size, err := FileSize(demoFS, file)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s %d bytes\n", file, size)
	}
}
