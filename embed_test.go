package golangembed

import (
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
