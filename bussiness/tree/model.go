package tree

// 一颗树的所有信息记录
type One struct {
	Level      int       // 树的高度
	Nodes      []*Node   // 所有节点
	Root       *Node     // 根
	LevelNodes [][]*Node // 二维数组，一维的含义是 节点所在的层次，顶层为0，二维是该层所有的节点
}

type Node struct {
	Id       int64 // 查询
	Pid      int64 // 查询
	Level    int   // 查询
	Children []*Node
	Marked   bool
}

// 节点列表数据需要由调用者来保证如下几点
// 1.节点列表不为空
// 2.节点列表的第一个元素是根
func NewOne(nodes []*Node, maxLevel int) *One {
	// Nodes
	one := &One{Nodes: nodes, Level: maxLevel, Root: nodes[0]}

	// LevelNodes init
	one.LevelNodes = make([][]*Node, maxLevel+1)
	for i := 0; i <= maxLevel; i++ {
		one.LevelNodes[i] = make([]*Node, 0)
	}

	// pidMap init
	pidMap := make(map[int64][]*Node, len(nodes))
	for _, node := range nodes {
		// LevelNodes
		one.LevelNodes[node.Level] = append(one.LevelNodes[node.Level], node)
		// pidMap
		if _, ok := pidMap[node.Pid]; !ok {
			pidMap[node.Pid] = make([]*Node, 0)
		}
		pidMap[node.Pid] = append(pidMap[node.Pid], node)
	}

	// tree
	for _, node := range nodes {
		node.Children = pidMap[node.Id]
	}

	return one
}
