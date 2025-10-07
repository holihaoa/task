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
	digits := base1.PlusOne([]int{0}) // 第五题 加一
	fmt.Println("PlusOne result is ", digits)
	num := base1.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}) // 第六题 去重
	fmt.Println("RemoveDuplicates num is ", num)
	mergeArr := base1.Merge([][]int{{4, 5}, {1, 4}, {0, 1}}) // 第七题 合并区间
	fmt.Println("mergeArr is ", mergeArr)
}
