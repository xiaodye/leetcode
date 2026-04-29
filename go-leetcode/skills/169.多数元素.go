package leetcode

func majorityElement(nums []int) int {
	// 统计每个数字出现的次数
	freq := make(map[int]int)

	for _, num := range nums {
		freq[num]++
	}

	// 遍历map, 找出count > arr.length/2 的数字
	for num, count := range freq {
		if count > len(nums)/2 {
			return num
		}
	}

	return -1
}
