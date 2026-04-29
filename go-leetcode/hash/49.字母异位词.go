package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	// key: 排序后的字符串，value: []，收集的字符串数组
	m := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// 字符排序
		s := []byte(strs[i])
		sort.Slice(s, func(a, b int) bool { return s[a] < s[b] })
		key := string(s)

		// if _, ok := m[key]; ok {
		// 	m[key] = append(m[key], strs[i])
		// } else {
		// 	m[key] = []string{strs[i]}
		// }

		// 直接 append，无需判断 key 是否存在
		m[key] = append(m[key], strs[i])
	}

	// 返回 map 的所有 values 作为二维切片
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
