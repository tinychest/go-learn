package pprof

/*
【目的】
gin web 服务的性能分析

【参考】
https://mp.weixin.qq.com/s/m68JmVxEW2NebjM7PMnE1A

生产一定不能这样搞，作者也说是测试模拟环境下

【具体操作】
[项目代码拓展]
import "github.com/gin-contrib/pprof"
pprof.Register(r)

[访问页面]
http://localhost:<项目端口>/debug/pprof/
可以从出来的页面查看各种信息，这里只说 profile，点击时会卡住，因为是默认采样 30s，30s 后给出采样结果文件

[分析]
- 使用 go tool 进行内存分析
go tool pprof http://localhost:<项目端口>/debug/pprof/heap
- 使用 go tool 进行 CPU 分析
go tool pprof http://localhost:<项目端口>/debug/pprof/profile

[但是更希望看到火焰图]
- go install github.com/uber-archive/go-torch@latest

参照 go-torch 的 README
- cd $GOPATH/src/github.com/uber/go-torch（实际没有看到 src 下的这个，所以到 bin 了，后面发现只要在环境变量中就行，以供 to-torch 调用）
- cd $GOPATH/bin
- git clone https://github.com/brendangregg/FlameGraph

[开始生成]
go-torch -u http://localhost:<项目端口>/ -t 5

GG：FATAL[15:20:43] Failed: could not generate flame graph: fork/exec ./FlameGraph/flamegraph.pl: %1 is not a valid Win32 application.
（这是一个其他系统上的脚本文件）
*/
