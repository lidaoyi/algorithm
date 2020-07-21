package main

func main() {

}

func generateTrees(n int) []*TreeNode {
	output := make([]*TreeNode, 10)
	return output
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
