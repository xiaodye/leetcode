package leetcode

import "strings"

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}

	currNum := 0
	currStr := ""

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c >= '0' && c <= '9' {
			// 外层数字情况，可能大于 10，需要累乘
			currNum = currNum*10 + int(c-'0')
		} else if c == '[' {
			// 进栈
			numStack = append(numStack, currNum)
			strStack = append(strStack, currStr)
			currNum = 0
			currStr = ""
		} else if c == ']' {
			// 出栈
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]

			prev := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]

			currStr = prev + strings.Repeat(currStr, count)
		} else {
			// 字母，累加
			currStr += string(c)
		}
	}

	return currStr
}
