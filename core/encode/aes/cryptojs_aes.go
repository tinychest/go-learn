package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"go-learn/core/encode/eutil"
	"go-learn/util"
)

// 参考（主要）：现成的 Go 实现代码
// 参考：github.com/silenceper/wechat/v2
// 参考：https://dequeue.blogspot.com/2014/11/decrypting-something-encrypted-with.html

const (
	// OpenSSL salt is always this string + 8 bytes of actual salt
	// 即`\x53\x61\x6c\x74\x65\x64\x5f\x5f`，[0x53616c74, 0x65645f5f] 提取自 node_modules/crypto-js/crypto-js.js:3847
	openSSLSaltHeader = "Salted__"
)

func CryptoJSAESEncrypt(aesKey, cipherText []byte) ([]byte, error) {
	salt := util.RandomBytes(8)
	aesKey, iv := AESKeyAndIV(aesKey, salt)
	cipherText = eutil.Cipher(cipherText)

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	cbc.CryptBlocks(cipherText, cipherText)

	result := append([]byte(openSSLSaltHeader), salt...)
	result = append(result, cipherText...)
	buf := make([]byte, base64.StdEncoding.EncodedLen(len(result)))
	base64.StdEncoding.Encode(buf, result)
	return buf, nil
}

func CryptoJSAESDecrypt(aesKey, encrypted []byte) ([]byte, error) {
	// 一、base64 解密
	buf := make([]byte, base64.StdEncoding.DecodedLen(len(encrypted)))
	n, err := base64.StdEncoding.Decode(buf, encrypted)
	if err != nil {
		return nil, err
	}
	buf = buf[:n]
	if len(buf) < aes.BlockSize {
		return nil, errors.New("非法的密文格式1")
	}

	// 二、取得密文 和 盐
	saltWithPrefix := buf[:aes.BlockSize]
	encryptText := buf[aes.BlockSize:]
	saltPrefix := saltWithPrefix[:8]
	salt := saltWithPrefix[8:]

	if string(saltPrefix) != openSSLSaltHeader {
		// data not appear to have been encrypted with OpenSSL, salt header missing.
		return nil, errors.New("非法的密文格式2")
	}
	aesKey, iv := AESKeyAndIV(aesKey, salt)

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

func AESKeyAndIV(aesKey, salt []byte) ([]byte, []byte) {
	var (
		m    = make([]byte, 48)
		prev []byte
	)
	for i := 0; i < 3; i++ {
		a := make([]byte, len(prev)+len(aesKey)+len(salt))
		copy(a, prev)
		copy(a[len(prev):], aesKey)
		copy(a[len(prev)+len(aesKey):], salt)

		h := md5.New()
		h.Write(a)
		prev = h.Sum(nil)

		copy(m[i*16:], prev)
	}
	return m[:32], m[32:]
}
