package zip

import (
	"archive/zip"
	"io"
	"os"
)

// Archive file descriptor pointing to the zip file
type Archive struct {
	zw *zip.Writer
	f  *os.File
}

// New creates a new instance to control
// the data that will be zipped.
func New(fileName string) (*Archive, error) {
	f, err := os.Create(fileName)
	zw := zip.NewWriter(f)
	return &Archive{
		zw: zw,
		f:  f,
	}, err
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

// Close the zip file descriptor
func (a *Archive) Close() error {
	err := a.zw.Close()
	if err != nil {
		return err
	}
	return a.f.Close()
}
