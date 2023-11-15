package main

import "fmt"

/*
*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。
*/
func main() {

	fmt.Println(isValid("()"))

}

func isValid(s string) bool {
	//是否是奇数

	if len(s)%2 == 1 {
		return false
	}
	var stack []byte

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	for i := 0; i < len(s); i++ {
		if v, ok := pairs[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != v {
				return false
			}

			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}

	}
	return len(stack) == 0
}
