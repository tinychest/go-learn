package decimal

// 因为计算机中二进制无法精确表示十进制小数，所以在进行相关计算，转化时，会出比较令你意外的结果
//
// 官方文档：http://docscn.studygolang.com/ref/spec#Comparison_operators
// 更专业的说法：只要浮点数遵循 IEEE 754 标准的编程语言都有这个问题
//
// 这里边想要重点强调两点：
// - 计算机中的二进制是计算机组成原理基垫，意味着只要你进行 10 进制的浮点数计算，就没有很好的方案来达到理想的效果
// - 既然无法避免，那就应该尽可能保证精度足够大，来将这个误差缩小，以贴近真实的结果
//
// 最佳实践：
// - 指定一个误差范围，两个浮点数的差值在此范围之内，则认为是相等的（阿里规范文档中提及，这不仅像是最佳实践，更像是一个计算机中的一个真理）
// - 不使用浮点数，使用整数。例：假定业务背景最小的金额单位是分，那么就应该将分作为系统最小的单位，以避免产生浮点数
//
// 有两篇讲的特别好的文章
// https://mp.weixin.qq.com/s?__biz=MzkyMDAzNjQxMg==&mid=2247484440&idx=1&sn=ed2e6bc81a6b40bf8bd1c2d6cf31d0af&scene=21#wechat_redirect
// https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651450239&idx=3&sn=9718870c742199724c229bacadde0349&scene=21#wechat_redirect
