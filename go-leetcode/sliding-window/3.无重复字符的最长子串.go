package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	// 滑动窗口
	i := 0
	j := 0
	maxLen := 0
	set := make(map[byte]bool)

	// 将字符串转为字节切片（适用于 Ascii 字符，若含 Unicode 请改用 []rune）
	bytes := []byte(s)

	for j < len(bytes) {
		for j < len(bytes) && !set[bytes[j]] {
			set[bytes[j]] = true
			if len(set) > maxLen {
				maxLen = len(set)
			}
			j++
		}

		// 收缩窗口
		delete(set, bytes[i])
		i++
	}

	return maxLen
}
