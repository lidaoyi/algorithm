package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4

	fmt.Println(findKthLargest(nums, k))
}

func findKthLargest(nums []int, k int) int {
	temp := nums[0]
	index := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if index < k {
			if num < temp {
				temp = num
				index++
			} else {
				index++
			}
		} else {
			if num > temp {

			}
		}
	}
	return temp
}
