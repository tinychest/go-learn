package encode

import (
	"fmt"
	"testing"
)

func TestAES(t *testing.T) {
	// len(key) 要求 16 24 32
	const key = "0123456789abcdef"

	// len(ori) 要求 >= 16
	ori := "0123456789abcdef"
	enc := AESEncrypt(key, ori)
	dec := AESDecrypt(key, enc)

	fmt.Println(enc)
	fmt.Println(dec == ori)
}
