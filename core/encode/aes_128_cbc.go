package encode

import (
	"crypto/aes"
	"crypto/cipher"
	"go-learn/core/encode/eutil"
)

// 对称加密算法：encrypted with OpenSSL using AES-256-CBC and a passphrase
// AES：标注了其基本算法
// 256：中间的数字其实是由 aesKey 的长度决定的，如：AES-128 AES-192 AES-256

// 解密要素
// - aesKey（密钥，16 字节）
// - iv（初始向量）
// - cipherText（密文）
// 		len(cipherText) >= 16（aes.BlockSize）
//		iv：cipherText[:16]
//		real_cipherText = cipherText[16:]
//		len(real_cipherText) 是 16 的整倍数

// 解密方法
func AES128CBC_PKCS7(aesKey, iv, cipherText []byte) []byte {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		panic(err)
	}
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(cipherText, cipherText)

	cipherText, err = eutil.Pkcs7Unpad(cipherText, block.BlockSize())
	if err != nil {
		panic(err)
	}
	return cipherText
}
