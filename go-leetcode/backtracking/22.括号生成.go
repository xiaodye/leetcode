package leetcode

func generateParenthesis(n int) []string {
	res := []string{}

	var backtrack func(left int, right int, str string)
	backtrack = func(left int, right int, str string) {
		// 边界条件
		//  1. 左括号数小于 右，直接结束
		//  2. 括号数量等于 n, push 结果，结束
		if left < right {
			return
		}
		if left+right == 2*n {
			res = append(res, str)
			return
		}
		if left < n {
			backtrack(left+1, right, str+"(")
		}
		if right < n {
			backtrack(left, right+1, str+")")
		}
	}

	backtrack(0, 0, "")
	return res
}
