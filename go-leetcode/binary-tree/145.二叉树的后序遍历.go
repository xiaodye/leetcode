package leetcode

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var postorder func(node *TreeNode)

	postorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		postorder(node.Left)
		postorder(node.Right)
		res = append(res, node.Val)
	}

	postorder(root)

	return res
}
