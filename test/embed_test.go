package test

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"testing"
)

//go:embed version.txt
var version string

// bikin test untuk Embed String
func TestString(t *testing.T) {
	// version diambil dari file version.txt
	// karena sudah diembed kedalam variable bernama version
	fmt.Println(version)
}

// kita bikin test Embed byte[]

//go:embed logo.png
var logo []byte

func TestByte(t *testing.T) {
	// kita coba isi data logo yang baru
	// nama_logo_baru, file binary, function
	err := ioutil.WriteFile("logo_new.png", logo, fs.ModePerm)

	if err != nil {
		panic(err)
	}
}

// harus embed yang tipenya FS (File System)
//go:embed files/a.txt
//go:embed files/b.txt
//go:embed files/c.txt
var files embed.FS

// gimaca cara baca file a, b, c? harus menggunakan file system yakni read file
func TestMultipleFiles(t *testing.T) {
	// balikannya 2 return value, yakni byte[] nya dan errornya
	a, _ := files.ReadFile("files/a.txt")
	fmt.Println(string(a)) // koversi ke string bisa byte bisa dibaca

	b, _ := files.ReadFile("files/b.txt")
	fmt.Println(string(b))

	c, _ := files.ReadFile("files/c.txt")
	fmt.Println(string(c))
}

// embed file dengan Path Matcher

//go:embed files/*.txt
var path embed.FS

func TestPathMatcher(t *testing.T) {
	// kita bisa langsung baca directory
	// isi dari dirEntries ini adalah []fs.DirEntry
	dirEntries, _ := path.ReadDir("files")

	// lalu kita interasi
	for _, entry := range dirEntries {
		// cek apakah dia directory/folder atau bukan, karena dia tidak tau entry itu apakah folder atau file
		// jika bukan directory
		if !entry.IsDir() {
			// ambil name nya, hasilnya adalah nama file (bukan directory)
			fmt.Println(entry.Name())
			// lalu baca isi filenya, tapi masukkan path directorynya
			file, _ := path.ReadFile("files/" + entry.Name())
			// cetak isi file ke string
			fmt.Println(string(file))
		}
	}

}
