package files

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"os"
)

var fileName string
var file File

//const MAXFILELENGTH = 1024 * 1024 * 1024
const MAXFILELENGTH = 1
const OFFSET = 64

func init() {
	flag.StringVarP(&fileName, "file-name", "f", "", "specify file name")
}

type File struct {
	name        string
	fileMD5     string
	lastLineMD5 string
}

func (f *File) readFileLastLine() string {
	var buf []byte
	var start int64
	file, err := os.Open(f.name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := os.Stat(f.name)
	if OFFSET > stat.Size() {
		buf = make([]byte, stat.Size())
		start = stat.Size()
	} else {
		buf = make([]byte, OFFSET)
		start = stat.Size() - OFFSET
	}

	file.ReadAt(buf, start)

	return fmt.Sprintf("%s\n", buf)

}

func (f *File) getLastLineMD5() string {
	data := []byte(f.readFileLastLine())
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func (f *File) getFileMD5() string {
	// return file md5

	file, err := os.Open(f.name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	h := md5.New()
	io.Copy(h, file)
	//file.Close()
	return hex.EncodeToString(h.Sum(nil))
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
	if size < MAXFILELENGTH {
		storeFileMD5 := f.fileMD5
		currentFileMD5 := f.getFileMD5()
		if storeFileMD5 != currentFileMD5 {
			f.fileMD5 = currentFileMD5
			return true
		}
		return false
	} else {
		storeLineMD5 := f.lastLineMD5
		currentLineMD5 := f.getLastLineMD5()
		if storeLineMD5 != currentLineMD5 {
			f.lastLineMD5 = currentLineMD5
			return true
		}
		return false
	}
}

func Do() bool {
	file.name = fileName
	return file.isFileChange()
}
