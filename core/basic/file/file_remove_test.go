package file

import (
	"os"
	"testing"
)

// 如果文件没有 Close 是无法 Remove 的
func TestFileRemove(t *testing.T) {
	var err error

	// 删除文件夹（不为空不让删除，返回 error） 或 文件
	if err = os.Remove("C:/Users/14590/Desktop/1"); err != nil {
		t.Logf("删除文件失败：%s\n", err)
	}

	// 删除文件夹（如果内部有东西，也全部一起删掉） 或 文件
	if err = os.RemoveAll("C:/Users/14590/Desktop/1"); err != nil {
		t.Logf("删除文件失败：%s\n", err)
	}
}
