package _46lru_cache

// Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.
//
// Implement the LRUCache class:
//
//    LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
//    int get(int key) Return the value of the key if the key exists, otherwise return -1.
//    void put(int key, int value) Update the value of the key if the key exists. Otherwise, add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation, evict the least recently used key.
//
// The functions get and put must each run in O(1) average time complexity.

// Constraints:
//
//    1 <= capacity <= 3000
//    0 <= key <= 104
//    0 <= value <= 105
//    At most 2 * 105 calls will be made to get and put.

// 实现一个满足 最近最少使用 策略的缓存
// 为了最佳的性能，自己使用的是 双链表 + map 的数据结构实现的
// 实现的效果和最佳的参考答案差不多
// 但是完成这个题，自己补充细节，调试花了太多的时间

type Node struct {
	k    int
	v    int
	pre  *Node
	next *Node
}

type LRUCache struct {
	l    int
	c    map[int]*Node
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		l: capacity,
		c: make(map[int]*Node, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.c[key]
	if ok {
		// 将元素移到队首
		if v != c.head {
			if v.pre != nil {
				v.pre.next = v.next
			}
			if v.next != nil {
				v.next.pre = v.pre
			}
			if c.tail == v {
				c.tail = v.pre
			}
			v.next = c.head
			v.pre = nil
			c.head.pre = v
			c.head = v
		}
		return v.v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	// 访问指定 key 对应的元素
	v := c.Get(key)
	if v != -1 {
		c.c[key].v = value
		return
	}

	if len(c.c) == c.l {
		// - map 中删除队尾
		delete(c.c, c.tail.k)

		// - 队列 中删除队尾
		if c.head == c.tail {
			c.head = nil
			c.tail = nil
		} else {
			c.tail.pre.next = nil
			c.tail = c.tail.pre
		}
	}
	// 插入队首
	n := &Node{
		k:    key,
		v:    value,
		pre:  nil,
		next: c.head,
	}
	// - map 中新增
	c.c[key] = n
	// - 队列 中新增
	if c.head == nil {
		c.tail = n
	} else {
		c.head.pre = n
	}
	c.head = n
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
