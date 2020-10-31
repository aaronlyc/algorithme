package __replace_space

func replaceSpace(s string) string {
	spaceNums := spaceNums(s)
	if spaceNums < 1 { //如果没有空格, 则返回原来s
		return s
	}

	addNums, sLen := 2*spaceNums, len(s)
	l, sIndex := sLen+addNums, sLen-1
	j := l - 1
	result := make([]byte, l, l)
	for sIndex >= 0 {
		if s[sIndex] == ' ' {
			result[j], result[j-1], result[j-2] = '0', '2', '%'
			j = j - 3
		} else {
			result[j] = s[sIndex]
			j--
		}
		sIndex--
	}
	return string(result)
}

// 获取s中的空格数量
func spaceNums(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			num++
		}
	}
	return num
}
