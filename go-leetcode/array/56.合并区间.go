package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	result := [][]int{}

	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		if i == 0 {
			result = append(result, intervals[i])
		} else if result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
		} else {
			arr := result[len(result)-1]
			result = result[:len(result)-1] // pop
			newEnd := arr[1]
			if intervals[i][1] > newEnd {
				newEnd = intervals[i][1]
			}
			result = append(result, []int{arr[0], newEnd})
		}
	}

	return result
}
