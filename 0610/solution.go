package main

import (
	"fmt"
)

func main() {
	num := 1001
	fmt.Println(isPalindrome(num))

	fmt.Println(lenNum(num))

	// fmt.Println(intAtIndex(12345, 0, 5))
	// fmt.Println(intAtIndex(12345, 1, 5))
	// fmt.Println(intAtIndex(12345, 2, 5))
	// fmt.Println(intAtIndex(12345, 3, 5))
	// fmt.Println(intAtIndex(12345, 4, 5))

}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	l := lenNum(x)

	for i := 0; i < l/2; i++ {
		if intAtIndex(x, i, l) != intAtIndex(x, l-1-i, l) {
			return false
		}
	}
	return true
}

func lenNum(num int) int {
	res := 0
	for ; num > 0; num = num / 10 {
		res++
	}
	return res
}

func intAtIndex(num int, index int, len int) int {
	for i := 0; i < len-1-index; i++ {
		num = num / 10
	}
	return num % 10
}
