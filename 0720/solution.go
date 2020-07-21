package main

import "fmt"

func main() {
	numbers := []int{2, 3, 4}
	out := twoSum(numbers, 6)
	fmt.Println(out)
}

func twoSum(numbers []int, target int) []int {
	output := make([]int, 2)
	size := len(numbers)
	idxMap := make(map[int]int)
	for i := 0; i < size; i++ {
		idx, ok := idxMap[target-numbers[i]]
		if ok {
			return []int{idx + 1, i + 1}
		}
		idxMap[numbers[i]] = i
	}
	return output
}
