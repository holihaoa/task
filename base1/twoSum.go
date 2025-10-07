package base1

func TwoSum(nums []int, target int) []int {
	result := make([]int, 0)
out:
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				break out
			}
		}
	}
	return result

	//更好的解法
	/*m := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if idx, exists := m[complement]; exists {
			return []int{idx, i}
		}
		m[num] = i
	}
	return []int{}*/
}
