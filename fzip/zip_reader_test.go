package fzip

import (
	"os"
	"fmt"
	"testing"
)

func TestUnZipFile(t *testing.T) {
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

	r, err := UnZipFile(fileZip)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(r)
}
