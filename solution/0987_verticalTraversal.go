package solution

import "sort"

func VerticalTraversal(root *TreeNode) [][]int {
	// 遍历出所有节点的坐标，使用层次遍历
	// 创建一个队列，用于存储节点以及它们所在的层次
	var queue []struct {
		node     *TreeNode
		level    int
		vertical int
	}
	var traversalRes []struct {
		node     *TreeNode
		level    int
		vertical int
	}
	// 将root压入队列
	queue = append(queue, struct {
		node     *TreeNode
		level    int
		vertical int
	}{node: root, level: 0, vertical: 0})

	for len(queue) != 0 {
		// 从队列中取出一个节点
		top := queue[0]
		queue = queue[1:]
		// 将坐标结果压入结果集
		traversalRes = append(traversalRes, top)
		// 如果该节点有左子节点，则将左子节点入队。
		if top.node.Left != nil {
			queue = append(queue, struct {
				node     *TreeNode
				level    int
				vertical int
			}{node: top.node.Left, level: top.level + 1, vertical: top.vertical - 1})
		}
		if top.node.Right != nil {
			queue = append(queue, struct {
				node     *TreeNode
				level    int
				vertical int
			}{node: top.node.Right, level: top.level + 1, vertical: top.vertical + 1})
		}
	}
	// 接下来遍历坐标的结果集，将对应的列加入垂直哈希表
	// 这里需要记录值、横坐标在mapping中
	var verticalMapping = make(map[int][]struct {
		value   int
		horizon int
	})
	for _, v := range traversalRes {
		// 如果没有实例化切片，则先实例化切片
		if verticalMapping[v.vertical] == nil {
			verticalMapping[v.vertical] = make([]struct {
				value   int
				horizon int
			}, 0)
		}
		verticalMapping[v.vertical] = append(verticalMapping[v.vertical], struct {
			value   int
			horizon int
		}{value: v.node.Val, horizon: v.level})
	}
	// 接下来遍历哈希表，对垂直列进行排序
	// 哈希表排序数组
	var mapSort []int
	for i, _ := range verticalMapping {
		mapSort = append(mapSort, i)
	}
	// 使用排序库
	sort.Ints(mapSort)
	// 对verticalMapping进行排序
	for _, col := range mapSort {
		/*
			 * 在sort.Slice的比较函数中，i和j代表要比较的两个元素在切片中的索引。
				这两个索引用于访问切片中的元素，以便比较它们的顺序。
				比较函数需要根据这两个索引来决定对应元素的排序关系。
			 * i索引：指向切片中当前正在考虑进行排序操作的一个元素。
			 * j索引：指向切片中另一个用于比较的元素。
			 * 比较函数的逻辑是根据这两个元素的属性（在您的情况下是horizon和value）来决定它们的排序顺序。
			 * 如果您希望索引i指向的元素排在索引j指向的元素之前，比较函数应该返回true；否则，返回false。
		*/
		sort.Slice(verticalMapping[col], func(i, j int) bool {
			// 先比较行位置
			if verticalMapping[col][i].horizon != verticalMapping[col][j].horizon {
				return verticalMapping[col][i].horizon < verticalMapping[col][j].horizon
			}
			// 行位置相同，则比较节点值
			return verticalMapping[col][i].value < verticalMapping[col][j].value
		})
	}
	// 输出
	var res [][]int
	for i := 0; i < len(mapSort); i++ {
		var saving []int
		for j := 0; j < len(verticalMapping[mapSort[i]]); j++ {
			saving = append(saving, verticalMapping[mapSort[i]][j].value)
		}
		res = append(res, saving)
	}
	return res
}
