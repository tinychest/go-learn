package encode

import (
	"testing"
)

func TestBeforeCryptoJSAES(t *testing.T) {
	aeskey := []byte("asdaf")
	ori := []byte("123")

	enc, err := CryptoJSAESPreEncrypt(aeskey, ori)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(enc))

	dec, err := CryptoJSAESPreDecrypt(aeskey, enc)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(dec))
}
