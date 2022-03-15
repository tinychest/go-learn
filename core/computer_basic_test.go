package core

import "testing"

// 计算机组成原理基础：为什么要有补码，计算机的硬件结构就只有加法器，所以遵从 《模》 的概念，将所有的减法转化成加法

// 回忆：负数补码的计算方法
// 负数对应正数的原码 → 取反 → +1 → 补码
// 负数原码 → 反码 → +1 → 补码

// 负数 和 正数 相加的结果 = 负数的补码 + 正数（原码、反码、补码相同）
// 即 结果（补） = 数1（补） + 数2（补）
// 也就是你看上去，你觉得的加减法，在计算机世界中只有：补码的加法

// -128 的补码为什么是 1000 0000？
// 一、数学角度
// 256 - 128 = 128
// → 256 + (-128的补码) = 128
// → (-128的补码) = 256 - 128 = 128 = 1000 0000

// 二、规定角度（既然规定了是有符号数的环境，那么最高位就是符号位）
// 0111 1111 → 127
// 0000 0000 → 0
// 1111 1111 → -127
// 1000 0000？无脑 → 128（上下文是有符号数喔） → -0 = 0（已经有了） → -128
func TestComputerBasic(t *testing.T) {
	// int8 数据类型范围 [-128, 127]

	var a int8 = -1
	var b int8 = -128 / a
	t.Log(b)

	// 127  0111 1111（道理）
	// -128 / -1 的结果很简单就是 128（理所当然），在内存的结果就是 0(符号位) 1000 0000，但是现在要转 int8 存不下额外的符号位信息

	// 说的再清楚一些，计算机中的计算、存储都是补码
	// 就是计算结果：0(符号位)000 0000 0000 0000 0000 0000 1000 0000，但是实际只有 8 位的存储空间
	// 那么就只能舍弃其他的位了，重要的就是符号位被丢弃了，所以 b变量存储的补码值就是 1000 0000，对应的十进制数就是 -128

	t.Log(-128 / -1)
	// t.Log(-128 / int8(-1)) // 编译不通过
}
