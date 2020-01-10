package fzip

import (
	"bytes"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	fileZip, err := os.Create("test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer fileZip.Close()
	defer os.Remove("test.zip")

	var r1 bytes.Buffer
	var r2 bytes.Buffer

	_, err = r1.Write([]byte("Данные первого файла"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r2.Write([]byte("Reader второго файла"))
	if err != nil {
		t.Fatal(err)
	}
	listNameReader := make([]*NameReader, 0)
	listNameReader = append(listNameReader, &NameReader{"1.txt", &r1})
	listNameReader = append(listNameReader, &NameReader{"2.txt", &r2})

	if err := Zip(listNameReader, fileZip); err != nil {
		t.Fatal(err)
	}
}

func TestZipFiles(t *testing.T) {
	fileZip, err := os.Create("test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer fileZip.Close()
	defer os.Remove("test.zip")

	f1, err := os.OpenFile("1.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.OpenFile("2.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	f3, err := os.OpenFile("3.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("1.txt")
	defer os.Remove("2.txt")
	defer os.Remove("3.txt")

	if err := ZipFiles([]*os.File{f1, f2, f3}, fileZip); err != nil {
		t.Fatal(err)
	}
}


func TestZipPath(t *testing.T) {
	fileZip, err := os.Create("test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer fileZip.Close()
	//defer os.Remove("test.zip")

	_, err = os.OpenFile("1.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.OpenFile("2.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.OpenFile("3.txt", os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("1.txt")
	defer os.Remove("2.txt")
	defer os.Remove("3.txt")

	if err := ZipPath([]string{"1.txt", "2.txt", "3.txt"}, "test.zip", false); err != nil {
		t.Fatal(err)
	}
}
