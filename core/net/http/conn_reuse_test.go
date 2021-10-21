package http

// 来源：https://mp.weixin.qq.com/s/PjSM2RNpIA934KBPDnf6MA
// 文章目的是，教你注意，基于 tcp 的 http 协议，Client 要复用（连接复用）
// 也就是需要注意超时时间不要设置错了，不要给 http.Client → http.Transport → net.Conn 设置了超时，应该直接给 Client 设置

// 但是，在了解完之后，突然“开瓢”，自己更关注复用连接，来提高内部服务调用的效率 - 了解了原理
// 查看源码：实际上，http.Get http.Post 等方法本质上，就是通过 http.DefaultClient 进行调用的，也就是默认具有连接复用的特性
// 找寻相关资料：https://ethantang.top/posts/go-http-keepalive/
