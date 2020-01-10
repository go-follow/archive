package fzip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//NameReader структура необходимая для метода Zip.
type NameReader struct {
	//Имя файла
	Name string
	//Reader - содержимое файла
	Reader io.Reader
}

//Zip создание zip архива по массиву NameReader
func Zip(listNameReader []*NameReader, w io.Writer) error {
	zipWriter := zip.NewWriter(w)
	if listNameReader == nil {
		return fmt.Errorf("empty slice NameReader")
	}

	for _, f := range listNameReader {
		if f == nil || f.Reader == nil || f.Name == "" {
			continue
		}
		zipEntry, err := zipWriter.Create(f.Name)
		if err != nil {
			return err
		}
		_, err = io.Copy(zipEntry, f.Reader)
		if err != nil {
			return err
		}
	}
	if err := zipWriter.Close(); err != nil {
		return err
	}
	return nil
}

//ZipFiles создание zip архива по переданным файлам
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
