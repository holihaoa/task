package base1

import (
	"fmt"
)

const PkgName string = "base1"

// 自己的解法
func GetOnceNumber(nums [5]int) int {
	var count map[int]int = make(map[int]int)
	fmt.Println(nums)
	for index := range nums {
		a := nums[index]
		value := count[a]
		count[a] = value + 1
	}
	for index, value := range count {
		if value == 1 {
			return index
		}
	}
	return -1
}

// 官方的解法
func singleNumber(nums []int) int {
	single := 0
	for _, num := range nums {
		single ^= num
	}
	return single
}
