package fzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

//UnZipFile разархивирование архива zip из файла
func UnZipFile(f *os.File) (*zip.Reader, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return zip.NewReader(f, fi.Size())
}

//UnZipPath разархивирование архива zip по пути до файла.
//deleteZip не обязательный параметр, если равен true - до архив после разархивирования удаляется
func UnZipPath(zipFile string, unZipDir string, deleteZip ...bool) error {
	f, err := os.Open(zipFile)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	r, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return err
	}
	if err := readZip(r, unZipDir); err != nil {
		return err
	}
	if len(deleteZip) > 0 && deleteZip[0] {
		if err := os.Remove(zipFile); err != nil {
			return err
		}
	}
	return nil
}

func readZip(r *zip.Reader, unZipDir string) error {
	if r == nil {
		return fmt.Errorf("не иницилизированный Reader")
	}
	for _, f := range r.File {
		if err := os.MkdirAll(unZipDir, 0766); err != nil {
			return err
		}
		if f.FileInfo().IsDir() {
			continue
		}
		nf, err := os.OpenFile(filepath.Join(unZipDir, f.Name),
			os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer nf.Close()
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		_, err = io.Copy(nf, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
