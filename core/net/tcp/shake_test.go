package tcp

import (
	"testing"
)

/*
为了了解清楚 tcp 建立连接 和 断开连接 的过程，对实际的过程进行抓包

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
*/

func TestShake_Server(t *testing.T) {
	c, err := server()
	if err != nil {
		t.Fatal(err)
	}
	// 确保断开连接流程，以完成四次挥手
	defer MustClose(c)

	// 添加消息读取，以达到预期的抓包效果
	bs := make([]byte, 1)
	_, err = c.Read(bs)
	if err != nil {
		t.Fatal(ErrWrap(err))
	}
}

func TestShare_Client(t *testing.T) {
	c, err := client()
	if err != nil {
		t.Fatal(err)
	}
	// 确保断开连接流程，以完成四次挥手
	defer MustClose(c)
}
