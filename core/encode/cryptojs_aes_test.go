package encode

import (
	"testing"
)

func TestCryptoJSAES(t *testing.T) {
	aeskey := []byte("asdaf")
	ori := []byte("123")

	enc, err := CryptoJSAESEncrypt(aeskey, ori)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(enc))

	dec, err := CryptoJSAESDecrypt(aeskey, enc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(dec))
}
