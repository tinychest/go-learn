package encode

import (
	"encoding/base64"
	"fmt"
)

// 编码
func encoding(encoding *base64.Encoding, rawMsg string) string {
	src := []byte(rawMsg)

	// 方式一：Encode
	// dstLen := base64.StdEncoding.EncodedLen(len(src))
	// dst := make([]byte, dstLen)
	// base64.StdEncoding.Encode(dst, src)
	// res := string(dst)

	// 方式二：EncodeToString
	res := encoding.EncodeToString(src)

	printlnProcess(rawMsg, "encoding to", res)
	return res
}

// 解码
func decoding(encoding *base64.Encoding, encodingMsg string) string {
	// 方式一：Decode
	// src := []byte(encodingMsg)
	// dstLen := base64.StdEncoding.DecodedLen(len(src))
	// dst := make([]byte, dstLen)
	// if _, err := encoding.Decode(dst, src); err != nil {
	// 	panic(err)
	// }
	// res := string(dst)

	// 方式二：DecodeToString
	dst, err := encoding.DecodeString(encodingMsg)
	if err != nil {
		panic(err)
	}
	res := string(dst)

	printlnProcess(encodingMsg, "decoding to", res)
	return res
}

func printlnProcess(raw, prefix, after string) {
	if len(raw) == 0 {
		raw = "空"
	}
	if len(after) == 0 {
		after = "空"
	}
	fmt.Printf("【%s】-- %s -->【%s】\n", raw, prefix, after)
}
