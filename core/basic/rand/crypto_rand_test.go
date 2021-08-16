package rand

import (
	cryptoRand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
)

func TestCryptoRand(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Print(cryptoRand1(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(cryptoRand2(6), " ")
	}
	println()
	for i := 0; i < 5; i++ {
		fmt.Print(cryptoRand3(6), " ")
	}
	println()
}

// crypto/rand 是为了提供更好的随机性满足密码对随机数的要求，在 linux 上已经有一个实现就是 /dev/urandom
// crypto/rand 就是从这个地方读“真随机”数字返回，但性能比较慢
func cryptoRand1(bit int) int64 {
	max := int(math.Pow10(bit)) - 1
	n, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(int64(max)))
	return n.Int64()
}

func cryptoRand2(bit int) int {
	max := int(math.Pow10(bit)) - 1

	res := make([]byte, 4*4)
	_, _ = cryptoRand.Read(res)

	return int(binary.BigEndian.Uint32(res)) % max
}

func cryptoRand3(bit int) string {
	max := "1" + strings.Repeat("0", bit)
	// SetString 参数1：生成随机数的最大值
	// SetString 参数2：随机大数（和 int_oto_string_test 中不同，这里能够表示的最大的进制：10 + 26 + 26 = 62）
	bi, _ := new(big.Int).SetString(max, 10)
	cryptoRandInt, _ := cryptoRand.Int(cryptoRand.Reader, bi)
	return cryptoRandInt.String()
}
