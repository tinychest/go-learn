package time

import (
	"fmt"
	"testing"
	"time"
)

func TestOperateTime(t *testing.T) {
	date := time.Now().AddDate(0, 0, -7)

	edgeTime := date.Truncate(time.Hour * 24)
	// golang 编译通过，go sdk 编译不通过，println 方法的类泛型参数类型定义，感觉只能接受最基本的数据类型
	// println(date)
	fmt.Println(date)
	// println(edgeTime)
	fmt.Println(edgeTime)
}
