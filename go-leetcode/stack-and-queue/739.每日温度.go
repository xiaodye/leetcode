package leetcode

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)

	// 单调递减栈，栈存储索引
	stack := []int{}

	// 初始化 0
	res := make([]int, n)

	for i := 0; i < n; i++ {
		// 维护一个单调递减栈
		for len(stack) != 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] = i - top
		}

		// 进栈
		stack = append(stack, i)
	}

	return res
}
