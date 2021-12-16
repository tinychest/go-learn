package todo

import (
	"testing"
)

func TestIO(t *testing.T) {
	// 【缓存】
	// strings.Builder{}
	// bytes.Buffer{}

	// 【[]byte → io.Reader】
	// strings.Reader（io.Reader） strings.NewReader([]byte)
	// bytes.Reader（io.Reader） bytes.NewReader(string)

	// 【io.Reader → io.CloseReader】
	// io.CloseReader ioutil.NopCloser(io.Reader)

	// 【io.Reader → []byte】
	// []byte io.ReadAll(io.Reader)
	// []byte ioutil.ReadAll(io.Reader)

	// 【其他】
	// 类比 /dev/null，这里也有黑洞的概念
	// io.Copy(ioutil.Discard, resp.Body)
	// 如果你不需要用它，可以考虑丢弃它，例如 HTTP 客户端的传输不会重用连接时，直到 body 被读完才会关闭

	// ioutil 下还有 ReadFile、WriteFile 方法
}
