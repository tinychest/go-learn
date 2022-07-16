package time

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	// - time.Duration 的底层值实际是纳秒值
	second := time.Second
	t.Log(second.Nanoseconds())
	t.Log(int64(second))

	// - time.Duration 的相关转化方法，返回值的类型是 float64，代表相关方法都是全值转化的意思
	dur := 10*time.Hour + 59*time.Minute + 59*time.Second
	t.Log(dur.Hours())
	t.Log(dur.Minutes())
	t.Log(dur.Seconds())

	t.Log(int64(dur / time.Hour))
	t.Log(int64(dur % time.Hour / time.Minute))
	t.Log(int64(dur % time.Minute / time.Second))
}
