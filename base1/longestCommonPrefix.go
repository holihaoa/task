package base1

func LongestCommonPrefix(strs []string) string {
	result := []rune(strs[0])
	for i := 1; i < len(strs); i++ {
		v := []rune(strs[i])
		if len(v) < len(result) {
			result = result[:len(v)]
		}
		if len(result) == 0 || len(v) == 0 {
			return ""
		}
		for i2 := 0; i2 < len(result); i2++ {
			if v[i2] == result[i2] {
				continue
			} else {
				if i2 == 0 {
					return ""
				}
				result = result[:i2]
				break
			}
		}
	}
	return string(result)
}
