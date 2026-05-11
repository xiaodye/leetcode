package leetcode

import (
	"sort"
	"strings"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	res := ""

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch >= '0' && ch <= '9' {
			// 外层数字情况，可能大于 10，需要累乘 123[a]
			num = num*10 + int(ch-'0')
		} else if ch == '[' {
			strStack = append(strStack, res)
			res = ""
			numStack = append(numStack, num)
			num = 0
		} else if ch == ']' {
			repeatTimes := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			res = prev + strings.Repeat(res, repeatTimes)
		} else {
			// 里层字符情况，可能存在多个
			res += string(ch)
		}
	}
	return res
}

func findKthLargest(nums []int, k int) int {
	sort.Slice(nums, func(i int, j int) bool {
		return nums[i] > nums[j]
	})

	return nums[k-1]
}
