package eutil

import (
	"crypto/aes"
	"crypto/sha256"
	"go-learn/util"
)

// IV 生成随机的向量 [16]byte
func IV() []byte {
	return util.RandomBytes(aes.BlockSize)
}

// Cipher 密文的长度得是 blockSize 的整倍数
func Cipher(cipher []byte) []byte {
	return Pkcs7Pad(cipher, aes.BlockSize)
}

// AESKey 合法长度 16 24 32 中的 32
func AESKey(aesKey []byte) []byte {
	bs := sha256.Sum256(aesKey)
	return bs[:]
}
