package encode

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

// TestMD5
func TestMD5(t *testing.T) {
	word := "123"
	t.Logf("%s MD5 To %s\n", word, MD5(word))
	t.Logf("%s MD5 To %s\n", word, MD52(word))
}

// 这个方法等同于前端 CryptoJS.MD5 方法
func MD5(word string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(word)))
}

func MD52(word string) string {
	// hex.EncodeToString(md5.Sum([]byte(word))[:])
	bs := md5.Sum([]byte(word))
	return hex.EncodeToString(bs[:])
}
