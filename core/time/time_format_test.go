package time

import (
	"go-learn/const/time_format"
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	dateAndTimeString1 := getCurrentFormatTimeString(time_format.DateTimeFormat)
	dateAndTimeString2 := getCurrentFormatTimeString(time_format.DateTimeFormat)
	dateString := getCurrentFormatTimeString(time_format.DateFormat)
	timeString := getCurrentFormatTimeString(time_format.TimeFormat)

	println(dateAndTimeString1)
	println(dateAndTimeString2)
	println(dateString)
	println(timeString)
}

func getCurrentFormatTimeString(format string) string {
	return time.Now().Format(format)
}
