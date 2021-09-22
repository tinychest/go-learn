package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"testing"
)

func TestPath(t *testing.T) {
	println(srcDir())
	println(srcDir2())
	println(srcDir3())
	println(srcDir4())

	println(execFile())
	println(execFile2())
}

func srcDir() string {
	f, _ := os.Getwd()
	return f
}

func srcDir2() string {
	f, _ := syscall.FullPath(".")
	return f
}

func srcDir3() string {
	f, _ := syscall.Getwd()
	return f
}

func srcDir4() string {
	// 参数带 拓展名，返回的结果才会带 拓展名
	f, _ := filepath.Abs(".")
	return f
}

func execFile() string {
	p, _ := exec.LookPath(os.Args[0])
	f, _ := filepath.Abs(p)
	return f
}

func execFile2() string {
	f, _ := os.Executable()
	return f

	// 拓展：会得到 .（细品）
	// println(path.Dir(f))
}
