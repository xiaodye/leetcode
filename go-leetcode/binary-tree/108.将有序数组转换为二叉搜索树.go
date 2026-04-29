package leetcode

func sortedArrayToBST(nums []int) *TreeNode {
	// 边界：空数组返回 nil
	if len(nums) == 0 {
		return nil
	}

	var buildBST func(left, right int) *TreeNode
	buildBST = func(left, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := left + (right-left)/2
		node := &TreeNode{Val: nums[mid]}
		node.Left = buildBST(left, mid-1)
		node.Right = buildBST(mid+1, right)
		return node
	}

	return buildBST(0, len(nums)-1)
}
