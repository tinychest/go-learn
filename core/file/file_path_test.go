package file

import (
	"os"
	"path"
	"path/filepath"
	"syscall"
	"testing"
)

func TestGetPath(t *testing.T) {
	wd1, _ := os.Getwd()
	wd2, _ := syscall.FullPath(".")
	wd3, _ := syscall.Getwd()
	wd4, _ := filepath.Abs(".")

	executable, _ := os.Executable()
	hostname, _ := os.Hostname()

	dir := path.Dir(executable)

	// 当前源码文件 所在目录 的绝对路径
	println(wd1)
	println(wd2)
	println(wd3)
	println(wd4) // 参数带 拓展名，返回的结果才会带 拓展名

	// 当前源码文件 生成的可执行文件 的绝对路径名
	println(executable)
	// DESKTOP-N3FSCS2
	println(hostname)
	// 会得到 .（你细品）
	println(dir)
}
