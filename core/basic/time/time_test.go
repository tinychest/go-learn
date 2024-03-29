package time

import (
	"go-learn/tool"
	"testing"
	"time"
)

// - Go(time.Time) - MySQL(timestamp) 默认转换关系，MySQL 驱动实际是调用 Unix 还是 哪个方法序列化入库（实际是 time.Time.Unix，也就是只到秒）
//
// 毫秒 3 个 0
// 微秒 6 个 0
// 纳秒 9 个 0

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
	t.Logf("%10s: %d\n", "UnixNano", now.UnixNano())     // 1614136265954607800
	t.Logf("%10s: %d\n", "Unix", now.Unix())             // 1614136265
	t.Logf("%10s: %d\n", "Nanosecond", now.Nanosecond()) // 954607800

	// [time.Time]
	// n := time.Now()
	// t.Log(n.Unix())      // 秒
	// t.Log(n.UnixMilli()) // 豪秒
	// t.Log(n.UnixMicro()) // 微秒
	// t.Log(n.UnixNano())  // 纳秒

	// [time.Duration]
	// s := time.Second
	// t.Log(s.Seconds())      // 秒
	// t.Log(s.Milliseconds()) // 毫秒
	// t.Log(s.Microseconds()) // 微秒
	// t.Log(s.Nanoseconds())  // 纳秒

	// [number → time.Duration]
	// t.Log(int(time.Second)) // 1? 错，单位是纳秒
	// t.Log(time.Duration(1))
	// t.Log(time.Nanosecond)

	// [实际开发中遇到的问题]
	// 表示复数个时间单位，常量可以直接相乘，变量则需要先转换一下
	// var d time.Duration
	// var sum int64 = 1
	//
	// d = 1 * time.Second
	// // d = sum * time.Second // 这个操作不允许，不能直接和变量相乘
	// d = time.Duration(sum * time.Second.Nanoseconds())
	// t.Log(d)

	// [时间戳 和 time.Time 互转]
	// now1 := time.Now()
	// timestamp := now1.Unix()
	// now2 := time.Unix(timestamp, 0)
	//
	// t.Log(now1.String())
	// t.Log(now2.String())
}

func TestDiff(t *testing.T) {
	before := tool.ParseTime("1997-10-05 11:00:00")
	after := tool.ParseTime("1997-10-05 12:00:00")

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
