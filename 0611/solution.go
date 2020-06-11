package main

import "fmt"

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures1(T))
}

func dailyTemperatures(T []int) []int {
	l := len(T)
	if l == 1 {
		return []int{0}
	}
	output := make([]int, l)

	for i := 0; i < l-1; i++ {
		if T[i] < T[i+1] {
			output[i] = 1
		} else {
			for j := i + 1; j < l; j++ {
				if T[j] > T[i] {
					output[i] = j - i
					break
				}
			}
		}
	}
	return output
}

func dailyTemperatures1(T []int) []int {
	l := len(T)
	if l == 1 {
		return []int{0}
	}
	output := make([]int, l)

	for i := 0; i < l-1; i++ {
		if T[i] < T[i+1] {
			output[i] = 1
		} else {
			for j := i + 1; j < l; j++ {
				if T[j] > T[i] {
					output[i] = j - i
					break
				}
			}
		}
	}
	return output
}
