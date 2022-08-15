package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	// 获取当前项目的执行路径
	fmt.Println(getAppPath())
}

// 在 goland console 中打印出来的相关 exe 路径会被标注为红色
func getAppPath() string {
	path := os.Args[0]
	println(path)

	path, err := exec.LookPath(path)
	println(path)
	if err != nil {
		panic(err)
	}

	path, err = filepath.Abs(path)
	println(path)
	if err != nil {
		panic(err)
	}

	return path[:strings.LastIndex(path, string(os.PathSeparator))]
}
