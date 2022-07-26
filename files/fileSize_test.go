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
