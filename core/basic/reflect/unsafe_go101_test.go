package _reflect

import (
	"reflect"
	"strings"
	"testing"
	"unsafe"
)

/*
【参见】
https://gfw.go101.org/article/unsafe.html
（如果想了解 unsafe 包，请直接去看原文，这里的记录仅仅为了促进理解）

【概念】
[源码注释]
unsafe contains operations that step around the type safety of Go programs.
直译过来就是可以绕过 Go 的安全检查，unsafe 命名很恰当。

[为什么使用 unsafe 包？]
直接引用原文的话：在某些情形下，通过非类型安全指针的帮助，我们可以写出效率更高的代码；
但另一方面，使用非类型安全指针也导致我们可能轻易地写出潜在的不安全的代码，这些潜在的不安全点很难在它们产生危害之前被及时发现。

[什么时候使用 unsafe 包？]
详见 unsafe 包的注释，给出了 6 个例子，详见下面的 TestPractice

[参照理解]
Go 的 星类型（*），也就是指针类型，可以理解为类型安全指针，它可以帮助我们写出安全的代码，但是某些性能要求极致的场景下，这限制是我们不能写出高效的代码（下面会给出例子）。
unsafe 包提供的 Pointer（unsafe.Pointer），即 Go 提供的非类型安全的指针，就和 C 指针类似，他们很强大，但同时也都很危险。

[Pointer]
type Pointer *ArbitraryType
ArbitraryType 代表 随意、任意 类型，即 可以被转换为任意类型安全指针（反之亦然）
类型零值为 nil

[地址值]
Go 中使用 uintptr 类型来表示一个指针指向的地址值，本质上就是一个整数值

【常见 api】
-- 下面 3 个方法 Go 1.17 就有，且这 3 个方法是安全的 --
unsafe.Alignof 详见 /performance/memory_align_test.go
unsafe.Offsetof 用于获取一个结构体值的某个字段的地址相对于此结构体值的地址的偏移
unsafe.Sizeof 用于获取一个值的尺寸（亦即此值的类型的尺寸）
-- Go 1.17 之后 --
unsafe.Add 在一个指针（非类型安全）基础之上，添加一个偏移量，然后返回表示新地址的一个指针
unsafe.Slice 从一个任意指针（类型安全）派生出一个指针长度的切片

【类型转换规则】
一个类型安全指针值可以被显式转换为一个非类型安全指针类型，反之亦然。
一个 uintptr 值可以被显式转换为一个非类型安全指针类型，反之亦然。但是，注意，一个 nil 非类型安全指针类型不应该被转换为 uintptr 并进行算术运算后再转换回来。

【点】
一个 uintptr 值并不引用任何值，它被看作一个整数，即，GC 工作时，可能会回收 uintptr 引用的内存空间

[回收时间点是不确定的] 所以，如果继续通过解引用 uintptr 的地址值进行内存操作，是非常危险的，因为，内存可能已经分配给其他地方使用了
[一个值的地址在程序运行中可能改变] 一个 Goroutine 的栈的大小发生改变时，开辟在次栈上的内存块需要移动，从而相应的地址将改变
[一个值的生命范围可能没有那么大] 一个结构体实例引用的指针空间，如果不再被用到，就可能会被回收
[*unsafe.Pointer 是一个类型安全的指针] 额...
*/

