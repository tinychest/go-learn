package compress

import (
	"archive/zip"
	"io"
	"os"
	"strings"
	"testing"
)

type F struct {
	Name    string
	Content string
}

func TestMemCompress(t *testing.T) {
	fs := []F{
		{
			Name:    "1.txt",
			Content: "1111111",
		},
		{
			Name:    "2.txt",
			Content: "2222222",
		},
		{
			Name:    "3.txt",
			Content: "333333",
		},
	}

	// 假如是在 web 环境中，只要将 dstFile 改成 ResponseWriter 就行了
	dst := "./mem.zip"
	dstFile, err := os.Create(dst)
	if err != nil {
		t.Fatal(err)
	}
	defer dstFile.Close()

	archive := zip.NewWriter(dstFile)
	defer archive.Close()

	for _, f := range fs {
		// 开始写入文件
		w, err := archive.CreateHeader(&zip.FileHeader{
			Name:   f.Name,
			Method: zip.Deflate,
		})
		if err != nil {
			t.Fatal(err)
		}

		_, err = io.Copy(w, strings.NewReader(f.Content))
		if err != nil {
			t.Fatal(err)
		}
	}
	t.Log("success")
}
