package stringer

// 一、直接执行 stringer -type=Pill 就能得到枚举描述方法的源码
// （如果提示 stringer 命令找不到，就执行 go get -u golang.org/x/tools/cmd/stringer）
// （会在当前目录下生成一个名为 pill_string.go 的源码文件，里边为 Pill 实现了 string 方法）

// 注意，当前文件名不能以 test 结尾，否则 stringer: no values defined for type Pill

// 二、放开下边的注释，直接在当前目录所在路径下，执行 go generate，也可以实现同样的效果
// //go:generate stringer -type=Pill
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
