package fzip

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

//NameReader структура необходимая для метода Zip.
type NameReader struct {
	//Имя файла
	Name string
	//Reader - содержимое файла
	Reader io.Reader
}

type listPathFiles struct {
	list []string
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

//ZipPath создание zip архива по массиву путей до файлов
func ZipPath(pathFiles []string, pathZip string, deleteFiles ...bool) error {
	listFiles := make([]*os.File, 0)

	listStringFiles := &listPathFiles{
		list: []string{},
	}

	for _, p := range pathFiles {
		file, err := os.Open(p)
		if err != nil {
			return err
		}
		defer file.Close()
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		if !fileInfo.IsDir() {
			listFiles = append(listFiles, file)
			continue
		}
		if err := getFilesInDir(p, listStringFiles); err != nil {
			return err
		}
	}

	if err := checkUniqueFilesName(listFiles); err != nil {
		return err
	}

	zipFile, err := os.Create(pathZip)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	listDirs, err := fillFilesFromStringDirs(listStringFiles.list)
	if err != nil {
		return err
	}

	if err := zipFilesFromDir(listFiles, listDirs, zipFile); err != nil {
		return err
	}

	//удаление файлов добавленных в архив
	if len(deleteFiles) > 0 && deleteFiles[0] {
		for _, f := range pathFiles {
			if err := os.RemoveAll(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func getFilesInDir(pathDir string, files *listPathFiles) error {
	filesInfo, err := ioutil.ReadDir(pathDir)
	if err != nil {
		return err
	}
	for _, f := range filesInfo {
		if !f.IsDir() {
			t := filepath.Join(pathDir, f.Name())
			files.list = append(files.list, t)
			continue
		}
		if err := getFilesInDir(filepath.Join(pathDir, f.Name()), files); err != nil {
			return err
		}
	}
	return nil
}

func checkUniqueFilesName(listFiles []*os.File) error {
	mapList := make(map[string]bool)
	for _, file := range listFiles {
		if mapList[filepath.Base(file.Name())] {
			return fmt.Errorf("file %v must be unique at the root of the archive",
				filepath.Base(file.Name()))
		}
		mapList[filepath.Base(file.Name())] = true
	}
	return nil
}

func fillFilesFromStringDirs(listPathDirs []string) ([]*os.File, error) {
	listDirs := make([]*os.File, 0)

	for _, f := range listPathDirs {
		file, err := os.Open(f)
		if err != nil {
			return nil, err
		}
		listDirs = append(listDirs, file)
	}
	return listDirs, nil
}

func zipFilesFromDir(listFiels []*os.File, listDirs []*os.File, zipFile io.Writer) error {
	writerZip := zip.NewWriter(zipFile)
	defer writerZip.Close()

	for _, file := range listFiels {
		f, err := writerZip.Create(filepath.Base(file.Name()))
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
	for _, file := range listDirs {
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
