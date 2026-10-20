package leetcode

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pairs := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []byte{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, ok := pairs[c]; ok {
			// 是左括号，进栈
			stack = append(stack, c)
		} else {
			// 是右括号，栈空或栈顶不匹配则非法
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 出栈
			if pairs[top] != c {
				return false
			}
		}
	}

	return len(stack) == 0
}
