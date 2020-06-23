package main

import (
	"fmt"
)

func main() {
	pattern := "aaaa"
	value := "dogcatdogcatdogcatdogcat"
	fmt.Println(patternMatching(pattern, value))
}

func patternMatching(pattern string, value string) bool {
	is, count := isSinglePattern(pattern)
	if is {
		size := len(value)
		subSize := size / count
		temp := value[0 : subSize-1]
		for i := 1; i < count; i++ {
			if temp != value[i*subSize:(i+1)*subSize-1] {
				return false
			}
		}
		return true
	} else {

	}
	return false
}

// isSinglePattern 是否 aaa 模式
// return 是/否 a的个数
func isSinglePattern(pattern string) (bool, int) {
	// temp:=
	temp := pattern[0:1]
	count := 0
	for _, c := range pattern {
		if string(c) != temp {
			return false, 0
		}
		count++
	}
	return true, count
}
