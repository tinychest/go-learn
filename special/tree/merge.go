package tree

import "go-learn/tool"

// 调用者保证 ids 不为空
func (o *One) MergeAndValid(ids []int64) ([]int64, error) {
	var idList = tool.Int64Slice(ids)
	// mark
	for _, node := range o.Nodes {
		if idList.Contains(node.Id) {
			node.Marked = true
		}
	}

	var result = make([]int64, 0, len(ids))
	traverse(o.Root, &result)
	return result, nil
}

// 乐观：深度先序遍历
func traverse(node *Node, result *[]int64) {
	if node.Marked {
		*result = append(*result, node.Id)
		return
	}
	for _, child := range node.Children {
		traverse(child, result)
	}
}

// 悲观 - 校验做法
// func (o *One) MergeAndValid(ids []int64) ([]int64, error) {
//	var (
//		idList = help.Int64Slice(ids)
//		idMap = make(map[int64]struct{}, len(ids))
//	)
//	// mark
//	for _, node := range o.Nodes {
//		if idList.Contains(node.Id) {
//			node.Marked = true
//		}
//	}
//
//	var level = o.Level - 1
//	for levelNodes := o.LevelNodes[level]; level >= 0; level-- {
//		for _, node := range levelNodes {
//			if node.Marked {
//				for _, child := range node.Children {
//					idMap[child.Id] = struct{}{}
//					if child.Marked == false {
//						return nil, help.LogError("参数非法，缺少", child.Id)
//					}
//				}
//			}
//		}
//	}
//
//	var result = make([]int64, 0, len(ids))
//	for _, id := range ids {
//		if _, ok := idMap[id]; !ok {
//			result = append(result, id)
//		}
//	}
//	return result, nil
// }
