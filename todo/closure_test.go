package todo

import (
	"testing"
)

type p struct {
	value    int
	children []*p
}

func traverse(node *p, do func(indicator *p)) {
	do(node)
	for _, v := range node.children {
		traverse(v, do)
	}
}

func TestTraverse(t *testing.T) {
	tree := p{
		value: 1,
		children: []*p{
			{
				value: 11,
				children: []*p{
					{value: 111},
				},
			},
			{
				value: 12,
				children: []*p{
					{value: 121},
				},
			},
			{
				value: 13,
				children: []*p{
					{value: 131},
				},
			},
		},
	}

	for _, v := range tree.children {
		traverse(v, func(node *p) {
			// 虽然 traverse 会进行递归调用，但是当前作为参数的方法的 v 的值就是闭包外的值，不会变动
			// 可以理解为，traverse 的两个参数中引用的 v 都是此时 for 上的 v
			// 虽然不知道为什么会有这种误区，但是解释一下，这里会有因为引用的变量导致最后的操作都是同一个值，这里不会吗，答：这里没有 defer...
			// 每次循环执行完，才会 改变 v 的值，进入下一次循环
			node.value = v.value
		})
	}
	t.Log(1)
}
