package main

import (
	"fmt"
	"strconv"
)

func main() {
	output := translateNum(122589)

	fmt.Println(output)

	fmt.Println(isUnique("leetcode"))
	nums := []int{2, 7, 11, 15}

	twoSum(nums, 9)

}

func isUnique(astr string) bool {
	charMap := make(map[string]bool)

	for i := 0; i < len(astr); i++ {
		item := astr[i : i+1]
		if charMap[item] {
			return false
		} else {
			charMap[item] = true
		}
	}
	return true
}

func twoSum(nums []int, target int) []int {
	tMap := make(map[int]int)
	output := make([]int, 2)

	for i, num := range nums {
		j, ok := tMap[target-num]
		if ok {
			if i != j {
				output = []int{i, j}
				break
			}
		} else {
			tMap[num] = i
		}
	}
	return output
}

func translateNum(num int) int {
	src := strconv.Itoa(num)
	p, q, r := 0, 0, 1
	for i := 0; i < len(src); i++ {
		p, q, r = q, r, 0
		r += q
		if i == 0 {
			continue
		}
		pre := src[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p
		}
	}
	return r
}
