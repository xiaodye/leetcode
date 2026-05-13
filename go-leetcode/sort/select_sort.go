package leetcode

func selectSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	n := len(arr)

	// 选择排序，从前往后排，外层为轮数
	// 内层为比较次数

	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}
