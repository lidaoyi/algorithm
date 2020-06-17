package main

import "fmt"

func main() {
	A := []int{2, 2, 2}
	fmt.Println(maxScoreSightseeingPair(A))
}

//（A[i] + A[j] + i - j）
func maxScoreSightseeingPair1(A []int) int {
	size := len(A)
	max := 0
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			temp := A[i] + A[j] + i - j
			if A[i]+A[j]+i-j > max {
				max = temp
				fmt.Println(i, j, temp)
			}
		}
	}
	return max
}

func maxScoreSightseeingPair(A []int) int {
	size := len(A)
	maxA := make([]int, size)
	maxA[0] = A[0] + 0
	maxB := make([]int, size)
	maxB[size-1] = A[size-1] - size + 1

	for i := 1; i < size; i++ {
		tempA := A[i] + i

		if tempA > maxA[i-1] {
			maxA[i] = tempA
		} else {
			maxA[i] = maxA[i-1]
		}

		tempB := A[size-1-i] - size + 1 + i
		if tempB > maxB[size-i] {
			maxB[size-1-i] = tempB
		} else {
			maxB[size-1-i] = maxB[size-i]
		}
	}
	fmt.Println(maxA)
	fmt.Println(maxB)

	max := 0
	for i := 1; i < size; i++ {
		temp := maxA[i-1] + maxB[i]
		if temp > max {
			max = temp
		}
	}
	return max
}
