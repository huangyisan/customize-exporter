package files

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	flag "github.com/spf13/pflag"
	"io"
	"os"
	"time"
)

var (
	fileName      string
	file          File
	fileIsChanged prometheus.Gauge
)

const (
	MAXFILELENGTH = 1024 * 1024 * 1024
	OFFSET        = 64
	TICK          = 60
)

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

func Do() {
	if fileName == "" {
		fmt.Println("Use -f to specify a file")
		os.Exit(1)
	}
	file.name = fileName
	fileIsChanged = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "file_change_detect",
		Help: fmt.Sprintf("The status code of file [%s] change events. change 1, keep 0", fileName),
	})
	go func() {
		c := time.Tick(TICK * time.Second)
		for range c {
			status := file.isFileChange()
			if status {
				fileIsChanged.Set(1)
			} else {
				fileIsChanged.Set(0)
			}
		}
	}()

}
