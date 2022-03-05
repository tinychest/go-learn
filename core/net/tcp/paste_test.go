package tcp

import (
	"testing"
)

/*
[现象]
可能会观察到控制台输出这样的日志：
等待客户端消息...
读取到客户端消息： I'm Client!I'm Client!
等待客户端消息...
读取到客户端消息： I'm Client!
等待客户端消息...
连接断开（正常）：EOF

会发现，第一次，一下读到了两条消息。我们将这种消息粘连在一起，无法区分的问题称作 “粘包”

为什么会出现这种问题，为什么每次服务端读取的消息，不是实际客户端每次发送的消息？

[客户端原因]
是因为发送端为了更加高效的将信息发送给接收端，TCP 默认采用了 Nagle 算法，该算法的机理是
将多次间隔较小、数据量较小的数据合并成一个大的数据块，封包进行发送，以减少网络中报文段的数量

解决：关闭 Nagle 算法，通过设置 TCP_NODELAY 为 true（在 Windows 中，需要修改注册表）
发送的时候，加上间隔停顿实践也可以从现象上解决这个问题

[服务端原因]
TCP 接收到数据包，并不会马上提交到应用层，也可以说应用层不会立即处理。
TCP 会将接收到的数据包保存在接收缓存里，应用程序主动从缓存读取，假如应用程序如果读取慢了，就有可能一下读到多个首尾相连的包。

解决：只能从应用层解决，如果从这个角度出发来解决，实际上是可以完全解决粘包问题的（客户端不用做任何处理，优化算法 Nagle 依旧开着）
其实就是约定好数据交互的规则：
- 为发送的数据添加指定的开始符和结束符，这样服务端就可以通过辨别特殊符号，将数据包区分开来（但是需要考虑数据内部不包含约定的特殊符号）
- 还有一种比较常见的做法，就是约定在发送的数据开头添加固定的字节数来表示数据的长度

[拓展]
UDP 是面向消息的，存在消息保护边界，即，接收方一次只接收一条独立的消息，因此不会有粘包问题。
举例
*/

func TestPaste_Server(t *testing.T) {
	c, err := server()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// 读取客户端消息
	var bs = make([]byte, 100)
	for {
		t.Log("等待客户端消息...")
		n, err := c.Read(bs)
		if err != nil {
			t.Fatal(ErrWrap(err))
		}
		t.Log("读取到客户端消息：", string(bs[:n]))
	}
}

func TestPaste_Client(t *testing.T) {
	c, err := client()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// 向服务端发送消息
	const write = "I'm Client!"
	for i := 0; i < 3; i++ {
		_, err = c.Write([]byte(write))
		if err != nil {
			t.Fatal(err)
		}
		t.Log("成功向服务端发送：", write)
	}
}
