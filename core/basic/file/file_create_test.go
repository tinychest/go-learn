package file

import (
	"fmt"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	createDirTest()
	createFileTest()
}

func createDirTest() {
	var err error

	// 创建参数中的最后一层文件夹
	// 前置路径对应的文件夹不存在，则创建失败 - mkdir /a/b/c: The system cannot find the path specified.
	if err = os.Mkdir("/a/b/c", os.ModeDir); err != nil {
		fmt.Printf("创建指定文件夹失败：%s\n", err)
	}

	// 如果目标文件夹已经存在，则会返回 nil
	// 前置路径对应的文件夹不存在，会创建，创建成功 - 创建到了 E 盘下，等同于参数是 E:/a/b/c
	if err = os.MkdirAll("/a/b/c", os.ModeDir); err != nil {
		fmt.Printf("创建层级文件夹失败：%s\n", err)
	}
}

func createFileTest() {
	var (
		file *os.File
		err  error
	)

	// 要求父级目录已经存在
	if file, err = os.Create("/a/b/c/text.txt"); err != nil {
		fmt.Printf("创建文件失败：%s\n", err)
	}
	_ = file.Close()
}
