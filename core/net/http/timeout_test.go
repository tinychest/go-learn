package http

import (
	"io"
	"net/http"
	"testing"
	"time"
)

// TODO 源码通读（连接池概念 - 连接复用测试 TCP 连接设置的超时时间）
// TODO 当前工具库，需要暴露出一个统一设置超时时间的配置项 或 参数

/*
http.Client.Timeout：每次请求的超时时间，不论发起请求、读取等等，就是客户端完成请求方法的 deadline
http.Client.Transport：控制更精细的 TCP 阶段的超时时间。默认值是 http.DefaultTransport（类型是 http.Transport）
  http.Transport.DialContext 默认值是 net.Dialer.DialContext
    net.Dialer{}.Timeout：TCP 连接的超时时间
    net.Dialer{}.KeepAlive：心跳的间隔时间（默认 15 秒）
    net.Dialer{}.Deadline：超时时间戳，超过了这个时间，连接会被强制关闭
  http.Transport{}.TLSHandshakeTimeout：TLS 握手超时时间
*/

func TestTimeout(t *testing.T) {
	// 详见：net/http/client.go:611
	http.DefaultClient.Timeout = 11 * time.Second

	for i := 0; i < 4; i++ {
		t.Logf("第 %d 次调用：", i)
		res, err := request()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(res)
	}
}

func request() (string, error) {
	resp, err := http.Post("http://127.0.0.1:2222/timeout", "application/json", nil)
	if err != nil {
		return "", err
	}

	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bs), err
}

type anonymousServer struct {
	http.Handler
}

func NewAnonymousServer(handlerFunc http.HandlerFunc) http.Handler {
	return &anonymousServer{Handler: handlerFunc}
}

func (s anonymousServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	s.Handler.ServeHTTP(resp, req)
}

func TestHTTPServer(t *testing.T) {
	handler := NewAnonymousServer(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte(`{"code": 0}`))
		if err != nil {
			t.Fatal(err)
		}
	})
	err := http.ListenAndServe("127.0.0.1:2222", handler)
	if err != nil {
		t.Fatal(err)
	}
}
