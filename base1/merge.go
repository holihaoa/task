package base1

func Merge(intervals [][]int) [][]int {
	for {
		length := len(intervals)
		for i := len(intervals) - 1; i >= 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if intervals[i][1] >= intervals[j][0] && intervals[i][0] <= intervals[j][0] {
					intervals[j][0] = intervals[i][0]
					if intervals[i][1] >= intervals[j][1] {
						intervals[j][1] = intervals[i][1]
					}
					intervals = append(intervals[:i], intervals[i+1:]...)
					i--
				} else if intervals[j][1] >= intervals[i][0] && intervals[j][0] <= intervals[i][0] {
					if intervals[j][1] < intervals[i][1] {
						intervals[j][1] = intervals[i][1]
					}
					intervals = append(intervals[:i], intervals[i+1:]...)
					i--
				}
			}
		}
		if length == len(intervals) {
			break
		}
	}
	return intervals

	//更好的解法
	/*slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	ans := [][]int{}
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return ans*/
}
