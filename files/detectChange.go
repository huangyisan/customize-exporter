package files

import (
	"crypto/md5"
	"fmt"
	flag "github.com/spf13/pflag"
	"os"
)

var fileName string

func init() {
	flag.StringVar(&fileName, "fileName", "", "specify file name")
}

type File struct {
	name        string
	fileMD5     string
	lastLineMD5 string
}

func (f *File) readFileLastLine() string {
	file, err := os.Open(f.name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buf := make([]byte, 62)
	stat, err := os.Stat(f.name)
	start := stat.Size() - 62
	_, err = file.ReadAt(buf, start)
	if err == nil {
		return fmt.Sprintf("%s\n", buf)
	}
	return ""
}

func (f *File) getLastLineMD5() string {
	data := []byte(f.readFileLastLine())
	has := md5.Sum(data)
	md5sum := fmt.Sprintf("%x", has)
	return md5sum
}

func (f *File) getFileMD5() string {
	// return file md5
	return "cccssdadsfasdfa"
}

func (f *File) getFileSize() int64 {
	stat, err := os.Stat(f.name)
	if err != nil {
		panic(err)
	}
	return stat.Size()
}

func (f *File) isFileChange() bool {
	size := f.getFileSize()
	fmt.Println(size)
	return true
}

func Do() {

}
