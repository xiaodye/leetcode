package leetcode

func partition(s string) [][]string {
	res := [][]string{}
	path := []string{}

	var backtrack func(startIndex int)
	backtrack = func(startIndex int) {
		if startIndex == len(s) {
			tmp := make([]string, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := startIndex; i < len(s); i++ {
			// startIndex - i 是否为回文
			if !isPalindrome(s, startIndex, i) {
				continue
			}
			subStr := s[startIndex : i+1]
			path = append(path, subStr)
			backtrack(i + 1)
			// 回溯
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return res
}

func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
