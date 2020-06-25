package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 2, 4, 8, 16, 32, 64, 128}
	fmt.Println(threeSumClosest(nums, 82))
}

func threeSumClosest(nums []int, target int) int {
	slice := []int{nums[0], nums[1], nums[2]}
	temp := slice[0] + slice[1] + slice[2]
	// sort.Ints(slice)
	size := len(nums)
	for i := 3; i < size; i++ {
		if temp == target {
			break
		} else {
			n := nums[i]
			sum := []int{n + slice[1] + slice[2], slice[0] + n + slice[2], n + slice[0] + slice[1]}
			idx := 0
			temp2 := sum[idx]
			for j := 1; j < 3; j++ {
				if math.Abs(float64(sum[j]-target)) < math.Abs(float64(temp2-target)) {
					idx = j
				}
			}
			if math.Abs(float64(sum[idx]-target)) < math.Abs(float64(temp-target)) {
				slice[idx] = n
				temp = slice[0] + slice[1] + slice[2]
			}
		}
	}
	return temp
}
