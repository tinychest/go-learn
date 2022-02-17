package http

import (
	"net/url"
	"testing"
)

/*
QueryEscape：非 数字 字母 的字符，进行编码
PathEscape：中文 特殊字符（不包括【= & @ $ _ + 】）进行编码
url.Values.Encode：特殊字符【= & _】进行编码
*/
func TestUrlParam(t *testing.T) {
	// 在 http 协议中， + / 都是特殊字符，为了避免其特殊含义，实际请求时，【'+' ←→ ' '】【'\' 会被省略】
	// 服务端请求时也会根据协议做对应的处理：【' ' ←→ '+'】，这也就是开发中可能会踩到的坑，“参数中要求指定的加号会怎么变成空格了呢，然后开始甩锅...”
	// 注意：不是说编码方式会擅自篡改字符，而是浏览器为了遵守协议，替换了字符，所以下边的测试是复现不出来，基于浏览器访问接口的问题的

	queryEscapeTest(t, " ")   // 【+】
	queryUnescapeTest(t,"+") // 【 】

	queryEscapeTest(t, "+")   // 【%2B】
	queryUnescapeTest(t, " ") // 【 】

	pathEscapeTest(t, "+")   // 【+】
	pathUnescapeTest(t, " ") // 【 】

	pathEscapeTest(t, " ")   // 【%20】
	pathUnescapeTest(t, "+") // 【+】
}

// test url.QueryEscape And url.QueryUnescape
func queryEscapeTest(t *testing.T, param string) {
	query := url.QueryEscape(param)
	t.Log("QueryEscape：【" + query + "】")
}

func queryUnescapeTest(t *testing.T, param string) {
	result, err := url.QueryUnescape(param)
	if err != nil {
		t.Fatal("QueryUnescape 失败")
	}
	t.Log("QueryUnescape：【" + result + "】")
}

// test url.PathEscape And url.UnPathEscape
func pathEscapeTest(t *testing.T, param string) {
	path := url.PathEscape(param)
	t.Log("PathEscape：【" + path + "】")
}

func pathUnescapeTest(t *testing.T, param string) {
	result, err := url.PathUnescape(param)
	if err != nil {
		t.Fatal("PathUnescape 失败")
	}
	t.Log("PathUnescape：【" + result + "】")
}

// test url.Values.Encode()
func TestUrlValuesEncode(t *testing.T) {
	// 会将特殊字符编码
	params := url.Values{
		"name": []string{"小(!@#$%^&*_+)明"},
		"age":  []string{"11"},
	}
	t.Log("url.Values.Encode：" + params.Encode())

	// 默认会根据键名的 ASCII 码从小到大排序（字典序）
	// 但是假如希望得到按键名字典序，但是不编码的，暂时除了自己通过 fmt.Sprintf 手动拼接，没有更好的办法
	params = url.Values{
		"b": []string{"B"},
		"c": []string{"C"},
		"a": []string{"A"},
	}
	t.Log(params.Encode())
}
