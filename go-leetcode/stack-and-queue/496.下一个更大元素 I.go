package leetcode

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := make([]int, len(nums1))
	for i := range ans {
		ans[i] = -1
	}
	stack := []int{}
	m := make(map[int]int)

	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > nums2[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m[nums2[top]] = nums2[i]
		}
		stack = append(stack, i)
	}

	for i, num := range nums1 {
		if val, ok := m[num]; ok {
			ans[i] = val
		}
	}

	return ans
}
