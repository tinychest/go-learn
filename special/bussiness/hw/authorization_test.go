package hw

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

const (
	Ak = "5FYXJPTSGTZKT0WRHWW6"
	Sk = "89ZKJvPelZuJ8eS5bqLg1fXRQkJXtHgkXY8zKnym"
)

func TestAuth(t *testing.T) {
	json := `{"categories":["porn","politics","ad","abuse","contraband","flood"],"items":[{"text":"666666luo聊请+110亚砷酸钾六位qq，fuck666666666666666","type":"content"}]}`

	Authorization("20210323T093454Z", []byte(json), "moderation.cn-north-1.myhuaweicloud.com")
}

func Authorization(xSdkDate string, jsonParamBytes []byte, host string) string {
	// 第 1 阶段
	sha256Bytes := sha256.Sum256(jsonParamBytes)
	contentSha256 := hex.EncodeToString(sha256Bytes[:])

	// 华为：a369bc8efb50cf243fb852caa580f52691ec37c277e87d057173d42960dd259b
	// 这里：a369bc8efb50cf243fb852caa580f52691ec37c277e87d057173d42960dd259b
	fmt.Println(contentSha256)

	// 第 2 阶段
	builder := strings.Builder{}
	builder.WriteString(http.MethodPost)
	builder.WriteString("\n")
	builder.WriteString("/v1.0/moderation/text/")
	builder.WriteString("\n\n")
	builder.WriteString("content-type:application/json; charset=UTF-8")
	builder.WriteString("\n")
	builder.WriteString("host:" + host)
	builder.WriteString("\n")
	builder.WriteString("x-sdk-date:" + xSdkDate)
	builder.WriteString("\n\n")
	builder.WriteString("content-type;host;x-sdk-date")
	builder.WriteString("\n")
	builder.WriteString(contentSha256)
	canonicalRequest := builder.String()

	// 第 3 阶段
	sha256Bytes = sha256.Sum256([]byte(canonicalRequest))
	contentSha256 = hex.EncodeToString(sha256Bytes[:])

	builder = strings.Builder{}
	builder.WriteString("SDK-HMAC-SHA256")
	builder.WriteString("\n")
	builder.WriteString(xSdkDate)
	builder.WriteString("\n")
	builder.WriteString(contentSha256)
	stringToSign := builder.String()

	signature := hmacSha256(stringToSign, Sk)

	// 第 4 阶段
	builder = strings.Builder{}
	builder.WriteString("SDK-HMAC-SHA256")
	builder.WriteString(" ")
	builder.WriteString("Access=" + Ak)
	builder.WriteString(", ")
	builder.WriteString("SignedHeaders=content-type;host;x-sdk-date")
	builder.WriteString(", ")
	builder.WriteString("Signature=" + signature)
	authorization := builder.String()
	return authorization
}
