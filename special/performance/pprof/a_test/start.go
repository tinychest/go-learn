package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
)

// 会在项目根目录下，创建 cpu.pprof
// 在项目根目录下，通过 go tool 分析：go tool pprof cpu.pprof
// 查看性能消耗排名：top10
// 会发现 runtime.mallocgc 占了一个不小的份量

// 整个测试衍生流程的建议可以通过 go build -gcflags=-m xxx.go 得到
// 注意 go 文件不能是 test 文件，不然会提示 no packages to build
func main() {
	cpufile, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	err = pprof.StartCPUProfile(cpufile)
	if err != nil {
		panic(err)
	}
	defer cpufile.Close()
	defer pprof.StopCPUProfile()

	// ensure function output is accurate
	if foo(12345) == "aajmtxaattdzsxnukawxwhmfotnm" {
		fmt.Println("Test PASS")
	} else {
		fmt.Println("Test FAIL")
	}

	for i := 0; i < 100; i++ {
		foo(rand.Int())
	}
}

func foo(n int) string {
	var buf bytes.Buffer
	for i := 0; i < 100000; i++ {
		buf.WriteString(strconv.Itoa(n))
	}
	sum := sha256.Sum256(buf.Bytes())

	var b []byte
	for i := 0; i < int(sum[0]); i++ {
		x := sum[(i*7+1)%len(sum)] ^ sum[(i*5+3)%len(sum)]
		c := strings.Repeat("abcdefghijklmnopqrstuvwxyz", 10)[x]
		b = append(b, c)
	}
	return string(b)
}