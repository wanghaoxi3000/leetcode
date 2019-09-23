package string

/*
20. Valid Parentheses 有效的括号

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：
左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:
输入: "()"
输出: true

示例 2:
输入: "()[]{}"
输出: true

示例 3:
输入: "(]"
输出: false

示例 4:
输入: "([)]"
输出: false

示例 5:
输入: "{[]}"
输出: true
*/

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	pare := make([]rune, len(s))
	top := -1
	for _, c := range s {
		if c == '{' || c == '[' || c == '(' {
			top++
			pare[top] = c
		} else {
			if (top < 0) || (c == '}' && pare[top] != '{') || (c == ']' && pare[top] != '[') || (c == ')' && pare[top] != '(') {
				return false
			}
			top--
		}
	}

	if top >= 0 {
		return false
	}

	return true
}
