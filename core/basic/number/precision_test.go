package number

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

const sum = 0.1

func TestPrintType(t *testing.T) {
	fmt.Println(1e-1)

	typ := reflect.TypeOf(sum)
	fmt.Println(typ)

	fmt.Println(0.8999999999+1e-9 >= 0.9)
}

// 精度
// 处理数位计算，在 Java 中是使用 BigDecimal，或者通用的是说，当在一定精度内相等，就认为是相等
func TestCal(t *testing.T) {
	sum1 := "0.1"
	sum2 := "0.2"
	sum3 := "0.3"

	s1, _ := strconv.ParseFloat(sum1, 64)
	s2, _ := strconv.ParseFloat(sum2, 64)
	s3, _ := strconv.ParseFloat(sum3, 64)

	// 编译其不提示，但是实际执行结果为 false
	fmt.Println(s1+s2 == s3)

	// 编译器提示 false，但是实际执行结果为 true
	fmt.Println(0.1+0.2 == 0.3)
}
