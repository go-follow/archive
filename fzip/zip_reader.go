package fzip

import (
	"archive/zip"
	"os"
)

func UnZipFile(f *os.File) (*zip.Reader, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return zip.NewReader(f, fi.Size())
}
