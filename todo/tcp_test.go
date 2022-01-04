package todo

import (
	"errors"
	"fmt"
	"io"
	"net"
	"testing"
)

/*
【坑】
- io.ReadAll 读取直到 io.EOF，而只有客户端调用 Close，服务端才会收到 io.EOF → 不能使用 io.ReadAll 来读取消息
- 自定义消息读取，字节数组 len 有多大，一次最多就读多少 → 需要自定义结束标识位
（别想偷懒，利用 io.ReadAll，即 基于 io.EOF，要清楚规范中，将 io.EOF 错误定义为读 “尽” 了，而连接没有断开，怎么能认为读尽了）

【WireShark 抓包】
Adapter for loopback traffic capture → tcp.port == 4444
 → Run TestTCPServer
 → Run TestTCPClient

[建立连接 三次握手]
（参考 https://mp.weixin.qq.com/s/NL7Jzh0lYoA395yzaGxBHw 有动画）
握三次；双方都能明确自己和对方的收、发能力是正常的。不能少么？少了，根本不能保证连接建立。不能多么？多了，在网络环境不好的情况下，如果三次都不能保证连接建立，那四次也不行，且浪费资源。

 → 捕获到的前 3 条：
   421	18.777198	127.0.0.1	127.0.0.1	TCP	56	52978 → 4444 [SYN] Seq=0 Win=65535 Len=0 MSS=65495 WS=256 SACK_PERM=1
    客户端发送 syn，状态变为 syn_sent；服务端收到，状态变为 syn_rcvd
   422	18.777252	127.0.0.1	127.0.0.1	TCP	56	4444 → 52978 [SYN, ACK] Seq=0 Ack=1 Win=65535 Len=0 MSS=65495 WS=256 SACK_PERM=1
    服务端发送 ack + syn；客户端收到，状态变为 established
   423	18.777289	127.0.0.1	127.0.0.1	TCP	44	52978 → 4444 [ACK] Seq=1 Ack=1 Win=2619648 Len=0
    客户端发送 syn；服务端收到，状态 established

【TCP 概念补充】
- [重传] 一方向另外一方发送消息，如果超过一定时间都没有收到另外一方的 “确认收到” 的回复，就会认为发送的消息丢失了，会进行重发。
- [去重] 所以，既然存在重传，也就会存在另外一方收到了两次一模一样的消息。
操作系统的网络内核将上面的两种现象都处理好了，用户层是不用关心的

- [优化策略] 如果一方连续发送了很多消息，另一方并不需要一直回复收到了，而是在收到了连续的消息后，一次性回复，我都收到了
- [滑动窗口] 发送方的发送速率和接收方的接收速率是需要双方沟通得到的，这就是滑动窗口

[断开连接 四次挥手]
（挥手参照这个去看 https://mp.weixin.qq.com/s/jTDU-zxP1INTYLpGLypjXQ 更好）
挥四次，那么多？因为 TCP 是全双工协议，也就是双方都要关闭，一方都要向另一方发送 FIN 和 回应的 ACK。

 → 倒数的 4 条（最开始服务端忘记关闭连接了，导致 WireShark 抓包还暴红了）
   267	13.671516	127.0.0.1	127.0.0.1	TCP	44	56280 → 4444 [FIN, ACK] Seq=1 Ack=1 Win=2619648 Len=0
    服务端发送 fin，状态变为 fin_wait_1；服务端收到，状态变为 close_wait
    为什么还有 ACK，这个不需要太在意，

   268	13.671532	127.0.0.1	127.0.0.1	TCP	44	4444 → 56280 [ACK] Seq=1 Ack=2 Win=2619648 Len=0
    服务端发送 ack，状态不发生改变；客户端收到，状态变为 fin_wait_2
   269	13.671568	127.0.0.1	127.0.0.1	TCP	44	4444 → 56280 [FIN, ACK] Seq=1 Ack=2 Win=2619648 Len=0
    服务端发送 fin，状态变为 last_ack；客户端收到，状态变为 time_wait
	（有时候这两步，会直接合并，这个时候 四次挥手 就变成了 三次）

   270	13.671587	127.0.0.1	127.0.0.1	TCP	44	56280 → 4444 [ACK] Seq=2 Ack=2 Win=2619648 Len=0
    客户端发送 ack；服务端收到，状态变为 close
    等一会，客户端状态变为 close

【TCP 概念补充】
- [time_wait] 主动关闭的一方在回复完对方的挥手后进入的一个长期状态，这个状态标准的持续时间是 4 分钟（可以调整），4 分钟后才会进入到 closed 状态，释放套接字资源
    作用1 是，重传最后一个 ack 报文，确保对方可以收到，没有收到，会重传 fin 报文
    作用2 是，消化残留报文（防止新连接重复利用端口时，残留报文对新连接产生影响），这个状态下的关闭方会丢弃所有收到的报文，4 分钟的时间足以使得这些残留报文彻底消逝
- [默认 4 分钟] 4 分钟就是 2 个 MSL，每个 MSL 是 2 分钟。MSL 就是 maximium segment lifetime（最长报文寿命）。RFC 规定
*/

// 收到一条回复一条，直至发生任何错误（包括连接断开）
func TestTCPServer(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:4444")
	if err != nil {
		t.Fatal(err)
	}

	// 阻塞等待客户端连接
	t.Log("等待客户端连接...")
	c, err := l.Accept()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("一个客户端建立连接")
	// 连接后，确保断开，以完成四次挥手
	defer func() {
		if err := c.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// 与客户端交互
	// const write = "OK"
	//
	// var bs = make([]byte, 100)
	// for {
	// 	// 读
	// 	t.Log("等待客户端消息...")
	// 	n, err := c.Read(bs)
	// 	if err := ErrHandle(err); err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Logf("读取到客户端消息：%s", string(bs[:n]))
	//
	// 	// 写
	// 	_, err = c.Write([]byte(write))
	// 	if err = ErrHandle(err); err != nil {
	// 		t.Log(err)
	// 		return
	// 	}
	// 	t.Logf("回复客户端消息：%s", write)
	// }
}

// 向服务端发送消息，每发送一条，就等待服务端响应是否接收到了
// （这里的业务背景和 tcp 协议核心的通信设计理念一模一样，任意一方发送消息，必须收到对方回复确认收到，才会认为对方收到了）
func TestTCPClient(t *testing.T) {
	c, err := net.Dial("tcp", "127.0.0.1:4444")
	if err != nil {
		t.Fatal(err)
	}
	// 不调用 Close 方法，就是暴力断开
	defer func() {
		if err := c.Close(); err != nil {
			t.Fatal(err)
		}
	}()

	// 与服务端交互
	// const write = "Holy Shit!"
	//
	// var bs = make([]byte, 100)
	//
	// for i := 0; i < 3; i++ {
	// 	// 写
	// 	_, err = c.Write([]byte(write))
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Logf("成功向服务端发送：%s", write)
	//
	// 	// 读
	// 	t.Log("等待服务端响应...")
	// 	n, err := c.Read(bs)
	// 	if err := ErrHandle(err); err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Logf("读取到服务端消息：%s", string(bs[:n]))
	// }
}

func ErrHandle(err error) error {
	if errors.Is(err, net.ErrClosed) {
		return fmt.Errorf("连接断开（暴力）：%w", err)
	}
	if errors.Is(err, io.EOF) {
		return fmt.Errorf("连接断开（正常）：%w", err)
	}
	return err
}
