package regex

import (
	"encoding/json"
	"regexp"
	"testing"
)

// 贪婪模式
// 非贪婪模式，和 ? 相关
// 独占模式，和 + 相关

func TestGreed(t *testing.T) {
	r := `(hello world) hello goland (hello world)`

	greed := regexp.MustCompile(`\(.+\)`)
	notGreed := regexp.MustCompile(`\(.+?\)`)

	res, _ := json.Marshal(greed.FindAllString(r, -1))
	t.Log(string(res))

	res, _ = json.Marshal(notGreed.FindAllString(r, -1))
	t.Log(string(res))
}