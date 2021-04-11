package _proxy

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// 这里的动态代理测试一下还没有完全吃透，目前就是搞清楚了程序功能结构，等有切实需求再来捡起来
// 为指定的 源码文件中 含有特定格式注释的结构体 生成代理类的源码
func Test_generate(t *testing.T) {
	want := `package _proxy

type UserProxy struct {
	child *User
}

func NewUserProxy(child *User) *UserProxy {
	return &UserProxy{child: child}
}

func (p *UserProxy) Login(username, password string) (r0 error) {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	r0 = p.child.Login(username, password)

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return r0
}
`
	got, err := generate("./static_proxy.go")
	require.Nil(t, err)
	assert.Equal(t, want, got)
}
