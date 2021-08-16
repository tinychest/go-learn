package file

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

func TestPath(t *testing.T) {
	println(projectPath())
	println(projectPath2())
	println(projectPath3())
	println(projectPath4())

	println(execPath())
	println(execPath2())
}

func projectPath() string {
	wd, _ := os.Getwd()
	return wd
}

func projectPath2() string {
	wd, _ := syscall.FullPath(".")
	return wd
}

func projectPath3() string {
	wd, _ := syscall.Getwd()
	return wd
}

func projectPath4() string {
	// 参数带 拓展名，返回的结果才会带 拓展名
	wd, _ := filepath.Abs(".")
	return wd
}

func execPath() string {
	f, _ := exec.LookPath(os.Args[0])
	p, _ := filepath.Abs(f)
	return p

	// 拓展：获取目录
	// p[:strings.LastIndex(p, string(os.PathSeparator))]
}

func execPath2() string {
	executable, _ := os.Executable()
	// 当前源码文件 生成的可执行文件 的绝对路径名
	return executable

	// 拓展 会得到 .（你细品）
	// println(path.Dir(executable))
}
