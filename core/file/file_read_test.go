package file

import (
	"bufio"
	"io"
	"os"
	"testing"
)

// 也许不同进程不能同时打开文件，但是多携程是可以同时打开一个文件的
func TestReadFile(t *testing.T) {
	// os.Open() = OpenFile(name, O_RDONLY, 0)
	filename := `C:/Users/14590/Desktop/abc.log`

	var (
		file  *os.File
		lines []string
		err   error
	)
	// 读模式
	if file, err = os.OpenFile(filename, os.O_RDONLY, 0); err != nil {
		t.Fatal(err)
	}
	// 读取
	if lines, err = read(file); err != nil {
		t.Fatal(err)
	}
	// 操作
	println(len(lines))
}

func read(file *os.File) ([]string, error) {
	var (
		reader = bufio.NewReader(file)

		lines = make([]string, 0)
		line  string

		err error
	)
	for {
		if line, err = reader.ReadString('\n'); err != nil && err != io.EOF {
			return nil, err
		}
		lines = append(lines, line)

		if err == io.EOF {
			return lines, nil
		}
	}
}
