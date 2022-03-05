package tcp

import "testing"

// 收到一条回复一条，直至发生任何错误（包括连接断开）
func TestCommunicate_Server(t *testing.T) {
	c, err := server()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// 与客户端交互
	const write = "I'm Server!"

	var bs = make([]byte, 100)
	for {
		// 读
		t.Log("等待客户端消息...")
		n, err := c.Read(bs)
		if err != nil {
			t.Fatal(ErrWrap(err))
		}
		t.Log("读取到客户端消息：", string(bs[:n]))

		// 写
		_, err = c.Write([]byte(write))
		if err = ErrWrap(err); err != nil {
			t.Log(ErrWrap(err))
			return
		}
		t.Log("回复客户端消息：", write)
	}
}

// 向服务端发送消息，每发送一条，就等待服务端响应，然后再发
func TestCommunicate_Client(t *testing.T) {
	c, err := client()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// 与服务端交互
	const write = "I'm Client!"

	var bs = make([]byte, 100)

	for i := 0; i < 3; i++ {
		// 写
		_, err = c.Write([]byte(write))
		if err != nil {
			t.Fatal(err)
		}
		t.Log("成功向服务端发送：", write)

		// 读
		t.Log("等待服务端响应...")
		n, err := c.Read(bs)
		if err = ErrWrap(err); err != nil {
			t.Fatal()
		}
		t.Log("读取到服务端消息：", string(bs[:n]))
	}
}