package main

import "fmt"

func main() {
	numbs := []int{1, 2, 3, 4}
	self := productExceptSelf(numbs)

	fmt.Println(self)
}

func productExceptSelf(nums []int) []int {
	length := len(nums)
	var output = make([]int, length)
	var left = make([]int, length)
	left[0] = 1
	output[length-1] = 1

	for i := 1; i < length; i++ {
		left[i] = left[i-1] * nums[i-1]
		output[length-1-i] = output[length-i] * nums[length-i]
	}

	for i := 0; i < length; i++ {
		output[i] = left[i] * output[i]
	}
	return output
}
