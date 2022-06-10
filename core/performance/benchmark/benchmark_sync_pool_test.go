package benchmark

import (
	"encoding/json"
	"sync"
	"testing"
)

/*
go test -bench="Unmarshal" . -benchmem

goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark/pool
Benchmark_Unmarshal-8                9969            127896 ns/op            1376 B/op          7 allocs/op
Benchmark_UnmarshalWithPool-8        9999            125913 ns/op             224 B/op          6 allocs/op
PASS
ok      go-learn/unit_test/benchmark/pool       2.771s

从测试结果上说，性能差距并不大的原因：
1.Student 结构体实例占用内存小
2.标准库 json 反序列化时利用了反射，效率比较低，这是占据时间的大头

消耗的内存完全是不同的，也很好理解，从 pool 里获取的方式，就是一直在从一个结构体实例中读取数据，和用一个创建一个，当然不同
主要影响的性能是 gc

TODO 其实这里给的例子并没有什么说服力，因为大多数场景都是要将某一结构的结构体多份数据形成一个列表，从而调用后面的方法
    还有就是说，单独引出一个池子的概念，是为了防止强引用导致的资源无法被回收
Go 语言标准库中大量使用了 sync.Pool，例如 fmt 和 encoding/json
*/

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "GeekFW", Age: 24})

var studentPool = sync.Pool{
	New: func() interface{} {
		return new(Student)
	},
}

func Benchmark_Unmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}
}

func Benchmark_UnmarshalWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := studentPool.Get().(*Student)
		_ = json.Unmarshal(buf, stu)
		studentPool.Put(stu)
	}
}
