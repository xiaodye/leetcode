package leetcode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			// 弹出队首节点
			node := queue[0]
			queue = queue[1:]

			// 记录每层最后一个节点
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}
