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
}
