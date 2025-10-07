package main

import (
	"fmt"

	"github.com/holihaoa/task/base1"
)

const mainName string = "main"

var nums [5]int = [5]int{1, 1, 3, 2, 2}

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	onceNumber := base1.GetOnceNumber(nums) // 第一题 只出现一次的数字
	fmt.Println("getOnceNumber is ", onceNumber)
	ispalindrome := base1.IsPalindrome(0) // 第二题 回文数
	fmt.Println("IsPalindrome is ", ispalindrome)
	isvalide := base1.IsValid("(]") // 第三题 有效括号
	fmt.Println("IsValid is ", isvalide)
	longestCommon := base1.LongestCommonPrefix([]string{"aaa", "aa", "aaa"}) // 第四题 公共前缀
	fmt.Println("LongestCommonPrefix is ", longestCommon)
	digits := base1.PlusOne([]int{0}) // 第四题 公共前缀
	fmt.Println("PlusOne result is ", digits)

}
