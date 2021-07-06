package _prototype

import (
	"encoding/json"
	"time"
)

/*
一、用自己的话，简单介绍一下：
利用已有的对象实例（这个就是原型）复制创建一个新对象实例的方式，来达到节省创建时间的目的

二、这里提一下适用场景
成本大。对象数据需要经过复杂的计算，排序 hash，或是 从 rpc、网络、数据库等慢io中获取

虽然是这个道理，但是并没有什么需要复制对象的使用场景（这可能和个人经历有关，因为 Java 和 C++ 这种面向对象的语言，这个模式确实用的少，javascript 用的就比较多了）

三、拓展一下，深拷贝 和 浅拷贝 其实就是从这里的概念引申出来的

下例的业务场景：
    有一部分数据，会在启动时加载到内存中，并且需要定时更新里面的数据
    同时展示给用户的数据每次必须要是相同版本的数据，不能一部分数据来自版本 1 一部分来自版本 2
    和加锁没关系！
*/

type Keyword struct {
	word      string
	visit     int
	UpdatedAt *time.Time
}

func (k *Keyword) Clone() *Keyword {
	// 这里使用序列化与反序列化的方式实现深拷贝
	var newKeyword Keyword

	b, _ := json.Marshal(k)
	_ = json.Unmarshal(b, &newKeyword)
	return &newKeyword
}

type Keywords map[string]*Keyword

func (words Keywords) Clone(updatedWords []*Keyword) Keywords {
	newKeywords := Keywords{}

	// 浅拷贝：直接拷贝了地址
	for k, v := range words {
		newKeywords[k] = v
	}

	// 深拷贝：替换掉需要更新的字段
	for _, word := range updatedWords {
		newKeywords[word.word] = word.Clone()
	}

	return newKeywords
}
