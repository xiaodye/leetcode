package leetcode

func singleNumber(nums []int) int {
	res := 0

	// singleNumber 找出数组中只出现一次的数字（其他数字均出现两次）
	// 利用异或运算：a ^ a = 0，a ^ 0 = a，且满足交换律与结合律
	// 遍历数组，所有数字依次异或，成对数字抵消，剩下的即为只出现一次的数字
	for _, num := range nums {
		res ^= num
	}

	return res
}
