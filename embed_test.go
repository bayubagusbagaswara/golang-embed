package golangembed

import (
	_ "embed"
	"fmt"
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
