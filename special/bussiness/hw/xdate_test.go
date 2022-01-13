package hw

import (
	"testing"
	"time"
)

// 华为：20210308T024104Z
// 本地：20210308T024550Z
func TestXDate(t *testing.T) {
	t.Log(XSdkData())
}

func XSdkData() string {
	return time.Now().UTC().Format("20060102T150405Z")
}
