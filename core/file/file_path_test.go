package file

import (
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestGetPath(t *testing.T) {
	wd, _ := os.Getwd()
	executable, _ := os.Executable()
	hostname, _ := os.Hostname()

	dir := path.Dir(executable)
	abs, _ := filepath.Abs(".")

	// 当前源码文件 所在目录 的绝对路径
	println(wd)
	// 当前源码文件 生成的可执行文件 的绝对路径名
	println(executable)
	// DESKTOP-N3FSCS2
	println(hostname)
	// 会得到 .（你细品）
	println(dir)
	// 获取相对路径对应的绝对路径（参数带 拓展名，返回的结果才会带 拓展名）
	print(abs)
}
