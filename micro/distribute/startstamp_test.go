package distribute

import (
	"testing"
	"time"
)

func TestTimeMilli(t *testing.T) {
	ti, err := time.Parse("2006-01-02 15:04:05.999", "2022-06-01 00:00:00.000")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ti.Unix())      // 距离 1970-01-01 00:00:00.000 的毫值
	t.Log(ti.UnixMilli()) // 距离 1970-01-01 00:00:00.0... 的毫秒值：1654041600000
	t.Log(ti.UnixMicro()) // 距离 1970-01-01 00:00:00.0... 的微秒值
	t.Log(ti.UnixNano())  // 距离 1970-01-01 00:00:00.0... 的纳秒值
}
