package encode

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestSHA256(t *testing.T) {
	password := "sda7f65asf7"
	// 注意参数类型是 int[] 返回值类型是 int[32]
	sum256 := sha256.Sum256([]byte(password))

	// println 不接受数组类型（编译不通过）
	println(sum256[:])
}

func TestSha1(t *testing.T) {
	param := "jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VPDfmWIpmAiMYVmDhkD5T2CbgFVT7iXnWQegREvWF7GWn-X97D_WnSbgMU4CxZACEg&noncestr=123&timestamp=123&url=http://qq.com/"

	theSha1 := sha1.New()
	if _, err := theSha1.Write([]byte(param)); err != nil {
		t.Fatal(err)
	}
	println(hex.EncodeToString(theSha1.Sum(nil)))

	// 每次都能得到一个恒定的结果：da39a3ee5e6b4b0d3255bfef95601890afd80709
	fmt.Println(hex.EncodeToString(sha1.New().Sum(nil)))
	// 这种是错误的，不能这样去写
	fmt.Println(hex.EncodeToString(sha1.New().Sum([]byte(param))))
}
