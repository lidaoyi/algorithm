package main

import (
	"fmt"
	"strings"
)

func main() {
	palindrome := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Println(palindrome)
}
func isPalindrome(s string) bool {
	s = strings.ToUpper(s)
	i := 0
	j := len(s) - 1
	for ; i < j; {
		if !isNeedChar(s[i]) {
			i++
			continue
		}
		if !isNeedChar(s[j]) {
			j--
			continue
		}
		fmt.Printf("i:%d ,j:%d, s[i]:%s, s[j]:%s \n", i, j, string(s[i]), string(s[j]))
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isNeedChar(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}
