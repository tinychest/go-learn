package encode

import (
	"encoding/base64"
	"fmt"
	"testing"
)

/*
【简介】
Base64 编码是网络上最常见的用于传输 8Bit 字节码的编码方式之一
是从二进制到字符的过程 ，可用于在 HTTP 环境下传递较长的标识信息
采用 Base64 编码具有不可读性，需要解码后才能阅读

【base64 类库中提供的加解码实例】
[StdEncoding]
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/（填充 =）
[StdEncoding]
字符集相同，不填充
[URLEncoding]
ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_（填充 =）
[RawURLEncoding]
字符集相同，不填充

【相关资料】
详见 Base64.md

【Base64 数据的内嵌样例】
html：<img width="40" height="30" src="data:image/jpg;base64,/9j/4QMZRXhpZgAASUkqAAgAAAAL...." /">
css：.demoImg{ background-image: url("data:image/jpg;base64,/9j/4QMZRXhpZgAASUkqAAgAAAAL...."); }
*/
func TestBase64Encode(t *testing.T) {
	// 带不带 Raw 的区别
	add1 := encoding(base64.URLEncoding, "+") // 编码为：Kw==
	decoding(base64.URLEncoding, add1)        // 能够正常解码出：+

	add2 := encoding(base64.RawURLEncoding, "+") // 编码为：Kw
	decoding(base64.RawURLEncoding, add2)        // 能够正常解码出：+
}

// 编码
func encoding(encoding *base64.Encoding, rawMsg string) string {
	src := []byte(rawMsg)

	// Encode
	dstLen := base64.StdEncoding.EncodedLen(len(src))
	dst := make([]byte, dstLen)
	base64.StdEncoding.Encode(dst, src)

	// EncodeToString
	encodingMsg := encoding.EncodeToString(src)

	printlnProcess(rawMsg, "Encoding encoded to", encodingMsg)
	return encodingMsg
}

// 解码
func decoding(encoding *base64.Encoding, encodingMsg string) string {
	src := []byte(encodingMsg)

	// Decode
	dstLen := base64.StdEncoding.DecodedLen(len(src))
	dst := make([]byte, dstLen)
	if _, err := encoding.Decode(dst, src); err != nil {
		panic(err)
	}

	// DecodeToString
	var err error
	if dst, err = encoding.DecodeString(encodingMsg); err != nil {
		panic(err)
	}

	dstStr := string(dst)
	printlnProcess(encodingMsg, "Encoding decoded to", dstStr)
	return dstStr
}

func printlnProcess(raw, prefix, after string) {
	if len(raw) == 0 {
		raw = "空"
	}
	if len(after) == 0 {
		after = "空"
	}
	fmt.Printf("【%s】 -- %s ->【%s】\n", raw, prefix, after)
}
