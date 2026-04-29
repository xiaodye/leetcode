package leetcode

func longestConsecutive(nums []int) int {
	max := 0

	// 把所有节点都存储到 set 中，并进行去重
	set := make(map[int]bool)
	for _, num := range nums {
		set[num] = true
	}
	// 将 set 中的 key 转为切片
	unique := make([]int, 0, len(set))
	for num := range set {
		unique = append(unique, num)
	}
	nums = unique

	// 全部放 set 里
	for i := 0; i < len(nums); i++ {
		// 如果该节点存在前驱节点，说明这个节点开始的序列肯定不是最长的
		if set[nums[i]-1] {
			continue
		}

		length := 1
		num := nums[i]

		for set[num+1] {
			length++
			num++
		}

		if length > max {
			max = length
		}
	}

	return max
}
