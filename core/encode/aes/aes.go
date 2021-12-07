package aes

import (
	"crypto/aes"
)

// Go 提供的基础 AES 算法的实现

func AESEncrypt(key, word string) string {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	dst := make([]byte, len(word))
	// 要求 len(dst) 和 len(src) 都必须 >= 16
	cipher.Encrypt(dst, []byte(word))
	return string(dst)
}

func AESDecrypt(key, word string) string {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	dst := make([]byte, len(word))
	// 要求 len(dst) 和 len(src) 都必须 >= 16
	cipher.Decrypt(dst, []byte(word))
	return string(dst)
}
