package tree_node

import (
	"encoding/json"
	"testing"
)

func p(i int) *int { return &i }

func TestFromArr(t *testing.T) {
	tree := FromArr(p(1), p(2), p(3))
	bs, _ := json.Marshal(tree)
	t.Log(string(bs))
}

func TestToArr(t *testing.T) {
	tree := FromArr(p(1), p(2), p(3))
	arr := tree.ToArr()
	bs, _ := json.Marshal(arr)
	t.Log(string(bs))
}

// TestComposite
// 看下是否符合 https://leetcode.com/problems/flatten-binary-tree-to-linked-list/ 中给出的遍历树的结果
func TestComposite(t *testing.T) {
	arr := []*int{p(1), p(2), p(3), p(4), p(5), p(6)}
	tree := new(TreeNode)
	p := tree
	for _, v := range arr {
		p.Right = &TreeNode{Val: *v}
		p = p.Right
	}
	res := tree.Right.ToArr()
	t.Log(res)
}
