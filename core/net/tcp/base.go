package tcp

import (
	"errors"
	"fmt"
	"io"
	"net"
)

// 为了方便、突出演示样例，封装建立连接的服务端、客户端、错误包装处理...
// 应该先运行 server，再运行 client

// 监听本地端口，等待客户端请求，建立连接并返回
func server() (net.Conn, error) {
	l, err := net.Listen("tcp", "127.0.0.1:4444")
	if err != nil {
		return nil, err
	}
	// 阻塞等待客户端连接
	return l.Accept()
}

// 向本地指定端口，发起请求，建立连接后并返回
func client() (net.Conn, error) {
	return net.Dial("tcp", "127.0.0.1:4444")
}

func ErrWrap(err error) error {
	if errors.Is(err, net.ErrClosed) {
		return fmt.Errorf("连接断开（暴力）：%w", err)
	}
	if errors.Is(err, io.EOF) {
		return fmt.Errorf("连接断开（正常）：%w", err)
	}
	return err
}

func MustClose(closer io.Closer) {
	err := closer.Close()
	if err != nil {
		panic(fmt.Errorf("关闭失败: %w", err))
	}
}
