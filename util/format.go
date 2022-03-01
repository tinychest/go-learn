package util

import (
	"go-learn/const/time_format"
	"time"
)

func ParseTime(s string) time.Time {
	return ParseFmtTime(s, time_format.FmtDateTime)
}

func ParseFmtTime(s string, format string) time.Time {
	if t, err := time.Parse(format, s); err != nil {
		panic(err)
	} else {
		return t
	}
}

func FormatTime(t time.Time, format string) string {
	return t.Format(format)
}

func Now() string {
	return FormatTime(time.Now(), time_format.FmtDateTime)
}

func NowCN() string {
	return FormatTime(time.Now(), time_format.FmtDateTimeCN)
}

func NowDate() string {
	return FormatTime(time.Now(), time_format.FmtDate)
}

func NowDateCN() string {
	return FormatTime(time.Now(), time_format.FmtDateCN)
}
