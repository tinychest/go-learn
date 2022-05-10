package main

import (
	"flag"
	"fmt"
	"os"
)

// 获取命令行参数 os.Args
//
// 这里的 flag 包是 Go 标准库中的，和启动命令参数相关的包
//
// - 命令行指定参数的方式
// [ ] name value
// [✔] -name value
// [✔] --name value
// [ ] name=value
// [✔] -name=value
// [✔] --name=value
// （布尔类型的参数必须使用等号的方式指定）
//
// 调用 Parse 方法，才会实际去将命令行的参数反序列化到变量中

func Usage() {
	w := flag.CommandLine.Output()

	_, err := fmt.Fprintf(w, "Usage of %s:\n", os.Args[0])
	_, err = fmt.Fprintf(w, "\tUsed Like: xxx -name <name> -pass <pass>\n")
	_, err = fmt.Fprintf(w, "For more information, see:\n")
	_, err = fmt.Fprintf(w, "\thttps://xxx.xxx/xxx/xxx\n")
	_, err = fmt.Fprintf(w, "Flags:\n")
	if err != nil {
		panic(err)
	}

	flag.PrintDefaults()
}

func main() {
	var name string
	var pass string

	// 命令、参数帮助
	flag.Usage = Usage

	flag.StringVar(&name, "name", "", "用户名")
	flag.StringVar(&pass, "pass", "", "密码")

	flag.Parse()

	fmt.Println(os.Args)
	fmt.Println(name)
	fmt.Println(pass)

	// 定义启动参数（方式一）
	// flag.Var()
	// flag.StringVar()
	// flag.IntVar()
	// flag.Int64Var()
	// flag.Float64Var()
	// flag.DurationVar()
	// flag.BoolVar()

	// 定义启动参数（方式二）
	// flag.String()
	// flag.Int()
	// flag.Int64()
	// flag.Float64()
	// flag.Duration()
	// flag.Bool()

	// 其他函数
	// flag.Args()  // 参数列表
	// flag.NFlag() // 参数个数
	// flag.NArg()  // 命令行参数后的其他参数个数
}
