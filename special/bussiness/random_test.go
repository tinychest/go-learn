package bussiness

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	mathRand "math/rand"
	"testing"
	"time"
)

// 一、关于随机数的一些点：
// 如果没有设置 seed，则默认 seed 为 1

// 对于不涉及到密码类的开发工作直接使用 math/rand + 基于时间戳的种子 rand.Seed(time.Now().UnixNano()) 一般都能满足需求
// 对于涉及密码类的开发工作一定要用 crypto/rand
// 如果想生成随机字符串，可以先列出字符串，然后基于随机数选字符的方式实现

// 二、关于 Go 随机数，这里有一个精华 https://medium.com/a-journey-with-go/go-how-are-random-numbers-generated-e58ee8696999
func TestRand(t *testing.T) {
	// 随机大数（和 int_oto_string_test 中不同，这里能够表示的最大的进制：10 + 26 + 26 = 62）
	// bit, _ := new(big.Int).SetString(strings.Repeat("Z", 99), 62)
	// cryptoRandInt, _ := cryptoRand.Int(cryptoRand.Reader, bit)

	mathRand.Seed(time.Now().UnixNano())

	for i := 0; i < 5; i++ {
		fmt.Print(fakeRandomSum1(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(fakeRandomSum2(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(randomSum1(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(randomSum2(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(randomSum3(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(randomSum4(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(randomSum5(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(verifyCode(6), " ")
	}
	println()
}

// 假随机数：不是说循环得到的数一样的，或是有规律的，而是你多次运行这个函数，会发现得到的结果是相同的
// 其实函数得到什么结果是固定的，是根据“种子”概念来的，决定了种子就是决定了一连串固定的事件、场景、操作（很多游戏都有这个机制：以撒）
// 为了使概念好展开，就拿以撒这个游戏继续说，以撒是怎么做到，每次开局的内容（房间、怪物、道具）都是不同的，即为什么每次开局，游戏的种子都是不同的 - 没错，时间
// （其实再往里边说，即使是同一个种子，怪物变异、攻克房间的基础掉落物也是不同的，那这里的随机概念，也和时间脱不了干系）
// 你每次开始游戏的时间都是不可能相同的，“这就像一个可以驱使既定程序改变的外在驱动力”
func fakeRandomSum1(bit int) int {
	max := int(math.Pow10(bit)) - 1
	mathRand.Seed(2)
	return mathRand.Intn(max)
}

// 和时间挂钩，每次运行程序的种子值都是不同的，那么这个程序每次运行的结果都是不同的
// 但是实际，在循环中调用这个方法，还是会获得相同的结果，也就说明自己觉得是的，实际并不是
// 进一步测试，发现 time.Now().UnixNano() 在短时间调用，返回的结果是相同的
func fakeRandomSum2(bit int) int {
	max := int(math.Pow10(bit)) - 1
	return mathRand.New(mathRand.NewSource(time.Now().UnixNano())).Intn(max)
}

// 上边错就错在，应该就一个种子不断的获取随机数，而不是每次重置种子，从头拿数
func randomSum1(bit int) int {
	max := int(math.Pow10(bit)) - 1
	return mathRand.Intn(max)
}

// 问题是使用 % 会存在指定范围内的数被随机到的概率是不同的
func randomSum2(bit int) int64 {
	// rand.Intn 底层就是这个，没什么好说的
	max := int(math.Pow10(bit)) - 1
	return mathRand.Int63() % int64(max)
}

// 这个概率也不是平等的，但是在生成指定位数的随机码的时候，位数越多，同一个数字一直出现（最终生成的随机码相同）的概率会逐步降低
func randomSum3(bit int) int64 {
	max := int(math.Pow10(bit)) - 1
	return mathRand.Int63() & int64(max)
}

// crypto/rand是为了提供更好的随机性满足密码对随机数的要求，在linux上已经有一个实现就是/dev/urandom，crypto/rand 就是从这个地方读“真随机”数字返回，但性能比较慢
// 注意要更换导包为 "crypto/rand"
func randomSum4(bit int) int64 {
	max := int(math.Pow10(bit)) - 1
	n, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max)))
	return n.Int64()
}

func randomSum5(bit int) int {
	max := int(math.Pow10(bit)) - 1

	res := make([]byte, 4*4)
	_, _ = cryptoRand.Read(res)

	return int(binary.BigEndian.Uint32(res)) % max
}

const VerifyCodeRune = "0123456789"

var rd = mathRand.New(mathRand.NewSource(time.Now().UnixNano()))

func verifyCode(bit int) string {
	bs := make([]byte, bit)
	for i := range bs {
		bs[i] = VerifyCodeRune[rd.Intn(len(VerifyCodeRune))]
	}
	return string(bs)
}
