package file

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	fileExists("E:/Learning-Workspace/hexo/blog/db.json")
	fileExists2("E:/Learning-Workspace/hexo/blog/db.json")
}

// beego 工具类中的源码
func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 文件是否存在（不分文件和文件夹）
func fileExists2(path string) bool {
	dirInfo, err := os.Lstat(path)
	if os.IsNotExist(err) {
		fmt.Printf("文件不存在：%s\n", err)
		return false
	}
	if err != nil {
		fmt.Printf("读取文件报错：%s", err)
		return false
	}

	file, _ := os.Open(path)

	// println(filepath.Abs(dirInfo.Name()), filepath.Dir(dirInfo.Name()))
	fmt.Printf("文件存在：%s，%s\n", file.Name(), dirInfo.Mode())
	return true
}
