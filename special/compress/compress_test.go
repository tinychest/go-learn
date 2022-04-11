package compress

import (
	"archive/zip"
	"io"
	"os"
	"testing"
)

func TestCompress(t *testing.T) {
	f1, err := os.Open("./1.txt")
	if err != nil {
		t.Fatal(err)
	}
	f2, err := os.Open("./2.txt")
	if err != nil {
		t.Fatal(err)
	}
	f3, err := os.Open("./3.txt")
	if err != nil {
		t.Fatal(err)
	}

	err = compress([]*os.File{f1, f2, f3}, "./test.zip")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
}

func compress(srcs []*os.File, dst string) error {
	// 防止无法覆盖
	if err := os.RemoveAll(dst); err != nil {
		return err
	}

	// 创建 zip 文件
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	archive := zip.NewWriter(dstFile)
	defer archive.Close()

	for _, f := range srcs {
		// 获取源文件信息
		info, err := f.Stat()
		if err != nil {
			return err
		}
		// 如果文件是文件夹就跳过
		if info.IsDir() {
			continue
		}
		// 获取 zip 文件信息、设置压缩算法
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Method = zip.Deflate

		// 将文件信息添加至目标压缩文件
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		// 开始写入
		_, err = io.Copy(writer, f)
		if err != nil {
			return err
		}
	}

	return nil
}