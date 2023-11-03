package zip

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
)

// Archive keeps data to be zipped in RAM
type Archive struct {
	zw   *zip.Writer
	buff *bytes.Buffer
}

// New creates a new instance to control
// the data that will be zipped.
func New() *Archive {
	buff := &bytes.Buffer{}
	zw := zip.NewWriter(buff)
	return &Archive{
		zw:   zw,
		buff: buff,
	}
}

// Add new files to the buffer
func (a *Archive) Add(
	fileName string,
	content []byte) error {
	var w io.Writer
	w, err := a.zw.Create(fileName)
	if err != nil {
		return err
	}
	_, err = w.Write(content)
	return err
}

// Save the buffer with compressed data to a zip file
func (a *Archive) Save(zipFile string) error {
	err := a.zw.Close()
	if err != nil {
		return err
	}
	return os.WriteFile(zipFile, a.buff.Bytes(), 0600)
}
