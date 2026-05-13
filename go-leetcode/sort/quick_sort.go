package leetcode

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	// 把第一个元素作为基准元素
	pivot := arr[0]

	leftArr := []int{}
	rightArr := []int{}

	// 对元素进行分堆，小于基准的放左边，大于的放右边
	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			leftArr = append(leftArr, arr[i])
		} else {
			rightArr = append(rightArr, arr[i])
		}
	}

	// 以此递归
	leftSorted := quickSort(leftArr)
	rightSorted := quickSort(rightArr)

	result := append(leftSorted, pivot)
	result = append(result, rightSorted...)

	return result
}
