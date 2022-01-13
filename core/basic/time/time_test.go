package time

import (
	"fmt"
	"go-learn/util"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	now := time.Now()

	yearSum, month, daySum := now.Date()
	yearSum = now.Year()
	month = now.Month()
	daySum = now.Day()
	weekday := now.Weekday()
	t.Logf("%10s: %d\n", "Year", yearSum)    // 2021
	t.Logf("%10s: %d\n", "Month", month)     // 2
	t.Logf("%10s: %d\n", "Day", daySum)      // 24
	t.Logf("%10s: %d\n", "Weekday", yearSum) // 2021
	t.Logf("%10s: %d\n", "Weak", weekday)    // 3

	// UnixNano = Unix + Nanosecond
	unixNano := now.UnixNano()
	unix := now.Unix()
	nanoSecond := now.Nanosecond()
	t.Logf("%10s: %d\n", "UnixNano", unixNano)     // 1614136265954607800
	t.Logf("%10s: %d\n", "Unix", unix)             // 1614136265
	t.Logf("%10s: %d\n", "Nanosecond", nanoSecond) // 954607800
}

func TestUnit(t *testing.T) {
	n := time.Now()
	t.Log(n.Unix())      // 秒
	t.Log(n.UnixMilli()) // 豪秒
	t.Log(n.UnixMicro()) // 微秒
	t.Log(n.UnixNano())  // 纳秒

	// 表示复数个时间单位，常量可以直接相乘，变量需要转换一下
	var d time.Duration
	var sum int64 = 1

	// d = sum * time.Second // 这个操作不允许，不能直接和变量相乘
	d = 1 * time.Second
	d = time.Duration(sum) * time.Second

	t.Log(d, sum)
}

func TestDiff(t *testing.T) {
	before := util.ParseTime("1997-10-05 11:00:00")
	after := util.ParseTime("1997-10-05 12:00:00")

	// 求两个时间相差的秒数
	t.Logf("%d\n", int(after.Sub(before).Seconds())) // 3600
	t.Logf("%d\n", after.Unix()-before.Unix())       // 3600
}

func TestOperate(t *testing.T) {
	now := time.Now()

	add := now.AddDate(0, 0, -7)
	t.Log(add)

	sub := now.Truncate(time.Hour * 24)
	t.Log(sub)
}