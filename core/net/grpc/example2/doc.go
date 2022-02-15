package example2

// Server stream RPC（客户端多响应）
// 从调整 proto 中接口定义开始，返回值前添加 stream
// - 差异 客户端实现的调用方式
// - 差异 服务端的实现方式
// - 细节 生成的客户端没有，对应服务端 Recv 的 Send 方法
