package eutil

import (
	"bytes"
	"errors"
)

var (
	// ErrInvalidBlockSize block size不合法
	ErrInvalidBlockSize = errors.New("invalid block size")
	// ErrInvalidPKCS7Data PKCS7数据不合法
	ErrInvalidPKCS7Data = errors.New("invalid PKCS7 data")
	// ErrInvalidPKCS7Padding 输入padding失败
	ErrInvalidPKCS7Padding = errors.New("invalid padding on input")
)

// 这里的写法可以从很多地方看到：
// 三方微信小程序 SDK:github.com/silenceper/wechat/v2
// 博文代码示例：https://dequeue.blogspot.com/2014/11/decrypting-something-encrypted-with.html

// Pkcs7Unpad returns slice of the original data without padding
func Pkcs7Unpad(data []byte, blockSize int) ([]byte, error) {
	if blockSize <= 0 {
		return nil, ErrInvalidBlockSize
	}
	if len(data)%blockSize != 0 || len(data) == 0 {
		return nil, ErrInvalidPKCS7Data
	}
	c := data[len(data)-1]
	n := int(c)
	if n == 0 || n > len(data) {
		return nil, ErrInvalidPKCS7Padding
	}
	for i := 0; i < n; i++ {
		if data[len(data)-n+i] != c {
			return nil, ErrInvalidPKCS7Padding
		}
	}
	return data[:len(data)-n], nil
}

func Pkcs7UnpadSimple(data []byte) []byte {
	length := len(data)
	padRune := int(data[length-1])
	return data[:length - padRune]
}

func Pkcs7Pad(data []byte, blockSize int) []byte {
	padRune := blockSize - len(data)%blockSize
	padding := bytes.Repeat([]byte{byte(padRune)}, padRune)
	return append(data, padding...)
}