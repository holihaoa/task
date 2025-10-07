package main

import (
	"fmt"

	"github.com/holihaoa/task/base1"
	_ "github.com/holihaoa/task/base1"
)

const mainName string = "main"

var mainVar string = getMainVar()

var nums [5]int = [5]int{1, 1, 3, 2, 2}

func init() {
	fmt.Println("main init method invoked")
}

func main() {
	onceNumber := base1.GetOnceNumber(nums)
	fmt.Println("getOnceNumber is ", onceNumber)
	ispalindrome := base1.IsPalindrome(0)
	fmt.Println("IsPalindrome is ", ispalindrome)
	isvalide := base1.IsValid("(]")
	fmt.Println("IsValid is ", isvalide)
	longestCommon := base1.LongestCommonPrefix([]string{"aaa", "aa", "aaa"})
	fmt.Println("LongestCommonPrefix is ", longestCommon)
}

func getMainVar() string {
	fmt.Println("main.getMainVar method invoked!")
	return mainName
}
