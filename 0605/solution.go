package main

import "fmt"

func main() {
	//[[1,2,3],[4,5,6],[7,8,9]]
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}

	self := spiralOrder(matrix)

	fmt.Println(self)
}

func spiralOrder(matrix [][]int) []int {

	xLen, yLen := len(matrix), 0
	if xLen > 0 {
		yLen = len(matrix[0])
	}
	output := make([]int, xLen*yLen)
	flag := 0
	xMin, yMin, xMax, yMax := 1, 0, xLen-1, yLen-1
	x, y := 0, 0

	for i := 0; i < xLen*yLen; i++ {
		output[i] = matrix[x][y]

		switch flag % 4 {
		case 0: // right
			if y == yMax {
				flag++
				yMax--
				x++
			} else {
				y++
			}
		case 1: // down
			if x == xMax {
				flag++
				xMax--
				y--
			} else {
				x++
			}
		case 2: // left
			if y == yMin {
				flag++
				yMin++
				x--
			} else {
				y--
			}
		case 3: // up
			if x == xMin {
				flag++
				xMin++
				y++
			} else {
				x--
			}
		}
	}
	return output
}
