package leetcode

func canJump(nums []int) bool {
	maxJump := 0

	for i := 0; i < len(nums); i++ {
		if i > maxJump {
			return false
		}

		// 我们依次遍历数组中的每一个位置，并实时维护 最远可以到达的位置
		if i+nums[i] > maxJump {
			maxJump = i + nums[i]
		}

		// 最远可以到达的位置 大于等于数组中的最后一个位置
		if maxJump >= len(nums)-1 {
			return true
		}
	}

	return false
}
