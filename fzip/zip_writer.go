package fzip

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
)

//ZipFiles архивирование zip архива по переданным файлам
func ZipFiles(files []*os.File, zipFile io.Writer) error {
	writerZip := zip.NewWriter(zipFile)
	defer writerZip.Close()

	for _, file := range files {
		f, err := writerZip.Create(file.Name())
		if err != nil {
			return err
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		_, err = f.Write(data)
		if err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}
	}
	return nil
}
