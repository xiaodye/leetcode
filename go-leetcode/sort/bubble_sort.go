package leetcode

func bubbleSort(arr []int) []int {
	n := len(arr)

	// 两层 for 循环，第一层代表轮次，比如 5 个数，每轮排好一个，只需要四轮，最后一个数自然就好了
	// 冒泡排序，其实是两两交换，从后往前排
	// 优化：如果发现一轮中，并没有发生交换，说明此时数组已经有序，不需要再排了
	// 时间复杂度 O(n * n)

	for i := 0; i < n-1; i++ {
		flag := false

		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}

		if !flag {
			return arr
		}
	}
	return arr
}
