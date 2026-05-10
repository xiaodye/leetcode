package leetcode

func productExceptSelf(nums []int) []int {
	// 前缀和
	// prefix[i]: i 左侧的数组乘积和
	// postfix[i]: i 右侧的数组的乘积和
	// answer[i] = prefix[i] * postfix[i]
	prefix := make([]int, len(nums))
	postfix := make([]int, len(nums))
	answer := make([]int, len(nums))

	prefix[0] = 1
	postfix[len(nums)-1] = 1

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := len(nums) - 2; i >= 0; i-- {
		postfix[i] = postfix[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = prefix[i] * postfix[i]
	}

	return answer
}
