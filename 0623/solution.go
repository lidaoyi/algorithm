package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	max, min := a, b
	if len(a) < len(b) {
		max, min = b, a
	}
	sizeMax := len(max)
	sizeMin := len(min)
	buf := make([]byte, sizeMax+1)
	carryOver := uint8(0)

	for i := 0; i < sizeMax; i++ {
		idxMax := sizeMax - 1 - i
		idxMin := sizeMin - 1 - i
		idxBuf := sizeMax - i
		mx := parseInt(max[idxMax])
		mn := uint8(0)
		if idxMin >= 0 {
			mn = parseInt(min[idxMin])
		}
		sum := mx + mn + carryOver
		switch sum {
		case 0:
			carryOver = 0
			buf[idxBuf] = '0'
		case 1:
			carryOver = 0
			buf[idxBuf] = '1'
		case 2:
			carryOver = 1
			buf[idxBuf] = '0'
		case 3:
			carryOver = 1
			buf[idxBuf] = '1'
		}
	}
	res := ""
	if carryOver == 1 {
		buf[0] = '1'
		res = string(buf)
	} else {
		res = string(buf[1 : sizeMax+1])
	}
	return res
}

func parseInt(s byte) uint8 {
	return uint8(s) - 48
}
