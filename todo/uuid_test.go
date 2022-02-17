package todo

import (
	"encoding/hex"
	"github.com/satori/go.uuid"
	"testing"
)

/*
首先声明，这个不是 Go 官方的类库，这是第三方类库

v4 版本的 uuid 是 32 个字符 + 4 个 '-'
*/

func TestUUID(t *testing.T) {
	uid := uuid.NewV4()

	// uid.Bytes() 可不等同于 []byte(uid.String())，原因在于 String 方法中有做处理和转换
	t.Log(hex.EncodeToString([]byte(uid.String())))
	// 除去 '-'，可以通过 hex.EncodeToString
	t.Log(hex.EncodeToString(uid.Bytes()))
}
