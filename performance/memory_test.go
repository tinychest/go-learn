package performance

import (
	"runtime"
	"testing"
)

// 运行时，程序申请占用的内存大小
func TestPrintMem(t *testing.T) {
	t.Helper()
	var rtm = new(runtime.MemStats)
	runtime.ReadMemStats(rtm)
	t.Logf("%.2f MB", float64(rtm.Alloc)/1024./1024.)
}
