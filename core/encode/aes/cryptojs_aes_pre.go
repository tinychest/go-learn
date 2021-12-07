package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"go-learn/core/encode/eutil"
)

// CryptoJSAESPreEncrypt
// 为了达到各项参数没有各项严谨的长度限制，需要在进行真正的原生 AES 算法处理之前，进行长度的处理（MD5 HASH 等给出特定位数结果的算法）

func CryptoJSAESPreEncrypt(aesKey, cipherText []byte) ([]byte, error) {
	iv        := eutil.IV()
	aesKey     = eutil.AESKey(aesKey)
	cipherText = eutil.Cipher(cipherText)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText, cipherText)

	// 最终结果 - 拼装
	result := append(iv, cipherText...)
	// 最终结果 - base64
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(result)))
	base64.StdEncoding.Encode(buf, result)
	return buf, nil
}

func CryptoJSAESPreDecrypt(aesKey, encrypted []byte) ([]byte, error) {
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(encrypted)))
	n, err := base64.StdEncoding.Decode(buf, encrypted)
	if err != nil {
		return nil, err
	}
	// 这一步很重要，因为多出了一位，base64.StdEncoding.DecodedLen 返回的不准确
	buf = buf[:n]
	if len(buf) < aes.BlockSize {
		return nil, errors.New("非法的密文格式")
	}

	aesKey = eutil.AESKey(aesKey)
	iv := buf[:aes.BlockSize]
	encryptText := buf[aes.BlockSize:]

	// 解密
	if len(encryptText) == 0 || len(encryptText)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("bad blocksize(%v), aes.BlockSize = %v\n", len(encryptText), aes.BlockSize)
	}
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCDecrypter(block, iv)
	cbc.CryptBlocks(encryptText, encryptText)
	return eutil.Pkcs7Unpad(encryptText, aes.BlockSize)
}
