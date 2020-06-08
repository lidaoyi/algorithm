package main

import "fmt"

func main() {
	equations := []string{"e==e","d!=e","c==d","d!=e"}
	output := equationsPossible(equations)
	fmt.Println(output)
}

func equationsPossible(equations []string) bool {

	equalVar := make(map[string]bool)

	unEqual := make(map[string]string)

	for i := 0; i < len(equations); i++ {
		equation := equations[i]
		x := equation[0:1]
		y := equation[3:4]

		flag := equation[1:3]

		switch flag {
		case "==":
			if x != y {
				equalVar[x] = true
				equalVar[y] = true
			}
		case "!=":
			if x == y {
				return false
			} else {
				unEqual[x] = y
			}
		}

		for k, v := range unEqual {
			if equalVar[k] && equalVar[v] {
				return false
			}
		}
	}
	return true
}
