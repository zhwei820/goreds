package goreds

// 是否含有汉字
func ContainsCh(str string) bool {
	r := []rune(str)
	for i := 0; i < len(r); i++ {
		if r[i] <= 40869 && r[i] >= 19968 {
			return true
		}
	}
	return false
}

// 将英文从汉字中分开
func SplitEnFromCh(str string) string {
	r := []rune(str)
	enstr := ""
	for i := 0; i < len(r); i++ {
		if r[i] > 40869 || r[i] < 19968 {
			enstr = enstr + string(r[i])
		}
	}
	return enstr
}
