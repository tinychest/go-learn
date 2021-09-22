package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

/*
	参考煎鱼的文章：https://mp.weixin.qq.com/s/QOA3Mk20M4rRM9oXOGB9HA

	只要在程序代码中有相对路径的概念，有相对，就会存在相对于谁，这个谁称之为基础路径
	在 Go 中，基础路径就是执行程序所在的路径

	这就意味着，当程序中存在任何相对路径，那么程序运行就会受到执行执行时所在路径的影响
	举例（./conf/app.conf）来说，你在什么目录下执行程序，就需要在在目录下创建对应的（./conf/app.conf）

	但是反过来说，就是 test，假如项目中有需求，单独配置一个测试环境的配置，那么就完全可以利用上相对路径这个特性，只要在 test 目录下创建对应目录结构的配置文件就可以了
*/

// 《go build》
// cd <cur>
// go build exec.go → exec.exe
// exec.exe → E:\Learning-Workspace\go\src\go-learn\core\basic\file\path
// cd .. → path\exec.exe | .\path\exec.exe → E:\Learning-Workspace\go\src\go-learn\core\basic\file\path

// 《go run》
// cd <cur>
// go run exec.go → C:\Users\xxx\AppData\Local\Temp\go-build3649388729\b001\exe
// cd .. → go run path\exec.go → C:\Users\xxx\AppData\Local\Temp\go-build3649388729\b001\exe

// 结论：main 函数返回的结果可以作为绝对路径来使用，配置文件只要相对于这个绝对路径来设置（也就是配置绝对路径），这样无论执行命令所在的路径是什么都能得到一样的结果

// 但是，面对 go run 依旧无法解决（也好在实际还是 go build 的应用场景广泛）
// 解决1：参照 beego 读取配置文件的方式（不能完全解决）
// 解决2：就不要写相对路径，写绝对路径
// 解决3：太过硬码的绝对不好，所以可以通过系统环境变量来指定配置文件的位置
func main() {
	p, _ := exec.LookPath(os.Args[0])
	f, _ := filepath.Abs(p)
	fmt.Println(f[:strings.LastIndex(f, string(os.PathSeparator))])
}
