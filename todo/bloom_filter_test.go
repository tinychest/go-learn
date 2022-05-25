package todo

import "testing"

// 布隆过滤器（Bloom Filter）是 1970 年由布隆提出的。它实际上是一个很长的二进制向量和一系列随机映射函数。主要用于判断一个元素是否在一个集合中。
//
// 通常我们会遇到很多要判断一个元素是否在某个集合中的业务场景，一般想到的是将集合中所有元素保存起来，然后通过比较确定。链表、树、散列表（又叫哈希表，Hash table）等等数据结构都是这种思路。
// 但是随着集合中元素的增加，我们需要的存储空间也会呈现线性增长，最终达到瓶颈。同时检索速度也越来越慢，上述三种结构的检索时间复杂度分别为$O(n)$，$O(logn)$，$O(1)$。
//
// 这个时候，布隆过滤器（Bloom Filter）就应运而生。
//
// [理解]
// 用自己的话说，拿 map 来说，每个 key 经由哈希算法最终映射到一个地址值，拿到对应的值；
// Bloom 在一段有限的 bit 数组空间中，是一个有着多个 hash 映射算法的，每个 key 经由这些映射规则会对应到数组空间中多个 bit 位并置为 1，
// 假如，我们获取一个不存在的 key，但是有可能对应的多个 bit 位都恰好因为其他 key 覆盖成了 1，那么布隆过滤器会告诉你，而实际是没有，所以这就是“有不一定有”
// 反之，如果说没有，那就是真的没有。
// 分析可得：
// - 布隆过滤器的实现，占用空间小，能支持非常大的数据量集
// - 布隆过滤器说有是不准确的，是否准确的概率，是受到 数组空间大小 和 映射规则算法 影响的

func TestBloomFilter(t *testing.T) {
	// 实现先放一放
}