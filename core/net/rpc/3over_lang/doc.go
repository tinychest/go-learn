package crpc

/*
实现跨语言的 RPC，即，数据传输协议采用 json

实际测试，没有效果？
通过原生的处理方式替代 rpc 服务端接收客户端发送的内容失败，不按照既定的数据交互逻辑，客户端一致发送真实数据
只能通过抓包：Wireshark → Adapter for loopback traffic capture → tcp.port == 1234（过滤条件）
*/
