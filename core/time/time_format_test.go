package time

import (
	"go-learn/util"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	dateAndTimeString1 := getCurrentFormatTimeString(util.DateTimeFormat)
	dateAndTimeString2 := getCurrentFormatTimeString(util.DateTimeFormat)
	dateString := getCurrentFormatTimeString(util.DateFormat)
	timeString := getCurrentFormatTimeString(util.TimeFormat)

	println(dateAndTimeString1)
	println(dateAndTimeString2)
	println(dateString)
	println(timeString)
}

func getCurrentFormatTimeString(format string) string {
	return time.Now().Format(format)
}
