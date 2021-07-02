package split

import "strings"

//Split("abcbe","b") => ["a","c","e"]
func Split(s, sep string) (ret []string) {
	// 提前使用make函数将result初始化为一个容量足够大的切片，而不再像之前一样通过调用append函数来追加
	ret = make([]string, 0, strings.Count(s, sep)+1)
	idx := strings.Index(s, sep)
	for idx > -1 {
		ret = append(ret, s[:idx])
		// s = s[idx+1:]
		s = s[idx+len(sep):]
		idx = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return
}
