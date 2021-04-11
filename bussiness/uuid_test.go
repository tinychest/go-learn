package bussiness

import (
    "encoding/hex"
    "fmt"
    "github.com/satori/go.uuid"
    "testing"
)

// 首先声明，这个不是 GO 官方的类库，这是第三方类库
func TestUUID(t *testing.T) {
    uuid := uuid.NewV4()
    uuidString := uuid.String()
    uuidBytes := uuid.Bytes()

    // 长度36，除去 - 符 长度 32
    // printUUID(uuid)

    // 除去 - 符，还可以用这样方法，原理是 uuid.Bytes 方法会将 - 符纳入考虑范围之外
    // println(hex.EncodeToString(uuidBytes))

    bytes := []byte(uuidString)
    // EncodeToString：源码注释是说将每个字节转成一个16进制数的字符串来表示
    // 源码：将参数 []byte 转成另外一个 大小为原来两倍的 []byte，原理是将参数字节切片中每一个字节用两位的十六进制数来表示（不足两位在第一位补零）
    println(hex.EncodeToString(bytes))
    println(hex.EncodeToString(uuidBytes))

    // 输出：010203，这也就是为什么说先转换，再转回来内容不一样的原因
    println(hex.EncodeToString([]byte{1, 2, 3}))
    // 输出：313233，其实就是每两位 16 进制转回来就是字符对应的 ascii 码
    println(hex.EncodeToString([]byte("123")))
}

func printUUID(uuid uuid.UUID) {
    uuidString := uuid.String()

    fmt.Printf("uuid：%s len：%d\n", uuidString, len(uuidString))
}
