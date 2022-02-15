package example3

// Bidirectional stream RPC
// 从调整 proto 中接口定义开始，参数 和 返回值前都添加 stream
// - 生成的客户端有 Send 方法，服务端有 Recv 方法
// - 客户端 和 服务端 都只能通过各自 Stream 的包装实例调用 发送、接收 等相关方法