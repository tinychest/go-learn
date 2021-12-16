package time

import (
	"encoding/json"
	"fmt"
	"go-learn/const/time_format"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	dateAndTimeString1 := forNow(time_format.FmtDateTime)
	dateAndTimeString2 := forNow(time_format.FmtDateTime)
	dateString := forNow(time_format.FmtDate)
	timeString := forNow(time_format.FmtTime)

	println(dateAndTimeString1)
	println(dateAndTimeString2)
	println(dateString)
	println(timeString)
}

func forNow(format string) string {
	return time.Now().Format(format)
}

// time 打印默认的格式：2006-01-02 15:04:05.999999999 -0700 MST（详见 Datetime.String）
// time json.Marshal 默认的格式：time.RFC3339Nano（详见 Datetime.MarshalJSON）
func TestPrintTime(t *testing.T) {
	// 直接 fmt.Println(time) 和 json.Marshal(time) 的结果是不同的
	// 结论：time.Datetime 实现了 MarshalJSON 方法
	n := time.Now()
	nJSON, _ := json.Marshal(n)

	fmt.Println(n)
	fmt.Println(string(nJSON))
}
