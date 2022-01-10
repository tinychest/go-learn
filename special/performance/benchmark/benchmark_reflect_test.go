package benchmark

import (
	"reflect"
	"testing"
)

/*
反射的主要作用：简化代码
反射的常用应用场景：将配置文件的数据映射到结构体中、ORM 中的数据表数据到结构体中、没有泛型想做到通用类型的方法是离不开反射的等等
*/

type Config struct {
	Name    string `json:"server_name"`
	IP      string `json:"server_ip"`
	URL     string `json:"server_url"`
	Timeout string `json:"timeout"`
}

/*
一、测试反射创建实例的性能

go test -bench="New$" .

goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
Benchmark_DirectNew-8            23528396                54.2 ns/op
Benchmark_ReflectNew-8           11646188                94.4 ns/op
PASS
ok      go-learn/unit_test/benchmark    2.581s

在创建实例上，存在不能忽视的性能差距；上面是 Go 1.16 的基准测试结果，Go 1.17 中，通过反射创建的性能得到了提升
*/
func Benchmark_DirectNew(b *testing.B) {
	var config *Config
	for i := 0; i < b.N; i++ {
		config = new(Config)
	}
	_ = config
}

func Benchmark_ReflectNew(b *testing.B) {
	var config *Config
	typ := reflect.TypeOf(Config{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config, _ = reflect.New(typ).Interface().(*Config)
	}
	_ = config
}

/*
二、测试反射修改结构体实例字段值的性能

go test -bench="Set$" .

goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
Benchmark_Set-8                          1000000000            0.626 ns/op
BenchmarkReflect_FieldSet-8             39997999              30.2 ns/op
BenchmarkReflect_FieldByNameSet-8       3986682               310 ns/op
PASS
ok      go-learn/unit_test/benchmark    3.702s

性能差距明显：
对于一个普通的拥有 4 个字段的结构体 Config 来说，使用反射给每个字段赋值，相比直接赋值，性能劣化约 100 - 1000 倍。其中，FieldByName 的性能相比 Field 劣化 10 倍
通过名称赋值和下标赋值的性能差距很好理解，因为根据名称是有一个查找过程的，时间复杂度为 O(n)
*/
func Benchmark_Set(b *testing.B) {
	config := new(Config)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		config.Name = "name"
		config.IP = "ip"
		config.URL = "url"
		config.Timeout = "timeout"
	}
}

func BenchmarkReflect_FieldSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(0).SetString("name")
		ins.Field(1).SetString("ip")
		ins.Field(2).SetString("url")
		ins.Field(3).SetString("timeout")
	}
}

func BenchmarkReflect_FieldByNameSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.FieldByName("Name").SetString("name")
		ins.FieldByName("IP").SetString("ip")
		ins.FieldByName("URL").SetString("url")
		ins.FieldByName("Timeout").SetString("timeout")
	}
}

/*
三、提高反射性能
1、选对工具类库
例如在 RPC 协议中，需要对结果提进行序列化和反序列化，这个时候要避免使用 Go 标准库 json 包下的 Marshal 和 Unmarshal
因为它们都是通过反射实现的，推荐 easyjson，性能在大部分场景下有 5 倍的提升

2、使用缓存
拿上例中的 BenchmarkReflect_FieldByNameSet 方法来说，直接上代码
*/
func BenchmarkReflect_FieldByNameCacheSet(b *testing.B) {
	typ := reflect.TypeOf(Config{})
	cache := make(map[string]int)
	for i := 0; i < typ.NumField(); i++ {
		cache[typ.Field(i).Name] = i
	}
	ins := reflect.New(typ).Elem()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ins.Field(cache["Name"]).SetString("name")
		ins.Field(cache["IP"]).SetString("ip")
		ins.Field(cache["URL"]).SetString("url")
		ins.Field(cache["Timeout"]).SetString("timeout")
	}
}

/*
go test -bench="Set$" .

goos: windows
goarch: amd64
pkg: go-learn/unit_test/benchmark
Benchmark_Set-8                                  1000000000               0.640 ns/op
BenchmarkReflect_FieldSet-8                     35322998                30.6 ns/op
BenchmarkReflect_FieldByNameSet-8                3896040               328 ns/op
BenchmarkReflect_FieldByNameCacheSet-8          17644048                69.7 ns/op
PASS
ok      go-learn/unit_test/benchmark    4.932s

BenchmarkReflect_FieldByNameCacheSet 和 BenchmarkReflect_FieldSet 的差距没那么大了，但是和直接设值还是没法比
*/
