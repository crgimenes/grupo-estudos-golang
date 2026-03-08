package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func ListFiles(root string) ([]string, error) {
	out := []string{}
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		out = append(out, filepath.Base(path))
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func ReadFileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}
