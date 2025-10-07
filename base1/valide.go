package base1

import "strings"

func IsValid(s string) bool {
	//错误方法
	/*if len(s) < 2 {
		return false
	}
	if string(s[0]) == ")" || string(s[0]) == "}" || string(s[0]) == "]" {
		return false
	}
	for i := 0; i < len(s)/2; i++ {
		if string(s[i]) == "(" && string(s[len(s)-i-1]) != ")" && string(s[i+1]) != ")" {
			return false
		}
		if string(s[i]) == "[" && string(s[len(s)-i-1]) != "]" && string(s[i+1]) != "]" {
			return false
		}
		if string(s[i]) == "{" && string(s[len(s)-i-1]) != "}" && string(s[i+1]) != "}" {
			return false
		}
	}
	return true*/
	//自己写的正确答案,复杂度较高
	for {
		if strings.Contains(s, "()") {
			s = strings.Replace(s, "()", "", 1)
		}
		if strings.Contains(s, "[]") {
			s = strings.Replace(s, "[]", "", 1)
		}
		if strings.Contains(s, "{}") {
			s = strings.Replace(s, "{}", "", 1)
		}
		if len(s) != 0 && (!strings.Contains(s, "{}") && !strings.Contains(s, "[]") && !strings.Contains(s, "()")) {
			return false
		}
		if len(s) == 0 {
			return true
		}
	}
	//官方答案
	/*list := make([]int, len(s))
	top := -1
	for _, i2 := range s {
		flag := getInt(i2)
		if flag%2 == 1 {
			top++
			list[top] = flag
		} else {
			if top >= 0 && list[top] == flag-1 {
				top--
			} else {
				return false
			}
		}
	}
	return top == -1*/
}

/*func getInt(x rune) int {
	switch x {
	case '(':
		return 1
	case ')':
		return 2
	case '[':
		return 3
	case ']':
		return 4
	case '{':
		return 5
	case '}':
		return 6
	default:
		return 0
	}
}*/
