package files

import (
	"fmt"
	"os"
	"testing"
)

func TestFileSize(T *testing.T) {
	stat, err := os.Stat("README.md")
	if err != nil {
		panic(err)
	}
	fmt.Println(stat.Size())
}

func TestReadAt(T *testing.T) {
	file, _ := os.Open("README.md")
	buf := make([]byte, 100)
	file.ReadAt(buf, 1)
	fmt.Printf("%s", buf)
}
