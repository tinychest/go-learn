package encode

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"testing"
)
/*
《概念》
1.Golang 默认不支持 UTF-8 以外的字符集
2.做编码转换，使用的是 transform 包

运行下面示例会提示你，下载必须的包
go get golang.org/x/text/encoding/simplifiedchinese@v0.3.3
go get golang.org/x/text/transform@v0.3.3
*/

func TestUTF8ToGBK(t *testing.T) {
	// UTF8 → GBK
	result, n, err := transform.String(simplifiedchinese.GBK.NewEncoder(), "小明")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n)
	fmt.Println(result)

	// GBK → UTF8（需要借助三方包 mahonia）
}
