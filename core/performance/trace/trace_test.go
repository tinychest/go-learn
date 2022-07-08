package trace

import (
	"fmt"
	"os"
	"runtime/trace"
	"testing"
)

// 使用 go tool trace test <TraceFileName> → 开启一个监听特定端口的服务，访问可以查看详细资源信息的页面
// 只监听 127.0.0.1 的 ip，并不能够通过局域网的 ip 访问

const TraceFileName = "trace.out"

func TestTrace(t *testing.T) {
	f, err := os.Create(TraceFileName)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	if err = trace.Start(f); err != nil {
		t.Fatal(err)
	}

	fmt.Println("Hello Trace")

	trace.Stop()
}
