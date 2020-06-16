package main

import "fmt"

func main() {

	input := []string{"flower", "flow", "flight"}
	fmt.Printf("output: %s ", longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	output := ""
	minSize := -1
	for _, str := range strs {

		fmt.Println(str)
		iSize := len(str)
		fmt.Println(iSize)
		if minSize == -1 || iSize < minSize {
			minSize = iSize
		}
	}

	fmt.Printf("minSize: %d\n", minSize)

	for i := 0; i < minSize; i++ {
		temp := strs[0][i : i+1]
		flag := true
		for _, str := range strs {
			if str[i:i+1] != temp {
				flag = false
				break
			}
		}
		if flag {
			output = output + temp
		} else {
			break
		}
	}
	return output
}
