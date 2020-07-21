package main

func main() {

}

func minSubArrayLen(s int, nums []int) int {
	size := len(nums)
	left, right := 0, 0
	sum := 0
	for right < size {
		if sum < s {
			sum = sum + nums[right]
			right++
		} else {

		}
	}
	return
}
