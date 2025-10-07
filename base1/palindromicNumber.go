package base1

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	str := strconv.Itoa(x)
	strs := []byte(str)
	for i := 0; i < len(strs)/2; i++ {
		if strs[i] != strs[len(strs)-i-1] {
			return false
		}
	}
	return true
}
