package hw

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestSha256(t *testing.T) {
	var json1, json2 string

	// Golang  ：a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
	// Java 原生：a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
	// Java 华为：a665a45920422f9d417e4867efdc4fb8a04a1f3fff1fa07e998e86f7f7a27ae3
	json1 = `123`
	// 和 Java 代码得出的结果一致
	// Golang  ：ee24506edc8b6f9943a3a9495a1e6b7b6475445db896b52f9dd1f62c0932c1d5
	// Java 原生：ee24506edc8b6f9943a3a9495a1e6b7b6475445db896b52f9dd1f62c0932c1d5
	// Java 华为：ee24506edc8b6f9943a3a9495a1e6b7b6475445db896b52f9dd1f62c0932c1d5
	json2 = `{"categories":["porn","politics","ad","abuse","contraband","flood"], "items":[{"text":"666666luo聊请+110亚砷酸钾六位qq，fuck666666666666666","type":"content"}]}`

	sha256Bytes := sha256.Sum256([]byte(json1))
	contentSha256 := hex.EncodeToString(sha256Bytes[:])
	t.Log(json1)
	t.Log(json2)
	t.Log(contentSha256)
}
