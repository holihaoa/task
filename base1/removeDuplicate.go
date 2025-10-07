package base1

func RemoveDuplicates(nums []int) int {
	//不知道为什么力扣不通过
	/*contrastMap := make(map[int]int)
	result := []int{}
	for i, num := range nums {
		if _, ok := contrastMap[num]; !ok {
			result = append(result, num)
			contrastMap[num] = i
		} else {
			continue
		}
	}
	nums = result
	return len(nums)*/
	//通过的解法
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
