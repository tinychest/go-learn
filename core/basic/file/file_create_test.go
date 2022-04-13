package file

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	// createDirTest(t)
	createFileTest(t)
}

func createDirTest(t *testing.T) {
	var err error

	// 创建参数中的最后一层文件夹
	// 前置路径对应的文件夹不存在，则创建失败 - mkdir /a/b/c: The system cannot find the path specified.
	if err = os.Mkdir("/a/b/c", os.ModeDir); err != nil {
		t.Logf("创建指定文件夹失败：%s\n", err)
	}

	// 如果目标文件夹已经存在，则会返回 nil
	// 前置路径对应的文件夹不存在，会创建，创建成功 - 创建到了 E 盘下，等同于参数是 E:/a/b/c
	if err = os.MkdirAll("/a/b/c", os.ModeDir); err != nil {
		t.Logf("创建层级文件夹失败：%s\n", err)
	}
}

func createFileTest(t *testing.T) {
	fp := "/a/b/c/text.txt"

	fp, err := filepath.Abs(fp)
	if err != nil {
		t.Fatalf("获取文件绝对路径失败: %s\n", err)
	}

	// 创建目录
	idx := strings.LastIndex(fp, "/")
	if idx == -1 {
		idx = strings.LastIndex(fp, "\\")
	}
	dir := fp[:idx]

	if err = os.MkdirAll(dir, os.ModeDir); err != nil {
		t.Fatalf("创建目录失败: %s\n", err)
	}

	// 创建文件（要求父级目录已经存在）
	f, err := os.Create(fp)
	if err != nil {
		t.Fatalf("创建文件失败: %s\n", err)
	}
	_ = f.Close()
}
