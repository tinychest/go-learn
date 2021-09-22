package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// 参考
// https://blog.csdn.net/whatday/article/details/118657417

// 引出
// Go 语言实现了自己的运行时，因此，对信号的默认处理方式和普通的 C 程序不太一样
// 信号量有默认的处理方式，有些信号，程序是可以改变默认行为的，这也就是 os/signal 包的用途（Windows 下只支持 os.SIGINT 信号）

// 信号量
// 源码中：go/src/syscall/types_windows.go:67
// Linux 中：可以通过 kill -l 命令查看（有 64 个）
// 1 - SIGHUP - hangup
// 2 - SIGINT- interrupt（Ctrl + C、kill <pid>）
// 3 - SIGQUIT - quit
// 9 - SIGKILL - killed（很特殊，如果处理信号的进程被强杀了，也就不存在处理一说了）
// 15 - SIGTERM - terminated（kill -9 <pid>）

// Notify
// 简单来说，Notify 可以理解为订阅，作为 Notify 参数的通道，当进程收到各种信号时，Go signal 只会发送订阅的信号
// 不加参数时，默认接收所有的信号
// 通道为什么要设置缓存，详见 Notify 源码注释

// 注意，如果存在通道订阅了 SIGINT 信号，那么程序默认的 Ctrl + C 行为将不再生效
// （不要有 Ctrl + C 等于结束程序先入为主的观念）

// Reset
// 和 Notify 正好成反义词

// Stop
// 那么假如你希望某个通道在整个生命周期里就只接收一次信号，并且希望这个通道一接收到信号程序就无条件结束，那么可以使用 Stop 方法
// 举例来说，当有通道订阅了 SIGINT 信号，会导致 Go 将该信号传递给该通道，不再有默认的结束程序行为，即程序不会结束
// （没有进一步调研，感觉没有任何应用场景）

// Ignored
// 没有测试出效果，无论在不在 Notify 之前，实际发送信号，程序里的通道还是可以收到信号

// 简单的测试结果如下（在 Windows 上不好测试，Linux 可以通过 kill 命令向指定的 pid 发送信号）
// Windows：Ctrl + C → interrupt
// Linux：Ctrl + C | kill -2 → interrupt
// kill -0：代码无法处理 - 通道不会收到信号
// kill -1：代码无法处理 - 挂起，会打印出 hangup
// kill -9：代码无法处理 - 终止程序，会打印出 killed
// kill -15：信号量不支持 16 进制
func main() {
	// 打印进程 pid
	fmt.Println(os.Getpid())

	// 监听信号
	c := make(chan os.Signal, 1)
	signal.Notify(c)
	for s := range c {
		fmt.Println("[End]", s)
	}
	time.Sleep(10 * time.Second)
}

func flagTest() {
	var appPath string
	flag.StringVar(&appPath, "app-path", "123", "项目路径")
	flag.Parse()
	fmt.Println(appPath)

	// getRunParam()
	getAppPath()

	// 获取完成的启动命令
	fmt.Println(os.Args)
}

func getAppPath() {
	f, _ := exec.LookPath(os.Args[0])
	p, _ := filepath.Abs(f)
	path := p[:strings.LastIndex(p, string(os.PathSeparator))]
	fmt.Println(path)
}
