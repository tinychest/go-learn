package yaml

import (
	"encoding/json"
	"testing"
)

// 测试结论：
// 1、支持 yaml 标签，不支持 yml 标签
// 2、映射关系：配置字段名 原封不动；结构体字段名、无标签=转化为全小写、有标签=转化为大写
type YS struct {
	TestTag1 string `yaml:"abc"`
	TestTag2 string `yml:"abc"`

	// test_match、testmatch、testMatch
	TestMatch string
}

var ys = new(YS)

// yaml 解析中 配置项的字段名和结构体字段名的（有标签就按照标签）
func TestBasic(t *testing.T) {
	configLoad("abc.yml", ys)
	result, _ := json.Marshal(ys)

	t.Log(string(result))
}
