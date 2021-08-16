package encode

import (
	"encoding/base64"
	"fmt"
	"testing"
)

/*
一、简介
Base64 编码是网络上最常见的用于传输 8Bit 字节码的编码方式之一
是从二进制到字符的过程 ，可用于在 HTTP 环境下传递较长的标识信息
采用 Base64 编码具有不可读性，需要解码后才能阅读

二、base64 类库中提供的加解码实例
StdEncoding
对应协议：RFC 4648
密文内容字符集：ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/

URLEncoding
对应协议：RFC 4648
密文内容字符集：ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_

1、区分一
StdEncoding URLEncoding
Std 和 URL 的区别在于密文内容的字符集，实际看下来就是一个字符集的特殊字符是 + / 而另一个是 - _

2、区分二
StdEncoding URLEncoding RawStdEncoding RawURLEncoding
名字打头是否带 Raw 的区别在于编解码时的填充字符是什么。不带 Raw：=，带 Raw：无填充字符

其他：
https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Basics_of_HTTP/Data_URIs

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

	printlnMsg(rawMsg, "Encoding encoded to", encodingMsg)
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
	printlnMsg(encodingMsg, "Encoding decoded to", dstStr)
	return dstStr
}

func printlnMsg(raw, prefix, after string) {
	if len(raw) == 0 {
		raw = "空"
	}
	if len(after) == 0 {
		after = "空"
	}
	fmt.Printf("【%s】 -- %s ->【%s】\n", raw, prefix, after)
}