func TestPractice(t *testing.T) {
	/* 一、将类型 *T1 的一个值转换为非类型安全指针值，然后将此非类型安全指针值转换为类型 *T2。 */
	// 答：只有在T1的尺寸不小于T2并且此转换具有实际意义的时候才应该实施这样的转换

	// 例1：math.Float64bits、math.Float64frombits（注意，和直接转换的结果不同）
	var f float64 = 1.11
	var _ int64 = int64(f)
	var _ uint64 = uint64(f)

	// 例2：避免额外的空间开辟和复制（通过安全的方式是无法实现的）
	type MyString string
	var ms []MyString
	var s []string

	// s = ms // 编译不通过
	// ms = s // 编译不通过

	ms = *(*[]MyString)(unsafe.Pointer(&s))
	s = *(*[]string)(unsafe.Pointer(&ms))

	// unsafe.Slice（Go 1.17）
	ms = unsafe.Slice(&ms[0], len(ms))            // 转换寂寞
	s = unsafe.Slice(&s[0], len(s))               // 转换寂寞
	ms = unsafe.Slice((*MyString)(&s[0]), len(s)) // Goland 抽风，这个不能省，却提示可以省
	s = unsafe.Slice((*string)(&ms[0]), len(ms))  // Goland 抽风，这个不能省，却提示可以省

	// 例2 拓展：不再使用的字节切片转换为一个字符串（从而避免对底层字节序列的一次开辟和复制）
	// strings.Builder{}.String()
	// 字节切片的尺寸比字符串的尺寸要大，并且它们的底层结构类似，所以此转换（对于当前的主流Go编译器来说）是安全的
	// 反之，因为 string 尺寸比 []string 小，所以 string → []byte 是危险的
	// PS：修改 unsafe 转化的 string 字节，是会 fatal 的
	// 虽然安全，但是不推荐在用户代码中使用，应该使用文案结尾 模式六 样例

	/* 二、将一个非类型安全指针值转换为一个 uintptr 值，然后使用此 uintptr 值。（作用不大） */
	type K struct{ a int }
	var k K
	t.Log(&t)                                   // 0xc6233120a8
	t.Logf("%p\n", &k)                          // 0xc6233120a8
	t.Logf("%x\n", uintptr(unsafe.Pointer(&k))) // c6233120a8

	/* 三、将一个非类型安全指针转换为一个 uintptr 值，然后此 uintptr 值参与各种算术运算，再将算术运算的结果 uintptr 值转回非类型安全指针。 */
	// 下面的操作更推荐使用 unsafe.Add
	// 还有其他较为深刻的说明，详见 Go 101 原文
	type M struct {
		x bool
		y [3]int16
	}
	const s1 = unsafe.Offsetof(M{}.y)
	const s2 = unsafe.Sizeof(M{}.y[0])
	m := M{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&m)
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + s1 + s2 + s2)) // m.y[2]的内存地址
	t.Log(*ty2)                                                // 789

	/* 四、将非类型安全指针值转换为 uintptr 值并传递给 syscall.Syscall 函数调用 */

	/* 五、将 reflect.Value.Pointer 或者 reflect.Value.UnsafeAddr 方法的 uintptr 返回值立即转换为非类型安全指针 */
	// （自己理解：这几个点都是要告诉你，避免使用 unsafe 时，GC 将该处内存空间回收了）

	// reflect.Value.Pointer 和 reflect.Value.UnsafeAddr 的方法返回值都是 uintptr，如此设计的意图
	// 引用原文：
	// 这样设计的目的是避免用户不引用unsafe标准库包就可以将这两个方法的返回值（如果是unsafe.Pointer类型）转换为任何类型安全指针类型
	// 这样的设计需要我们将这两个方法的调用的uintptr结果立即转换为非类型安全指针。
	// 否则，将出现一个短暂的可能导致处于返回的地址处的内存块被回收掉的时间窗。
	// 此时间窗是如此短暂以至于此内存块被回收掉的几率非常之低，因而这样的编程错误造成的bug的重现几率亦十分得低

	// 样例（安全） p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
	// 样例（危险）
	// u := reflect.ValueOf(new(int)).Pointer()
	// // 在这个时刻，处于存储在u中的地址处的内存块
	// // 可能会被回收掉。
	// p := (*int)(unsafe.Pointer(u))

	/* 六、将一个 reflect.SliceHeader 或者 reflect.StringHeader 值的 Data 字段转换为非类型安全指针，以及其逆转换 */
	// 原文：
	// reflect 标准库包中的 SliceHeader 和 StringHeader 类型的 Data 字段的类型被指定为 uintptr，而不是 unsafe.Pointer

	// TODO 我们可以将一个字符串的指针值转换为一个 *reflect.StringHeader 指针值，从而可以对此字符串的内部进行修改。
	//  类似地，我们可以将一个切片的指针值转换为一个*reflect.SliceHeader指针值，从而可以对此切片的内部进行修改。

	// 一般说来，我们只应该从一个已经存在的字符串值得到一个*reflect.StringHeader指针，
	// 或者从一个已经存在的切片值得到一个*reflect.SliceHeader指针，而不应该从一个StringHeader值生成一个字符串，或者从一个SliceHeader值生成一个切片

	// 样例：给出使用非类型安全途径将一个字符串转换为字节切片的方法（使用非类型安全途径避免了复制一份底层字节序列）
	// 下面两个函数实例的前提都是：一定不要修改字节数组的值
	String2ByteSlice := func(str string) (bs []byte) {
		strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
		sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
		sliceHdr.Data = strHdr.Data
		sliceHdr.Cap = strHdr.Len
		sliceHdr.Len = strHdr.Len
		return
	}
	ByteSlice2String := func(bs []byte) (str string) {
		sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
		strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
		strHdr.Data = sliceHdr.Data
		strHdr.Len = sliceHdr.Len
		return
	}

	_ = ByteSlice2String

	// 否则。为什么说否则，因为这个表现已经和 Go 语言定义标准大相径庭了（字符串开辟在不可修改内存区）
	str := strings.Join([]string{"Go", "land"}, "")
	ss := String2ByteSlice(str)
	t.Logf("%s\n", ss) // Goland
	ss[5] = 'g'
	t.Log(str) // Golang

	/* 作者给出了几个非常好的参考例子 */

	/* 总结 */
	// ONE、没有一定需要极致性能的场景下，没有能独当一面的 Go 工程师，千万不要使用 unsafe，结合实际项目冗杂的业务逻辑，出的问题是极难发现的，并且几乎无法复现和调试
	// （解决问题原则：发现问题永远是最难、最花时间的，因此我们一直在各种思考，前进的路上，以此避免各种各样的错误）
	// TWO、不得不使用 unsafe 的时候，一定要遵守上面的模式和原则
	// THREE、当前的非类型安全机制规则和使用模式可能在以后的Go版本中完全失效（作者的 Go 环境是 Go 1.17）
	// FOUR、Go 官方工具链 1.14 中加入了一个 -gcflags=all=-d=checkptr 编译器动态分析选项（在Windows平台上推荐使用工具链1.15+）。
	// 当此选项被使用的时候，编译出的程序在运行时会监测到很多（但并非所有）非类型安全指针的错误使用。一旦错误的使用被监测到，恐慌将产生
}
