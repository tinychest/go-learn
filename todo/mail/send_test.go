package mail

import "testing"

func TestSend(t *testing.T) {
	// Strings.Builder 的长度为 290
	SendTo("mincong.wang@shanghairanking.com", "小明", "123456")
}
