package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
	"time"
)

// 参照：https://mp.weixin.qq.com/s/mJqUE_EIEFDUi10dy4yghQ

const (
	Key = "123"
)

var (
	KeyFunc = func(token *jwt.Token) (interface{}, error) { return []byte(Key), nil }
)

// TokenInfo token 中存储的自定义信息
// jwt 中这个称之为 Claims，要求实现自定义 Valid 方法，但是可以通过内嵌 StandardClaim 来使用默认的协议行为
type TokenInfo struct {
	jwt.StandardClaims
	Id int64 `json:"id"`
}

func TestJWT(t *testing.T) {
	info := &TokenInfo{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() - 1, // 过期时间（整好等于，不算过期）
			Issuer:    "test",
		},
		Id: 1,
	}

	tokenStr, err := Sign(info)
	if err != nil {
		fmt.Println("签发失败", err)
	}

	tokenInfo, err := Parse(tokenStr)
	if err != nil {
		fmt.Println("解析失败", err)
		return
	}

	fmt.Printf("%+#v", tokenInfo)
}

// Sign 签发
func Sign(claims *TokenInfo) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(Key))
}

// Parse 校验解析
func Parse(tokenStr string) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenInfo{}, KeyFunc)
	return token.Claims.(*TokenInfo), err
}
