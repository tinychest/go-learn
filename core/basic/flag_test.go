package basic

import (
	"flag"
	"os"
	"testing"
)

// 主旨：获取命令行参数
// 核心：os.Args

/*
在 main.go 中（下面的代码拿得到）
go run main.go --app-path "xxx"
go run main.go --app-path="xxx"

在 test 中（下面的代码拿不到）
go test -test.run TestFlag$ -v -args app-path="123"
go test -test.run=TestFlag$ -v -args app-path="123"
*/
func TestFlag(t *testing.T) {
	const defaultValue = "./go-learn"

	var appPath string
	flag.StringVar(&appPath, "app-path", defaultValue, "项目路径")
	flag.Parse()
	t.Log(appPath)
}

// 只能解析串了
// go test -test.run TestFlag2 -v -args app-path="123"（-args 去掉语法都不正确了）
func TestFlag2(t *testing.T) {
	t.Log(os.Args)
}